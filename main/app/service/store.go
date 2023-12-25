package service

import (
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
	"selfx/constant"
	"strings"
	"time"
)

var Store = new(StoreService)

type StoreService struct {
}

func (s *StoreService) Create(item *model.Store) (err error) {
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.StoreCreateTime == 0 {
		item.StoreCreateTime = time.Now().Unix()
	}
	return repo.Store.Create(item)
}

func (s *StoreService) Update(item *model.Store) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	return repo.Store.Update(item)
}

func (s *StoreService) postCheck(item *model.Store) error {
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

func (s *StoreService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	return repo.Store.Delete(id)
}

func (s *StoreService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *StoreService) Get(id int) (res *model.Store, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Store.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	return
}

func (s *StoreService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Store.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *StoreService) CountTotal() (int64, error) {
	return repo.Store.CountTotal()
}

// CountToday 统计今日添加数量
func (s *StoreService) CountToday() (int64, error) {
	return repo.Store.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *StoreService) CountYesterday() (int64, error) {
	return repo.Store.CountYesterday()
}

func (s *StoreService) List(ctx *context.Context) (res []model.Store, err error) {
	res, err = repo.Store.List(ctx)
	return
}

func (s *StoreService) ListByCategoryID(ctx *context.Context, id int) (res []model.Store, err error) {
	return s.ListByCategoryIds(ctx, []int{id})
}

// ListByCategoryIds 通过分类ID调用文章列表
func (s *StoreService) ListByCategoryIds(ctx *context.Context, ids []int) (res []model.Store, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Store.ListByCategoryIds(ctx, ids)
	return
}
