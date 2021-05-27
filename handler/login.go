package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"gocms/model"
	"gocms/pkg/auth"
	"gocms/pkg/captcha"
	"gocms/pkg/config"
	"gocms/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/patrickmn/go-cache"
)

var (
	tokenCache = cache.New(time.Minute*5, time.Hour)
	session    = new(auth.Session)
)

// Login 登录页
func Login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.Set("Ref", c.Request.Referer())
		c.Set("Captcha", config.Captcha())
		c.HTML(http.StatusOK, "login.html", c.Keys)
		return
	}

	var p *model.AdminLogin

	switch c.Query("action") {
	case "verfiy":
		// 二步验证
		token := strings.TrimSpace(c.Query("token"))
		if v, ok := tokenCache.Get(token); !ok {
			c.Error(errors.ErrRequest)
			return
		} else if p, ok = v.(*model.AdminLogin); !ok {
			c.Error(errors.ErrRequest)
			return
		}
		if p.IP != c.ClientIP() {
			c.Error(errors.ErrRequest)
			return
		}
		code := strings.TrimSpace(c.Query("code"))
		if len(code) < 1 {
			c.Error(errors.ErrCode)
			return
		}
		p.Verifyed = true
		// 移除 token
		tokenCache.Delete(token)

	default:
		p = &model.AdminLogin{
			IP: c.ClientIP(), UA: c.Request.UserAgent(),
		}
		err := c.MustBindWith(p, binding.FormPost)
		if err != nil {
			return
		}
		if config.Captcha() != nil {
			if err = captcha.Recaptcha(c.Request); err != nil {
				c.Error(errors.ErrCapcha)
				return
			}
		}
	}

	u, err := p.Login()
	if err != nil {
		c.Error(err)
		return
	}
	if err == nil {
		if err = session.Set(c.Request, c.Writer, u); err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, errors.OK("/admin/"))
		return
	}
	if u != nil {
		hash := hmac.New(sha256.New, net.ParseIP(c.ClientIP()))
		fmt.Fprint(hash, u.ID, c.Request.UserAgent(), time.Now().Unix())
		token := base64.StdEncoding.EncodeToString(hash.Sum(nil))
		tokenCache.SetDefault(token, p)

		c.Header("X-Form-Action", "?action=verfiy")
		c.HTML(http.StatusOK, "login-verify", gin.H{
			"Token": token, "Email": p.Email,
		})
		return
	}
	c.AbortWithStatus(http.StatusForbidden)
}

// Logout 登出
func Logout(c *gin.Context) {
	session.Flush(c.Request, c.Writer)
	c.Redirect(http.StatusFound, "/login")
}

func init() {
	auth.Register(session)
}
