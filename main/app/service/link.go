package service

import (
	"selfx/app/event"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
	"selfx/constant"
	"strings"
	"time"
)

var Link = new(LinkService)

type LinkService struct {
	CreateBeforeEvents []event.LinkCreateBefore
	CreateAfterEvents  []event.LinkCreateAfter
	UpdateBeforeEvents []event.LinkUpdateBefore
	UpdateAfterEvents  []event.LinkUpdateAfter
	DeleteBeforeEvents []event.LinkDeleteBefore
	DeleteAfterEvents  []event.LinkDeleteAfter
	GetAfterEvents     []event.LinkGetAfter
	ListAfterEvents    []event.LinkListAfter
}

func (s *LinkService) AddCreateBeforeEvents(ev ...event.LinkCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *LinkService) AddCreateAfterEvents(ev ...event.LinkCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *LinkService) AddUpdateBeforeEvents(ev ...event.LinkUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *LinkService) AddUpdateAfterEvents(ev ...event.LinkUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *LinkService) AddDeleteBeforeEvents(ev ...event.LinkDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *LinkService) AddDeleteAfterEvents(ev ...event.LinkDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *LinkService) AddGetAfterEvents(ev ...event.LinkGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *LinkService) AddListAfterEvents(ev ...event.LinkListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *LinkService) listAfterEvents(list []model.Link) {
	for _, e := range s.ListAfterEvents {
		e.LinkListAfter(list)
	}
}

func (s *LinkService) Save(item *model.Link) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *LinkService) Create(item *model.Link) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.LinkCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repo.Link.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.LinkCreateAfter(item)
	}
	return
}

func (s *LinkService) Update(item *model.Link) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.LinkUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.Link.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.LinkUpdateAfter(item)
	}
	return
}

func (s *LinkService) postCheck(item *model.Link) error {
	if item.Name == "" {
		return constant.ErrNameRequired
	}
	if item.URL == "" {
		return constant.ErrUrlRequired
	}
	return nil
}

func (s *LinkService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.LinkDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repo.Link.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.LinkDeleteAfter(id)
	}
	return
}

func (s *LinkService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *LinkService) Get(id int) (res *model.Link, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Link.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.LinkGetAfter(res)
	}
	return
}

func (s *LinkService) ExistsURL(url string) (bool, error) {
	if url == "" {
		return false, constant.ErrUrlRequired
	}
	id, err := repo.Link.GetIdByURL(url)
	return id > 0, err
}

//////////////////

func (s *LinkService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Link.CountByWhere(where)
}

// CountTotal 统计总数
func (s *LinkService) CountTotal() (int64, error) {
	return repo.Link.CountTotal()
}

func (s *LinkService) DisableLink(id int) error {
	return repo.Link.DisableLink(id)
}

func (s *LinkService) EnableLink(id int) error {
	return repo.Link.EnableLink(id)
}

func (s *LinkService) List(ctx *context.Context) (res []model.Link, err error) {
	res, err = repo.Link.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 根据id调用列表
func (s *LinkService) ListByIds(ctx *context.Context, ids []int) (res []model.Link, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Link.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListPublic 前台公开列表
func (s *LinkService) ListPublic(ctx *context.Context) (res []model.Link, err error) {
	if res, err = repo.Link.ListPublic(ctx); err != nil {
		return
	}
	s.listAfterEvents(res)
	return
}

// ListLikeURL 相似链接列表
func (s *LinkService) ListLikeURL(ctx *context.Context, url string) (res []model.Link, err error) {
	if url == "" {
		return nil, constant.ErrUrlRequired
	}
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "//")
	return repo.Link.ListLikeURL(ctx, url)
}

// ListDetectLink 开启检查的链接列表
func (s *LinkService) ListDetectLink(ctx *context.Context) (res []model.Link, err error) {
	list, err := repo.Link.ListDetectLink(ctx)
	if err != nil {
		return
	}
	// 排除延迟检测
	now := time.Now().Unix()
	for _, item := range list {
		if item.DetectDelay == 0 || now >= item.CreateTime+item.DetectDelay*60 {
			res = append(res, item)
		}
	}
	s.listAfterEvents(res)
	return
}
