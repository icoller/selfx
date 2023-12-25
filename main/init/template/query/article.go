/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:33:25
 * @Desc:
 */
package query

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/init/log"
)

type Article struct {
	limit   int
	order   string
	comment string
}

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) Limit(val int) *Article {
	a.limit = val
	return a
}

func (a *Article) Order(val string) *Article {
	a.order = val
	return a
}

func (a *Article) Comment(val string) *Article {
	a.comment = val
	return a
}
func (a *Article) context() *context.Context {
	if a.limit == 0 {
		a.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(a.limit, a.order, a.comment)
}

// Get by id
func (a *Article) Get(id int) *model.Article {
	res, err := service.Article.Get(id)
	log.WarnShortcut("template query error", err)
	return res
}

// List 调用文章列表
func (a *Article) List() (res []model.ArticleBase) {
	res, err := service.Article.List(a.context())
	log.WarnShortcut("template query error", err)
	return
}

// ListByID 根据ID调用文章列表
func (a *Article) ListByID(ids ...int) (res []model.ArticleBase) {
	res, err := service.Article.ListByIds(a.context(), ids)
	log.WarnShortcut("template query error", err)
	return
}

// ListByCategoryID 根据分类ID查询文章列表
func (a *Article) ListByCategoryID(id, categoryId int) (res []model.ArticleBase) {
	res, err := service.Article.ListByCategoryIDNotId(a.context(), id, categoryId)
	log.WarnShortcut("template query error", err)
	return
}

// ListByTags 方便模板中可以直接通过tags实体调用
func (a *Article) ListByTags(tags []model.Tag) []model.ArticleBase {
	var ids []int
	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}
	return a.ListByTagID(ids...)
}

// ListByTagID 通过tagId调用相关文章
func (a *Article) ListByTagID(ids ...int) (res []model.ArticleBase) {
	res, err := service.Article.ListByTagIds(a.context(), ids)
	log.WarnShortcut("template query error", err)
	return
}

// PseudorandomList 伪随机列表
func (a *Article) PseudorandomList() (res []model.ArticleBase) {
	res, err := service.Article.PseudorandomList(a.context())
	log.WarnShortcut("template query error", err)
	return
}
