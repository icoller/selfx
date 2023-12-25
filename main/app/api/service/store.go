/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:31:10
 * @Desc:
 */
package service

import (
	"selfx/app/mapper"
	"selfx/app/model"
	"selfx/app/service"
)

// StorePost 发布仓库文章
func StorePost(id int) (post *model.ArticlePost, err error) {
	item, err := service.Store.Get(id)
	if err != nil {
		return
	}
	post = mapper.StoreToArticlePost(item)
	if err = ArticlePost("create", post); err != nil {
		return
	}
	return post, service.Store.Delete(id)
}
