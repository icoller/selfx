/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 23:27:35
 * @Desc: 发送邮件
 */
package plugins

import (
	"crypto/tls"
	"fmt"
	"selfx/app/model"
	"selfx/app/plugin/entity"
	"selfx/app/service"
	"selfx/utils/isx"

	"gopkg.in/gomail.v2"
)

type SendEmail struct {
	Enable   bool   `json:"enable"`   // 启用
	FormMail string `json:"formMail"` // 来自地址
	FormName string `json:"formName"` // 来自名称
	Smtp     string `json:"smtp"`     // 发送Smtp
	Password string `json:"password"` // 密码
	SSL      bool   `json:"ssl"`      // 是否使用SSL
}

func NewSendEmail() *SendEmail {
	return &SendEmail{Enable: true, Smtp: "smtp.88.com", Password: "UEzRNRiuT44tQIXw", FormMail: "colo@88.com", FormName: "coller"}
}

func (d *SendEmail) Info() *entity.PluginInfo {
	return &entity.PluginInfo{
		ID:    "SendEmail",
		About: "发送邮件",
	}
}

func (d *SendEmail) Load(ctx *entity.Plugin) error {
	service.Verify.AddCreateBeforeEvents(d)
	return nil
}

func (s *SendEmail) VerifyCreateBefore(item *model.Verify) (err error) {
	if s.Enable {
		if item.Username != "" && isx.IsEmail(item.Username) {
			s.SmsEmail(item.Username, item.Code)
		}
	}
	return
}

func (s *SendEmail) SmsEmail(email, code string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(s.FormMail, s.FormName))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "您的验证码是")
	m.SetBody("text/html", `<div style="background-color:#ffffff;margin:0px auto;max-width:600px; padding-top:30px; overflow: hidden;font-family: Arial;"><div style="font-size:24px;font-weight:600;line-height:normal;text-align:center;color:#242833;">欢迎回来</div><div style="font-size:14px;line-height:normal;text-align:left;color:#000000;"><p>Hi,coller!</p><p>We received your email verification request. The verification code is:</p></div><div style="font-size:14px;line-height:normal;text-align:center;color:#000000;"><a href="#"style="display:inline-block;background: #000000 ;color: #ffffff ;font-family: inherit ;font-size:14px;font-weight:400;line-height:120%;margin:0;text-decoration:none;text-transform:none;padding:16px 32px 16px 32px; border-radius:4px;"target="_blank">`+code+`</a></div><div style="font-size:14px;line-height:normal; margin-top: 20px; text-align:left;color:#474F5E;background-color:#f5f7f8;vertical-align:top;padding:16px;"><p style="text-align:center">Need Help or Have Questions?Feel free to contact us.</p><p style="text-align: center;"><a style="color: #276eaf;"href="#"target="_blank">Unsubscribe</a></p></div></div>`) //发件信息和服务器信息

	d := gomail.NewDialer(
		s.Smtp,
		25,
		s.FormMail,
		s.Password,
	)

	if !s.SSL {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if err = d.DialAndSend(m); err != nil {
		fmt.Println("发送邮件出错" + err.Error())
	}
	return nil
}

func (d *SendEmail) Run(ctx *entity.Plugin) (err error) {
	return nil
}
