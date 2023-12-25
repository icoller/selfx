/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:34:48
 * @Desc:
 */
package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/init/db"
	"time"
)

var Link = new(LinkRepo)

type LinkRepo struct {
}

func (r *LinkRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.Link{})
}

func (r *LinkRepo) Create(item *model.Link) error {
	return db.DB.Create(item).Error
}

func (r *LinkRepo) Update(item *model.Link) error {
	return db.DB.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
}

func (r *LinkRepo) Delete(id int) error {
	return db.DB.Delete(&model.Link{ID: id}).Error
}

func (r *LinkRepo) Get(id int) (*model.Link, error) {
	var res model.Link
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (r *LinkRepo) GetIdByURL(url string) (id int, err error) {
	err = db.DB.Model(&model.Link{}).Where("url = ?", url).Limit(1).Pluck("id", &id).Error
	return
}

/////////////////////////////////

// CountByWhere 通过where获取统计结果
func (r *LinkRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(&model.Link{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计总数
func (r *LinkRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(model.Link{}).Count(&res).Error
	return
}

func (r *LinkRepo) DisableLink(id int) error {
	return db.DB.Model(&model.Link{}).Where("id = ?", id).UpdateColumn("status", false).Error
}

func (r *LinkRepo) EnableLink(id int) error {
	return db.DB.Model(&model.Link{}).Where("id = ?", id).UpdateColumn("status", true).Error
}

// List 调用列表
func (r *LinkRepo) List(ctx *context.Context) (res []model.Link, err error) {
	err = db.DB.Model(&model.Link{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 根据id调用列表
func (r *LinkRepo) ListByIds(ctx *context.Context, ids []int) (res []model.Link, err error) {
	err = db.DB.Model(&model.Link{}).Scopes(gormx.Context(ctx, gormx.WhereIds(ids))).Find(&res).Error
	return
}

// ListPublic 前台公开列表
func (r *LinkRepo) ListPublic(ctx *context.Context) (res []model.Link, err error) {
	err = db.DB.Model(&model.Link{}).Where("status = ? and (expire_time = 0 or expire_time > ?)", true, time.Now().Unix()).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListLikeURL 相似链接列表
func (r *LinkRepo) ListLikeURL(ctx *context.Context, url string) (res []model.Link, err error) {
	err = db.DB.Model(&model.Link{}).Where("url like ?", "%"+url+"%").Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListDetectLink 开启检查的链接列表
func (r *LinkRepo) ListDetectLink(ctx *context.Context) (res []model.Link, err error) {
	err = db.DB.Model(&model.Link{}).Where("detect = ?", true).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}
