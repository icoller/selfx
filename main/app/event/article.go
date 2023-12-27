/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:28:55
 * @Desc: 文章事件
 */
package event

import "selfx/app/model"

type ArticleCreateBefore interface {
	ArticleCreateBefore(*model.Article) error
}

type ArticleCreateAfter interface {
	ArticleCreateAfter(*model.Article)
}

type ArticleUpdateBefore interface {
	ArticleUpdateBefore(item *model.Article) error
}

type ArticleUpdateAfter interface {
	ArticleUpdateAfter(item *model.Article)
}

type ArticleDeleteBefore interface {
	ArticleDeleteBefore(id int) error
}

type ArticleDeleteAfter interface {
	ArticleDeleteAfter(id int)
}

type ArticleGetAfter interface {
	ArticleGetAfter(*model.Article)
}

type ArticleListAfter interface {
	ArticleListAfter([]model.ArticleBase)
}
