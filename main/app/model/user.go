/*
 * @Author: Coller
 * @Date: 2022-05-08 11:59:54
 * @LastEditTime: 2023-12-26 17:14:46
 * @Desc: 用户表
 */
package model

import "selfx/app/model/helper"

type User struct {
	ID        uint            `gorm:"column:id;primaryKey;type:uint;size:30;comment:主键" json:"id"`
	Username  string          `gorm:"column:username;type:varchar(70);comment:账号名称;NOT NULL;" json:"username"` // 用户名
	Email     string          `gorm:"column:email;type:varchar(50);comment:邮箱地址;" json:"email"`                // 邮箱地址
	Mobile    string          `gorm:"column:mobile;type:varchar(30);comment:手机号码;NOT NULL;" json:"mobile"`     // 手机号
	Salt      string          `gorm:"column:salt;type:varchar(6);comment:安全符;" json:"salt"`                    // hash值
	Password  string          `gorm:"column:password;type:varchar(128);comment:账号密码;" json:"-"`
	SetPass   int             `gorm:"-" json:"setPass"`                                   // 是否设置密码
	Setting   helper.JSON     `gorm:"column:setting;type:json;comment:设置" json:"setting"` // 设置
	AppId     string          `gorm:"column:app_id;type:varchar(17);comment:appId;" json:"appId"`
	AppSecret string          `gorm:"column:app_secret;type:varchar(32);comment:appSecret;" json:"-"`
	Status    int             `gorm:"column:status;type:int;size:4;default:10;comment:状态;" json:"status"` // 状态
	LoginAt   helper.DateTime `gorm:"column:login_at;comment:上次登录时间" json:"loginAt,omitempty"`
	ModelTime
}
