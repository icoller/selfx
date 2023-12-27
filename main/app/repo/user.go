/*
 * @Author: coller
 * @Date: 2023-12-27 12:30:38
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:10:36
 * @Desc: 用户
 */
package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/init/db"
	"selfx/utils/date"
	"time"
)

var User = new(UserRepo)

type UserRepo struct {
}

func (r *UserRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.User{})
}

func (r *UserRepo) Create(item *model.User) error {
	return db.DB.Create(item).Error
}

func (r *UserRepo) CreateInBatches(items []model.User, batchSize int) error {
	return db.DB.CreateInBatches(items, batchSize).Error
}

func (r *UserRepo) Update(item *model.User) error {
	return db.DB.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
}

func (r *UserRepo) Delete(id uint) error {
	return db.DB.Delete(&model.User{ID: id}).Error
}

func (r *UserRepo) Get(id uint) (*model.User, error) {
	var res model.User
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (r *UserRepo) GetByMobile(mobile string) (user *model.User, err error) {
	err = db.DB.Model(model.User{}).Where("mobile = ?", mobile).First(&user).Limit(1).Error
	return
}

func (r *UserRepo) GetByEmail(email string) (user *model.User, err error) {
	err = db.DB.Model(model.User{}).Where("email = ?", email).First(&user).Limit(1).Error
	return
}

func (r *UserRepo) GetByUsername(username string) (user *model.User, err error) {
	err = db.DB.Model(model.User{}).Where("username = ?", username).First(&user).Limit(1).Error
	return
}

func (r *UserRepo) GetIdByUsername(username string) (id int, err error) {
	err = db.DB.Model(model.User{}).Where("username = ?", username).Limit(1).Pluck("id", &id).Error
	return
}

func (r *UserRepo) GetIdByMobile(mobile string) (id int, err error) {
	err = db.DB.Model(model.User{}).Where("mobile = ?", mobile).Limit(1).Pluck("id", &id).Error
	return
}

func (r *UserRepo) GetIdByEmail(email string) (id int, err error) {
	err = db.DB.Model(model.User{}).Where("email = ?", email).Limit(1).Pluck("id", &id).Error
	return
}

// CountByWhere 通过where获取统计结果
func (r *UserRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(model.User{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计总数
func (r *UserRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(model.User{}).Count(&res).Error
	return
}

// List 调用列表
func (r *UserRepo) List(ctx *context.Context) (res []model.User, err error) {
	err = db.DB.Model(model.User{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 通过ids获取列表
func (r *UserRepo) ListByIds(ctx *context.Context, ids []int) (res []model.User, err error) {
	err = db.DB.Model(model.User{}).Scopes(gormx.Context(ctx), gormx.WhereIds(ids)).Find(&res).Error
	return
}

// CountToday 统计今日添加数量
func (r *UserRepo) CountToday() (res int64, err error) {
	err = db.DB.Model(model.User{}).Where("create_at >= ?", date.TodayBeginTime().Unix()).Count(&res).Error
	return
}

// CountYesterday 统计昨日添加数量
func (r *UserRepo) CountYesterday() (res int64, err error) {
	today := date.TodayBeginTime()
	yesterday := today.AddDate(0, 0, -1)
	err = db.DB.Model(model.User{}).Where("create_time >= ? and create_time < ?", yesterday.Unix(), today.Unix()).Count(&res).Error
	return
}

// CountLastFewDays 统计最近几日的数据
func (r *UserRepo) CountLastFewDays(n int) (res int64, err error) {
	today := date.TodayBeginTime()
	days := today.AddDate(0, 0, -n)
	err = db.DB.Model(model.User{}).Where("create_time >= ?", days.Unix()).Count(&res).Error
	return
}

func (r *UserRepo) UpdateLoginAtById(id uint) (err error) {
	return db.DB.Model(&model.User{}).Where("id = ?", id).UpdateColumn("login_at", time.Now()).Error
}
