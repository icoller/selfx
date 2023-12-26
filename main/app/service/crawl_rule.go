/*
 * @Author: coller
 * @Date: 2023-12-25 16:39:04
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 17:41:44
 * @Desc:
 */
package service

import (
	"selfx/app/model"
	"selfx/app/repo"
	"time"
)

func (s *CrawlService) RuleAdd(item *model.Crawl) (err error) {
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CrawlCreateTime == 0 {
		item.CrawlCreateTime = time.Now().Unix()
	}
	return repo.Crawl.Create(item)
}
