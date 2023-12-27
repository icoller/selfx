/*
 * @Author: Coller
 * @Date: 2022-01-04 20:02:21
 * @LastEditTime: 2023-12-27 13:51:47
 * @Desc: 用户登录错误表
 */
package model

type LoginFail struct {
	ID     uint   `gorm:"column:id;primaryKey;type:uint;size:30;comment:主键" json:"id"`
	UserID uint   `gorm:"column:user_id;type:uint;size:30;comment:用户ID;" json:"userId"`
	Count  int    `gorm:"column:count;type:int;size:4;default:1;comment:错误次数" json:"count"`
	Ip     string `gorm:"column:ip;type:varchar(80);comment:登录IP;" json:"ip"`
	ModelTime
}
