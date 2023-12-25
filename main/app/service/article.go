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

var Article = new(ArticleService)

type ArticleService struct {
	CreateBeforeEvents []event.ArticleCreateBefore
	CreateAfterEvents  []event.ArticleCreateAfter
	UpdateBeforeEvents []event.ArticleUpdateBefore
	UpdateAfterEvents  []event.ArticleUpdateAfter
	DeleteBeforeEvents []event.ArticleDeleteBefore
	DeleteAfterEvents  []event.ArticleDeleteAfter
	GetAfterEvents     []event.ArticleGetAfter
	ListAfterEvents    []event.ArticleListAfter
}

func (s *ArticleService) AddCreateBeforeEvents(ev ...event.ArticleCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *ArticleService) AddCreateAfterEvents(ev ...event.ArticleCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *ArticleService) AddUpdateBeforeEvents(ev ...event.ArticleUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *ArticleService) AddUpdateAfterEvents(ev ...event.ArticleUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *ArticleService) AddDeleteBeforeEvents(ev ...event.ArticleDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *ArticleService) AddDeleteAfterEvents(ev ...event.ArticleDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *ArticleService) AddGetAfterEvents(ev ...event.ArticleGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *ArticleService) AddListAfterEvents(ev ...event.ArticleListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *ArticleService) listAfterEvents(list []model.ArticleBase) {
	for _, e := range s.ListAfterEvents {
		e.ArticleListAfter(list)
	}
}

// getAfterEvents
func (s *ArticleService) getAfterEvents(item *model.Article) {
	for _, e := range s.GetAfterEvents {
		e.ArticleGetAfter(item)
	}
}

func (s *ArticleService) Get(id int) (res *model.Article, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Article.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *ArticleService) GetBySlug(slug string) (res *model.Article, err error) {
	if slug == "" {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Article.GetBySlug(slug); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *ArticleService) Save(item *model.Article) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *ArticleService) Create(item *model.Article) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.ArticleCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repo.Article.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.ArticleCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *ArticleService) CreateInBatches(items []model.Article) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.ArticleCreateBefore(&items[k]); err != nil {
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
	if err = repo.Article.CreateInBatches(items); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.ArticleCreateAfter(&item)
		}
	}
	return
}

func (s *ArticleService) Update(item *model.Article) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.ArticleUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.Article.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.ArticleUpdateAfter(item)
	}
	return
}

func (s *ArticleService) postCheck(item *model.Article) error {
	if item.Slug == "" {
		return constant.ErrSlugRequired
	}
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

func (s *ArticleService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.ArticleDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repo.Article.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.ArticleDeleteAfter(id)
	}
	return
}

func (s *ArticleService) ExistsSlug(slug string) (bool, error) {
	if slug == "" {
		return false, constant.ErrSlugRequired
	}
	id, err := repo.Article.GetIdBySlug(slug)
	return id > 0, err
}

func (s *ArticleService) ExistsTitle(title string) (bool, error) {
	if title == "" {
		return false, constant.ErrTitleRequired
	}
	id, err := repo.Article.GetIdByTitle(title)
	return id > 0, err
}

func (s *ArticleService) UpdateViewsBySlug(id string, n int) error {
	return repo.Article.UpdateViewsBySlug(id, n)
}

//////-------list ------

// List 调用文章列表
func (s *ArticleService) List(ctx *context.Context) (res []model.ArticleBase, err error) {
	res, err = repo.Article.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListExistThumbnail 调用有缩略图文章列表
func (s *ArticleService) ListExistThumbnail(ctx *context.Context) (res []model.ArticleBase, err error) {
	res, err = repo.Article.ListExistThumbnail(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 根据id调用文章列表
func (s *ArticleService) ListByIds(ctx *context.Context, ids []int) (res []model.ArticleBase, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Article.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

func (s *ArticleService) ListByCategoryID(ctx *context.Context, id int) (res []model.ArticleBase, err error) {
	return s.ListByCategoryIds(ctx, []int{id})
}

func (s *ArticleService) ListByCategoryIDNotId(ctx *context.Context, id, categoryId int) (res []model.ArticleBase, err error) {
	res, err = repo.Article.ListByCategoryIDNotId(ctx, id, categoryId)
	s.listAfterEvents(res)
	return
}

// ListByCategoryIds 通过分类ID调用文章列表
func (s *ArticleService) ListByCategoryIds(ctx *context.Context, ids []int) (res []model.ArticleBase, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Article.ListByCategoryIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (s *ArticleService) ListAfterCreateTime(ctx *context.Context, t int64) (res []model.ArticleBase, err error) {
	res, err = repo.Article.ListAfterCreateTime(ctx, t)
	s.listAfterEvents(res)
	return
}

func (s *ArticleService) ListByTagID(ctx *context.Context, tagID int) (res []model.ArticleBase, err error) {
	return s.ListByTagIds(ctx, []int{tagID})
}

// ListByTagIds 根据标签ID调用文章列表
func (s *ArticleService) ListByTagIds(ctx *context.Context, tagIds []int) (res []model.ArticleBase, err error) {
	ids, err := Mapping.ListArticleIdsByTagIds(ctx, tagIds)
	if err != nil {
		return
	}
	res, err = s.ListByIds(context.NewContext(ctx.Limit, "id desc"), ids)
	return
}

// PseudorandomList 伪随机列表
func (s *ArticleService) PseudorandomList(ctx *context.Context) (res []model.ArticleBase, err error) {
	maxID, err := repo.Article.MaxID()
	if err != nil {
		return
	}
	return s.ListByIds(ctx, pseudorandomIds(maxID, ctx.Limit))
}

// ListDetail 调用详情表文章列表
func (s *ArticleService) ListDetail(ctx *context.Context) (res []model.ArticleDetail, err error) {
	res, err = repo.Article.ListDetail(ctx)
	return
}

func (s *ArticleService) ListDetailByIds(ctx *context.Context, ids []int) (res []model.ArticleDetail, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.Article.ListDetailByIds(ctx, ids)
	return
}

// CountByCategoryID 根据分类ID统计文章数
func (s *ArticleService) CountByCategoryID(categoryID int) (int64, error) {
	return repo.Article.CountByCategoryID(categoryID)
}

func (s *ArticleService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Article.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *ArticleService) CountTotal() (int64, error) {
	return repo.Article.CountTotal()
}

// CountToday 统计今日添加数量
func (s *ArticleService) CountToday() (int64, error) {
	return repo.Article.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *ArticleService) CountYesterday() (int64, error) {
	return repo.Article.CountYesterday()
}

// CountLastFewDays 统计最近几日的
func (s *ArticleService) CountLastFewDays(n int) (int64, error) {
	return repo.Article.CountLastFewDays(n)
}

func (s *ArticleService) MergeBaseListAndDetailList(v1 []model.ArticleBase, v2 []model.ArticleDetail) (res []model.Article) {
	for _, v := range v1 {
		detail, found := s.FindDetailListByID(v2, v.ID)
		if !found {
			continue
		}
		res = append(res, model.Article{ArticleBase: v, ArticleDetail: detail})
	}
	return
}

func (s *ArticleService) FindDetailListByID(list []model.ArticleDetail, id int) (res model.ArticleDetail, found bool) {
	for _, v := range list {
		if v.ArticleID == id {
			return v, true
		}
	}
	return res, false
}

// BatchSetCategory 批量设置分类
func (s *ArticleService) BatchSetCategory(categoryID int, ids []int) error {
	return repo.Article.BatchSetCategory(categoryID, ids)
}
