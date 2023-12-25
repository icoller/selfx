/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:33:38
 * @Desc:
 */
package query

import (
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/init/log"
)

type Link struct {
	limit   int
	order   string
	comment string
}

func NewLink() *Link {
	return &Link{}
}

func (l *Link) Limit(val int) *Link {
	l.limit = val
	return l
}

func (l *Link) Order(val string) *Link {
	l.order = val
	return l
}

func (l *Link) Comment(val string) *Link {
	l.comment = val
	return l
}

func (l *Link) context() *context.Context {
	if l.limit == 0 {
		l.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(l.limit, l.order, l.comment)
}

// Get by id
func (l *Link) Get(id int) *model.Link {
	res, err := service.Link.Get(id)
	log.WarnShortcut("template query error", err)
	return res
}

// List 调用文章列表
func (l *Link) List() (res []model.Link) {
	res, err := service.Link.List(l.context())
	log.WarnShortcut("template query error", err)
	return
}

// ListByID 根据ID调用列表
func (l *Link) ListByID(ids ...int) (res []model.Link) {
	res, err := service.Link.ListByIds(l.context(), ids)
	log.WarnShortcut("template query error", err)
	return
}

// ListPublic 公开的链接列表
func (l *Link) ListPublic() (res []model.Link) {
	res, err := service.Link.ListPublic(l.context())
	log.WarnShortcut("template query error", err)
	return
}
