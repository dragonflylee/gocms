package mail

import (
	"log"

	"gocms/pkg/config"
	"gocms/pkg/errors"

	"gopkg.in/gomail.v2"
)

func SendEmail(email *Email) error {
	// 定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	conf := config.Smtp()
	if conf == nil {
		return errors.ErrConfig
	}
	m := gomail.NewMessage()
	m.SetHeader("From", conf.User)             // 发件人
	m.SetHeader("To", string(email.Recipient)) // 发送给多个用户
	m.SetHeader("Subject", email.Subject)      // 设置邮件主题
	m.SetBody("text/html", email.Body)         // 设置邮件正文

	d := gomail.NewDialer(conf.Host, conf.Port, conf.User, conf.Pass)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("SendEmail %s failed: %v", email.Recipient, err)
		return nil
	}
	log.Printf("SendEmail %s OK %s", email.Recipient, email.Subject)
	return nil
}
