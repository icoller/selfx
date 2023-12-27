/*
 * @Author: coller
 * @Date: 2023-12-26 17:44:21
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:09:46
 * @Desc: 注册
 */
package dto

import (
	"selfx/app/model/helper"
)

type UserRegister struct { // 注册
	Username string `json:"username" validate:"required" info:"请输入用户名"` // 手机/邮箱/用户名
	Password string `json:"password" validate:"required" info:"请输入密码"`  // 密码
	Mode     string `json:"mode" validate:"required" info:"请输入注册类型"`    // 注册类型
	Code     string `json:"code"`                                       // 验证码
}

type UserLogin struct { // 登录
	Username string `json:"username" validate:"required" info:"请输入用户名"` // 手机/邮箱/用户名
	Password string `json:"password"`                                   // 密码
	Mode     string `json:"mode" validate:"required" info:"请输入登录类型"`    // 注册类型
	Code     string `json:"code"`                                       // 验证码
	Ip       string `json:"ip"`
}

type UserLoginInfo struct {
	UserAuth
	UserInfo *UserInfo `json:"userInfo"`
}

type UserAuth struct {
	XAuth     string `json:"xAuth"`
	ExpiresIn int64  `json:"expiresIn"`
}

type UserInfo struct {
	Id       uint        `redis:"id" json:"id"`
	Username string      `redis:"username" json:"username"`
	Email    string      `redis:"email" json:"email"`
	Mobile   string      `redis:"mobile" json:"mobile"`
	Avatar   string      `redis:"avatar" json:"avatar"`
	Salt     string      `redis:"salt" json:"salt"`
	Setting  helper.JSON `redis:"setting" json:"setting"`
	SetPass  int         `redis:"setPass" json:"setPass"`
}
