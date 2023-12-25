/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:33:42
 * @Desc:
 */
package query

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/init/log"
)

type Tag struct {
	limit   int
	order   string
	comment string
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Limit(val int) *Tag {
	t.limit = val
	return t
}

func (t *Tag) Order(val string) *Tag {
	t.order = val
	return t
}

func (t *Tag) Comment(val string) *Tag {
	t.comment = val
	return t
}

func (t *Tag) context() *context.Context {
	if t.limit == 0 {
		t.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(t.limit, t.order, t.comment)
}

// Get by id
func (t *Tag) Get(id int) *model.Tag {
	res, err := service.Tag.Get(id)
	log.ErrorShortcut("template query error", err)
	return res
}

func (t *Tag) List() (res []model.Tag) {
	res, err := service.Tag.List(t.context())
	log.ErrorShortcut("template query error", err)
	return
}

func (t *Tag) ListByArticleID(ids ...int) (res []model.Tag) {
	res, err := service.Tag.ListByArticleIds(t.context(), ids)
	log.ErrorShortcut("template query error", err)
	return
}

func (t *Tag) ListByID(ids ...int) (res []model.Tag) {
	res, err := service.Tag.ListByIds(t.context(), ids)
	log.ErrorShortcut("template query error", err)
	return
}

// PseudorandomList 伪随机列表
func (t *Tag) PseudorandomList() (res []model.Tag) {
	if t.limit == 0 {
		t.limit = 20
	}
	res, err := service.Tag.PseudorandomList(t.context())
	log.ErrorShortcut("template query error", err)
	return
}
