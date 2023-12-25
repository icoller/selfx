/*
 * @Author: coller
 * @Date: 2023-12-25 12:35:25
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:30:01
 * @Desc:
 */
package dto

import (
	"selfx/config/service"
	"selfx/utils/timex"
)

// ConfigInfo 配置信息对象
type ConfigInfo struct {
	ID   string         `json:"id"`
	Data service.Config `json:"data"`
}

// ConfigAdminInit 管理员初始化对象
type ConfigAdminInit struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConfigAdminLogin 管理员登录对象
type ConfigAdminLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captchaID"`
}

// ConfigAdminPost 管理员更新对象
type ConfigAdminPost struct {
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	LoginExpire timex.Duration `json:"login_expire"`
}
