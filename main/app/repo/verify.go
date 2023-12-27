/*
 * @Author: coller
 * @Date: 2023-12-27 10:33:00
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 10:50:24
 * @Desc:
 */
package repo

import (
	"selfx/app/model"
	"selfx/constant"
	"selfx/init/db"
)

func init() {}

var Verify = new(VerifyRepo)

type VerifyRepo struct{}

func (r *VerifyRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.Verify{})
}

func (r *VerifyRepo) GetByCodeUsernameTypeIdStatus(username, code string, typeId uint, status uint) (varify *model.Verify, err error) {
	if err = db.DB.Where("username = ? AND code = ? AND type_id = ? AND status = ?", username, code, typeId, status).First(&varify).Error; err != nil {
		return varify, err
	}
	return varify, nil
}

func (r *VerifyRepo) UpdateVerifiedById(id uint) (err error) {
	if err = db.DB.Model(&r).Where("id = ?", id).Update("status", constant.VerifyStatusSucc).Error; err != nil {
		return err
	}
	return nil
}
