/*
 * @Author: Coller
 * @Date: 2022-01-04 19:49:04
 * @LastEditTime: 2023-12-27 23:14:00
 * @Desc: 发送消息
 */
package service

import (
	"errors"
	"fmt"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/utils"

	"gopkg.in/gomail.v2"
)

var Send = new(SendService)

type SendService struct{}

func (s *SendService) SmsCode(username string, typeId uint, ip string) (err error) {
	code := utils.RandInt(1000, 9999)
	err = repo.Verify.Create(&model.Verify{
		Username: username,
		Code:     code,
		TypeId:   typeId,
		IP:       ip,
	})
	if err != nil {
		return errors.New("验证码发送错误")
	}
	// aliyun := &aliyun.SmsConfig{AccessKeyId: "", AccessKeySecret: "", EndPoint: ""}
	// if err = aliyun.SendCode(username, typeId, code); err != nil {
	// 	return err
	// }
	return nil
}

func (s *SendService) SmsEmail(username string, typeId uint, ip string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("colo@88.com", "coller"))
	m.SetHeader("To", "coller@tom.com")
	m.SetHeader("Subject", "您的验证码是")
	m.SetBody("text/plain", `<div style="background-color:#ffffff;margin:0px auto;max-width:600px; padding-top:30px; overflow: hidden;font-family: Arial;"><div style="font-size:24px;font-weight:600;line-height:normal;text-align:center;color:#242833;">欢迎回来</div><div style="font-size:14px;line-height:normal;text-align:left;color:#000000;"><p>Hi,coller!</p><p>We received your email verification request. The verification code is:</p></div><div style="font-size:14px;line-height:normal;text-align:center;color:#000000;"><a href="#"style="display:inline-block;background: #000000 ;color: #ffffff ;font-family: inherit ;font-size:14px;font-weight:400;line-height:120%;margin:0;text-decoration:none;text-transform:none;padding:16px 32px 16px 32px; border-radius:4px;"target="_blank">Visit store</a></div><div style="font-size:14px;line-height:normal; margin-top: 20px; text-align:left;color:#474F5E;background-color:#f5f7f8;vertical-align:top;padding:16px;"><p style="text-align:center">Need Help or Have Questions?Feel free to contact us.</p><p style="text-align: center;"><a style="color: #276eaf;"href="#"target="_blank">Unsubscribe</a></p></div></div>`) //发件信息和服务器信息

	d := gomail.NewDialer(
		"smtp.88.com",
		25,
		"colo@88.com",
		"UEzRNRiuT44tQIXw",
	)

	//	if !emailInfo.SSL {
	//	  d.TLSConfig := &tls.Config{InsecureSkipVerify:true}
	//	}

	if err = d.DialAndSend(m); err != nil {
		fmt.Println("发送邮件出错" + err.Error())

	}

	return nil
}
