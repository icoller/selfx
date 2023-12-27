/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 11:13:27
 * @Desc:
 */
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

var Tag = new(TagService)

type TagService struct {
	CreateBeforeEvents []event.TagCreateBefore
	CreateAfterEvents  []event.TagCreateAfter
	UpdateBeforeEvents []event.TagUpdateBefore
	UpdateAfterEvents  []event.TagUpdateAfter
	DeleteBeforeEvents []event.TagDeleteBefore
	DeleteAfterEvents  []event.TagDeleteAfter
	GetAfterEvents     []event.TagGetAfter
	ListAfterEvents    []event.TagListAfter
}

func (s *TagService) AddCreateBeforeEvents(ev ...event.TagCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *TagService) AddCreateAfterEvents(ev ...event.TagCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *TagService) AddUpdateBeforeEvents(ev ...event.TagUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *TagService) AddUpdateAfterEvents(ev ...event.TagUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *TagService) AddDeleteBeforeEvents(ev ...event.TagDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *TagService) AddDeleteAfterEvents(ev ...event.TagDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *TagService) AddGetAfterEvents(ev ...event.TagGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *TagService) AddListAfterEvents(ev ...event.TagListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *TagService) listAfterEvents(list []model.Tag) {
	for _, e := range s.ListAfterEvents {
		e.TagListAfter(list)
	}
}

// getAfterEvents
func (s *TagService) getAfterEvents(item *model.Tag) {
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(item)
	}
}

func (s *TagService) Save(item *model.Tag) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *TagService) Create(item *model.Tag) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.TagCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repo.Tag.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.TagCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *TagService) CreateInBatches(items []model.Tag, batchSize int) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.TagCreateBefore(&items[k]); err != nil {
				return
			}
		}
		if err = s.postCheck(&items[k]); err != nil {
			return
		}
		if items[k].CreateTime == 0 {
			items[k].CreateTime = time.Now().Unix()
		}
	}
	if err = repo.Tag.CreateInBatches(items, batchSize); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.TagCreateAfter(&item)
		}
	}
	return
}

func (s *TagService) Update(item *model.Tag) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.TagUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.Tag.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.TagUpdateAfter(item)
	}
	return
}

func (s *TagService) postCheck(item *model.Tag) error {
	if item.Slug == "" {
		return constant.ErrSlugRequired
	}
	if item.Name == "" {
		return constant.ErrNameRequired
	}
	if strings.HasPrefix(item.Slug, " ") {
		return constant.ErrSlugStartSpaceRequired
	}
	if strings.HasSuffix(item.Slug, " ") {
		return constant.ErrSlugEndSpaceRequired
	}
	return nil
}

func (s *TagService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.TagDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repo.Tag.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.TagDeleteAfter(id)
	}
	return
}

func (s *TagService) Get(id int) (res *model.Tag, err error) {
	if id == 0 {
		return nil, constant.ErrSlugRequired
	}
	if res, err = repo.Tag.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(res)
	}
	return
}

func (s *TagService) GetBySlug(slug string) (res *model.Tag, err error) {
	if slug == "" {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Tag.GetBySlug(slug); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(res)
	}
	return
}

func (s *TagService) ExistsSlug(slug string) (bool, error) {
	id, err := s.GetIdBySlug(slug)
	return id > 0, err
}

func (s *TagService) ExistsName(name string) (bool, error) {
	id, err := s.GetIdByName(name)
	return id > 0, err
}

// GetIdByName 通过name获取ID
func (s *TagService) GetIdByName(name string) (id int, err error) {
	if name == "" {
		return 0, constant.ErrNameRequired
	}
	return repo.Tag.GetIdByName(name)
}

// GetIdBySlug 通过slug获取ID
func (s *TagService) GetIdBySlug(slug string) (id int, err error) {
	if slug == "" {
		return 0, constant.ErrSlugRequired
	}
	return repo.Tag.GetIdBySlug(slug)
}

// GetIdByNameOrCreate 通过name获取主键,不存在则创建
func (s *TagService) GetIdByNameOrCreate(name string) (id int, err error) {
	if id, err = s.GetIdByName(name); err != nil {
		return
	}
	if id == 0 {
		id, err = s.CreateByNameReturnID(name)
	}
	return
}

func (s *TagService) CreateByName(name string) (res *model.Tag, err error) {
	res = &model.Tag{Name: name}
	err = s.Create(res)
	return
}

func (s *TagService) CreateByNameReturnID(name string) (int, error) {
	item, err := s.CreateByName(name)
	return item.ID, err
}

func (s *TagService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Tag.CountByWhere(where)
}

// CountTotal 统计总数
func (s *TagService) CountTotal() (int64, error) {
	return repo.Tag.CountTotal()
}

///////////////////////////////

func (s *TagService) List(ctx *context.Context) (res []model.Tag, err error) {
	res, err = repo.Tag.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 通过tagID获取列表
func (s *TagService) ListByIds(ctx *context.Context, ids []int) (res []model.Tag, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Tag.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListByArticleID 通过文章ID获取列表
func (s *TagService) ListByArticleID(ctx *context.Context, id int) (res []model.Tag, err error) {
	return s.ListByArticleIds(ctx, []int{id})
}

// ListByArticleIds 通过文章ID获取列表
func (s *TagService) ListByArticleIds(ctx *context.Context, articleIds []int) (res []model.Tag, err error) {
	ids, err := Mapping.ListTagIdByArticleIds(ctx, articleIds)
	if err != nil {
		return
	}
	res, err = s.ListByIds(nil, ids)
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (s *TagService) ListAfterCreateTime(ctx *context.Context, t int64) (res []model.Tag, err error) {
	res, err = repo.Tag.ListAfterCreateTime(ctx, t)
	s.listAfterEvents(res)
	return
}

// PseudorandomList 伪随机列表
func (s *TagService) PseudorandomList(ctx *context.Context) (res []model.Tag, err error) {
	maxID, err := repo.Tag.MaxID()
	if err != nil {
		return
	}
	return s.ListByIds(ctx, pseudorandomIds(maxID, ctx.Limit))
}
