/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 14:56:41
 * @Desc:
 */
package service

import (
	"errors"
	"selfx/config"
	"selfx/config/service"
	"selfx/init/captcha"
)

// AdminExists 判断管理员是否存在
func AdminExists() bool {
	return config.Config.Admin.Username != ""
}

// AdminCreate 创建管理员
func AdminCreate(username, password string) error {
	if AdminExists() {
		return errors.New("administrator already exists")
	}
	if err := config.Config.Admin.InitAdministrator(username, password); err != nil {
		return err
	}
	return service.Push(config.Config.Admin)
}

func AdminLogin(username, password, captchaAnswer, captchaID string) (token string, err error) {
	if err = captcha.Client.Verify(captchaID, captchaAnswer); err != nil {
		return
	}
	captcha.Client.Delete(captchaID) // 验证码成功后，清除旧的验证码，防止重复使用
	return config.Config.Admin.Login(username, password)
}

func AdminCaptcha() (bs64 string, id string) {
	return captcha.Client.StringSimple(4).Base64()
}
