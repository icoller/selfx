/*
 * @Author: coller
 * @Date: 2023-12-25 13:53:45
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 14:58:42
 * @Desc: 系统配置实体
 */
package entity

import "time"

type System struct {
	JwtSigningKey  string        `json:"jwtSigningKey"`  // jwt签名钥匙
	JwtExpiresTime time.Duration `json:"jwtExpiresTime"` // jwt过期时间
}

func NewSystem() *System {
	return &System{
		JwtSigningKey:  "selex",
		JwtExpiresTime: 24 * 7, // 小时
	}
}

func (*System) ConfigID() string {
	return "system"
}
