/*
 * @Author: Coller
 * @Date: 2022-01-06 12:32:54
 * @LastEditTime: 2023-12-27 15:28:10
 * @Desc: 记录登录表
 */
package model

type LoginRecord struct {
	ID      uint   `gorm:"column:id;primaryKey;type:uint;size:30;comment:主键" json:"id"`
	UserId  uint   `gorm:"column:user_id;type:uint;size:30;comment:用户ID;" json:"userId"`
	Mode    string `gorm:"column:mode;type:varchar(30);comment:方式;" json:"mode"`
	Ip      string `gorm:"column:ip;type:varchar(64);comment:IP地址;" json:"ip"`
	Region  string `gorm:"column:region;type:varchar(128);comment:区域;" json:"region"`
	Browser string `gorm:"column:browser;type:varchar(64);comment:浏览器;" json:"browser"`
	Os      string `gorm:"column:os;type:varchar(64);comment:操作系统;" json:"os"`
	Remark  string `gorm:"column:remark;type:varchar(500);comment:user-agent" json:"remark"`
	ModelTime
}
