/*
 * @Author: coller
 * @Date: 2023-12-25 13:39:39
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:39:48
 * @Desc:
 */
package model

type Config struct {
	ID   string `gorm:"primaryKey;type:varchar(100);"`
	Data string `gorm:"type:string;"`
}
