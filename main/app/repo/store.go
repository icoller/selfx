package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/config"
	"selfx/constant"
	"selfx/init/db"
	"selfx/utils/date"

	"gorm.io/gorm"
)

var Store = new(StoreRepo)

type StoreRepo struct {
}

func (r *StoreRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.Store{})
}

func (r *StoreRepo) Create(item *model.Store) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		return tx.Create(item).Error
	})
}

func (r *StoreRepo) Update(item *model.Store) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		return tx.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
	})
}

func (r *StoreRepo) checkPost(tx *gorm.DB, item *model.Store) error {
	var id int
	// 判断 slug 是否存在
	if item.Slug != "" {
		if err := tx.Model(&model.Store{}).Where("slug = ? and id != ?", item.Slug, item.ID).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return constant.ErrSlugAlreadyExists
		}
		// 检查 article base 表是否存在标题
		if err := tx.Model(&model.ArticleBase{}).Where("slug = ?", item.Slug).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return constant.ErrSlugAlreadyExists
		}
	}
	// 判断 title 是否存在
	if config.Config.More.UniqueTitle {
		if err := tx.Model(&model.Store{}).Where("title = ? and id != ?", item.Title, item.ID).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return constant.ErrTitleAlreadyExists
		}
		// 检查 article 表是否存在标题
		if err := tx.Model(&model.Article{}).Where("title = ?", item.Title).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return constant.ErrTitleAlreadyExists
		}
	}
	return nil
}

func (r *StoreRepo) Delete(id int) error {
	return db.DB.Delete(&model.Store{ID: id}).Error
}

func (r *StoreRepo) Get(id int) (*model.Store, error) {
	var res model.Store
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

// CountByWhere 通过where获取统计结果
func (r *StoreRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(&model.Store{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计文章总数
func (r *StoreRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(model.Store{}).Count(&res).Error
	return
}

// CountToday 统计今日添加数量
func (r *StoreRepo) CountToday() (res int64, err error) {
	err = db.DB.Model(model.Store{}).Where("store_create_time > ?", date.TodayBeginTime().Unix()).Count(&res).Error
	return
}

// CountYesterday 统计昨日添加数量
func (r *StoreRepo) CountYesterday() (res int64, err error) {
	today := date.TodayBeginTime()
	yesterday := today.AddDate(0, 0, -1)
	err = db.DB.Model(model.Store{}).Where("store_create_time > ? and store_create_time < ?", yesterday.Unix(), today.Unix()).Count(&res).Error
	return
}

// List 调用列表
func (r *StoreRepo) List(ctx *context.Context) (res []model.Store, err error) {
	err = db.DB.Model(model.Store{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByCategoryIds 根据分类ID调用文章列表
func (r *StoreRepo) ListByCategoryIds(ctx *context.Context, categoryIds []int) (res []model.Store, err error) {
	err = db.DB.Model(&model.Store{}).Scopes(gormx.Context(ctx, gormx.WhereCategoryIds(categoryIds))).Find(&res).Error
	return
}
