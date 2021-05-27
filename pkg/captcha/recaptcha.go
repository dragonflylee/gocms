package captcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gocms/pkg/config"
	"gocms/pkg/errors"
)

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// Recaptcha 确认验证码正确
func Recaptcha(r *http.Request) error {
	response := strings.TrimSpace(r.PostFormValue("g-recaptcha-response"))
	if len(response) <= 0 {
		return errors.ErrCapcha
	}
	conf := config.Captcha()
	data := make(url.Values)
	data.Add("secret", conf.Secret)
	data.Add("remoteip", r.RemoteAddr)
	data.Add("response", response)
	resp, err := http.PostForm(conf.API, data)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	var v recaptchaResponse
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}
	if v.Success {
		return nil
	}
	return errors.New(strings.Join(v.ErrorCodes, ","))
}
