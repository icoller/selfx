/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:25:30
 * @Desc:
 */
package repo

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/repo/gormx"
	"selfx/init/db"
)

var Mapping = new(MappingRepo)

type MappingRepo struct {
}

func (r *MappingRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&model.Mapping{})
}

// CreateArticleTag 创建文章标签映射
func (r *MappingRepo) CreateArticleTag(articleID, tagID int) error {
	return db.DB.Create(&model.Mapping{ArticleID: articleID, TagID: tagID}).Error
}

// DeleteArticleTag 删除文章标签映射 通过文章ID和标签ID
func (r *MappingRepo) DeleteArticleTag(articleID, tagID int) error {
	return db.DB.Where("article_id = ? and tag_id = ?", articleID, tagID).Delete(&model.Mapping{}).Error
}

// DeleteArticleTagByTagIds 删除文章标签映射 通过文章ID和标签id列表
func (r *MappingRepo) DeleteArticleTagByTagIds(articleID int, tagIds []int) error {
	return db.DB.Where("article_id = ? and tag_id in ?", articleID, tagIds).Delete(&model.Mapping{}).Error
}

// DeleteArticle 删除文章
func (r *MappingRepo) DeleteArticle(articleID int) error {
	return db.DB.Where("article_id = ?", articleID).Delete(&model.Mapping{}).Error
}

// DeleteTag 删除标签
func (r *MappingRepo) DeleteTag(tagID int) error {
	return db.DB.Where("tag_id = ?", tagID).Delete(&model.Mapping{}).Error
}

// ListArticleIdsByTagIds 通过标签ID查询文章ID列表
func (r *MappingRepo) ListArticleIdsByTagIds(ctx *context.Context, tagIds []int) (articleIds []int, err error) {
	err = db.DB.Model(&model.Mapping{}).Scopes(gormx.Context(ctx)).Where("tag_id in ?", tagIds).Pluck("article_id", &articleIds).Error
	return
}

// ListTagIdByArticleIds 通过文章ID获取tagID列表
func (r *MappingRepo) ListTagIdByArticleIds(ctx *context.Context, articleIds []int) (tagIds []int, err error) {
	err = db.DB.Model(&model.Mapping{}).Scopes(gormx.Context(ctx)).Where("article_id in ?", articleIds).Pluck("tag_id", &tagIds).Error
	return
}

// CountByTagID 根据tagID统计
func (r *MappingRepo) CountByTagID(tagID int) (res int64, err error) {
	err = db.DB.Model(model.Mapping{}).Where("tag_id = ?", tagID).Count(&res).Error
	return
}
