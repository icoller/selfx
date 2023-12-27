/*
 * @Author: coller
 * @Date: 2023-12-27 12:30:38
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:24:38
 * @Desc: 登录记录
 */
package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/init/db"
)

var LoginRecord = new(LoginRecordRepo)

type LoginRecordRepo struct {
}

func (r *LoginRecordRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.LoginRecord{})
}

func (r *LoginRecordRepo) Create(item *model.LoginRecord) error {
	return db.DB.Create(item).Error
}

func (r *LoginRecordRepo) Delete(id uint) error {
	return db.DB.Delete(&model.LoginRecord{ID: id}).Error
}

func (r *LoginRecordRepo) List(ctx *context.Context) (res []model.LoginRecord, err error) {
	err = db.DB.Model(model.LoginRecord{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}
