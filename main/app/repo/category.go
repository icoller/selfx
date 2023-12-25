package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/init/db"
)

var Category = new(CategoryRepo)

type CategoryRepo struct {
}

func (r *CategoryRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.Category{})
}

func (r *CategoryRepo) Create(item *model.Category) error {
	return db.DB.Create(item).Error
}

func (r *CategoryRepo) CreateInBatches(items []model.Category, batchSize int) error {
	return db.DB.CreateInBatches(items, batchSize).Error
}

func (r *CategoryRepo) Update(item *model.Category) error {
	return db.DB.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
}

func (r *CategoryRepo) Delete(id int) error {
	return db.DB.Delete(&model.Category{ID: id}).Error
}

func (r *CategoryRepo) Get(id int) (*model.Category, error) {
	var res model.Category
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetByName(name string) (*model.Category, error) {
	var res model.Category
	err := db.DB.Where("name = ?", name).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetBySlug(slug string) (*model.Category, error) {
	var res model.Category
	err := db.DB.Where("slug = ?", slug).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetIdByName(name string) (id int, err error) {
	err = db.DB.Model(model.Category{}).Where("name = ?", name).Limit(1).Pluck("id", &id).Error
	return
}

func (r *CategoryRepo) GetIdBySlug(slug string) (id int, err error) {
	err = db.DB.Model(model.Category{}).Where("slug = ?", slug).Limit(1).Pluck("id", &id).Error
	return
}

// MaxID 获取最大ID
func (r *CategoryRepo) MaxID() (res int, err error) {
	err = db.DB.Model(model.Category{}).Limit(1).Order("id desc").Limit(1).Pluck("id", &res).Error
	return
}

// CountByWhere 通过where获取统计结果
func (r *CategoryRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计总数
func (r *CategoryRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(model.Category{}).Count(&res).Error
	return
}

////////////////////////

// List 调用列表
func (r *CategoryRepo) List(ctx *context.Context) (res []model.Category, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 通过ids获取列表
func (r *CategoryRepo) ListByIds(ctx *context.Context, ids []int) (res []model.Category, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Context(ctx), gormx.WhereIds(ids)).Find(&res).Error
	return
}

// ListByParentIds 通过parentID获取列表
func (r *CategoryRepo) ListByParentIds(ctx *context.Context, ids []int) (res []model.Category, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Context(ctx, gormx.WhereParentIds(ids))).Find(&res).Error
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (r *CategoryRepo) ListAfterCreateTime(ctx *context.Context, t int64) (res []model.Category, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Context(ctx, gormx.WhereCreateTimeAfter(t))).Find(&res).Error
	return
}

// ListBeforeCreateTime 根据创建时间调用列表
func (r *CategoryRepo) ListBeforeCreateTime(ctx *context.Context, t int64) (res []model.Category, err error) {
	err = db.DB.Model(model.Category{}).Scopes(gormx.Context(ctx, gormx.WhereCreateTimeBefore(t))).Find(&res).Error
	return
}

// GetWithParent 获取分类和其夫分类
func (r *CategoryRepo) GetWithParent(id int) (res []model.Category, err error) {
	var current model.Category
	if err = db.DB.Where("id = ?", id).Find(&current).Error; err != nil || current.ID == 0 {
		return
	}
	res = append(res, current)
	if current.ParentID == 0 {
		return
	}
	var parent model.Category
	if err = db.DB.Where("id = ?", current.ParentID).Find(&parent).Error; err != nil || parent.ID == 0 {
		return
	}
	res = append(res, parent)
	return
}

// BatchSetParentCategory 批量设置父分类
func (r *CategoryRepo) BatchSetParentCategory(parentID int, ids []int) error {
	return db.DB.Model(&model.Category{}).Where("id in ?", ids).UpdateColumn("parent_id", parentID).Error
}
