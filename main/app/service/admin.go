/*
 * @Author: coller
 * @Date: 2023-12-27 11:48:47
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 11:53:07
 * @Desc: 管理
 */
package service

import (
	"selfx/config"
	"selfx/config/service"
	"selfx/utils/timex"
)

type AdminService struct{}

// 更新管理员配置
func (s *AdminService) Update(username, password string, loginExpire timex.Duration) error {
	if err := config.Config.Admin.Update(username, password, loginExpire); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey() // 重置 jwtKey 主动使所有已登录失效
	return service.Push(config.Config.Admin)
}

// 更新管理员用户名
func (s *AdminService) UsernameUpdate(username string) error {
	if err := config.Config.Admin.UpdateUsername(username); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey()
	return service.Push(config.Config.Admin)
}

// 更新管理员密码
func (s *AdminService) PasswordUpdate(password string) error {
	if err := config.Config.Admin.UpdatePassword(password); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey()
	return service.Push(config.Config.Admin)
}

// 更新管理路径
func (s *AdminService) PathUpdate(path string) error {
	if err := config.Config.Router.UpdateAdminPath(path); err != nil {
		return err
	}
	return service.Push(config.Config.Router)
}
