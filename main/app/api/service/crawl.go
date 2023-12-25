/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 15:30:13
 * @Desc:
 */
package service

import (
	"selfx/app/mapper"
	"selfx/app/model"
	"selfx/app/service"
)

// CrawlPost 发布仓库文章
func CrawlPost(id int) (post *model.ArticlePost, err error) {
	item, err := service.Crawl.Get(id)
	if err != nil {
		return
	}
	post = mapper.CrawlToArticlePost(item)
	if err = ArticlePost("create", post); err != nil {
		return
	}
	return post, service.Crawl.Delete(id)
}
