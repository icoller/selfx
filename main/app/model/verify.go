/*
 * @Author: Coller
 * @Date: 2022-06-12 13:35:25
 * @LastEditTime: 2023-12-27 10:52:58
 * @Desc: 验证码
 */
package model

import "time"

type Verify struct {
	ID        uint      `gorm:"column:id;primaryKey;type:uint;size:30;comment:主键" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(50);comment:验证对象;" json:"username"`
	Code      string    `gorm:"column:code;type:varchar(10);comment:验证码;" json:"code"`
	TypeId    uint      `gorm:"column:typeId;type:int;size:4;default:10;comment:验证类型;" json:"typeId"`
	Status    int       `gorm:"column:status;type:int;size:4;default:10;comment:状态;" json:"status"`
	IP        string    `gorm:"column:ip;type:varchar(30);comment:IP地址;" json:"ip"`
	ExpiredAt time.Time `gorm:"column:expired_at;comment:过期时间" json:"expiredAt"`
	ModelTime
}
