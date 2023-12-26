/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 17:01:01
 * @Desc: 采集规则
 */
package model

type CrawlRule struct {
	ID   int    `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(120);default:''" json:"name"`

	ModelTime
}
