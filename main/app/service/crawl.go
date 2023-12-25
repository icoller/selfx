package service

import (
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
	"selfx/constant"
	"strings"
	"time"
)

var Crawl = new(CrawlService)

type CrawlService struct {
}

func (s *CrawlService) Create(item *model.Crawl) (err error) {
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CrawlCreateTime == 0 {
		item.CrawlCreateTime = time.Now().Unix()
	}
	return repo.Crawl.Create(item)
}

func (s *CrawlService) Update(item *model.Crawl) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	return repo.Crawl.Update(item)
}

func (s *CrawlService) postCheck(item *model.Crawl) error {
	if item.Title == "" {
		return constant.ErrTitleRequired
	}
	if item.Content == "" {
		return constant.ErrContentRequired
	}
	if strings.HasPrefix(item.Slug, " ") {
		return constant.ErrSlugStartSpaceRequired
	}
	if strings.HasSuffix(item.Slug, " ") {
		return constant.ErrSlugEndSpaceRequired
	}
	return nil
}

func (s *CrawlService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	return repo.Crawl.Delete(id)
}

func (s *CrawlService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *CrawlService) Get(id int) (res *model.Crawl, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Crawl.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	return
}

func (s *CrawlService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Crawl.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *CrawlService) CountTotal() (int64, error) {
	return repo.Crawl.CountTotal()
}

// CountToday 统计今日添加数量
func (s *CrawlService) CountToday() (int64, error) {
	return repo.Crawl.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *CrawlService) CountYesterday() (int64, error) {
	return repo.Crawl.CountYesterday()
}

func (s *CrawlService) List(ctx *context.Context) (res []model.Crawl, err error) {
	res, err = repo.Crawl.List(ctx)
	return
}

func (s *CrawlService) ListByCategoryID(ctx *context.Context, id int) (res []model.Crawl, err error) {
	return s.ListByCategoryIds(ctx, []int{id})
}

// ListByCategoryIds 通过分类ID调用文章列表
func (s *CrawlService) ListByCategoryIds(ctx *context.Context, ids []int) (res []model.Crawl, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Crawl.ListByCategoryIds(ctx, ids)
	return
}
