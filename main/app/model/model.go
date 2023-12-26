/*
 * @Author: coller
 * @Date: 2023-12-25 13:16:27
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 16:49:31
 * @Desc:
 */
package model

import (
	"selfx/app/model/helper"

	"gorm.io/gorm"
)

type ModelInterface interface {
	Article | Category | Tag | Link | Crawl
}

// 创建更新人
type ControlBy struct {
	UserId       uint `gorm:"column:user_id;type:uint;size:30;comment:创建者" json:"userId,omitempty"`
	UpdateUserId uint `gorm:"column:update_user_id;type:uint;size:30;comment:更新者" json:"updateUserId,omitempty"`
}

// 创建更新时间
type ModelTime struct {
	CreatedAt *helper.DateTime `gorm:"column:created_at;comment:创建时间" json:"createdAt,omitempty"`
	UpdatedAt *helper.DateTime `gorm:"column:updated_at;comment:更新时间" json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-" `
}
