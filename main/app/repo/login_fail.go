/*
 * @Author: coller
 * @Date: 2023-12-27 12:30:38
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:21:46
 * @Desc: 用户
 */
package repo

import (
	"selfx/app/model"
	"selfx/init/db"
)

var LoginFail = new(LoginFailRepo)

type LoginFailRepo struct {
}

func (r *LoginFailRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.LoginFail{})
}

func (r *LoginFailRepo) Create(item *model.LoginFail) error {
	return db.DB.Create(item).Error
}

func (r *LoginFailRepo) Delete(id uint) error {
	return db.DB.Delete(&model.LoginFail{ID: id}).Error
}

func (r *LoginFailRepo) GetByUserIdIp(userId uint, ip string) (res *model.LoginFail, err error) {
	err = db.DB.Where("user_id = ? AND ip = ?", userId, ip).First(&res).Error
	return res, err
}

func (r *LoginFailRepo) UpdateCountByUserIdIp(count int, userId uint, ip string) (err error) {
	return db.DB.Model(&model.LoginFail{}).Where("user_id = ? AND ip = ?", userId, ip).UpdateColumn("count", count).Error
}
