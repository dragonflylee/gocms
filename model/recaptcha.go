package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
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
		return errors.New("验证码非法")
	}
	data := make(url.Values)
	data.Add("secret", Config.Captcha.Secret)
	data.Add("remoteip", r.RemoteAddr)
	data.Add("response", response)
	resp, err := http.PostForm(Config.Captcha.API, data)
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
