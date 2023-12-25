package service

import (
	"selfx/app/dto"
	"selfx/app/event"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
	"selfx/constant"
	"selfx/init/log"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/slice"
	"go.uber.org/zap"
)

var Category = new(CategoryService)

type CategoryService struct {
	CreateBeforeEvents []event.CategoryCreateBefore
	CreateAfterEvents  []event.CategoryCreateAfter
	UpdateBeforeEvents []event.CategoryUpdateBefore
	UpdateAfterEvents  []event.CategoryUpdateAfter
	DeleteBeforeEvents []event.CategoryDeleteBefore
	DeleteAfterEvents  []event.CategoryDeleteAfter
	GetAfterEvents     []event.CategoryGetAfter
	ListAfterEvents    []event.CategoryListAfter
}

func (s *CategoryService) AddCreateBeforeEvents(ev ...event.CategoryCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *CategoryService) AddCreateAfterEvents(ev ...event.CategoryCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *CategoryService) AddUpdateBeforeEvents(ev ...event.CategoryUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *CategoryService) AddUpdateAfterEvents(ev ...event.CategoryUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *CategoryService) AddDeleteBeforeEvents(ev ...event.CategoryDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *CategoryService) AddDeleteAfterEvents(ev ...event.CategoryDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *CategoryService) AddGetAfterEvents(ev ...event.CategoryGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *CategoryService) AddListAfterEvents(ev ...event.CategoryListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

func (s *CategoryService) listAfterEvents(list []model.Category) {
	for _, e := range s.ListAfterEvents {
		e.CategoryListAfter(list)
	}
}

func (s *CategoryService) Save(item *model.Category) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *CategoryService) Create(item *model.Category) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.CategoryCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repo.Category.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.CategoryCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *CategoryService) CreateInBatches(items []model.Category, batchSize int) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.CategoryCreateBefore(&items[k]); err != nil {
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
	if err = repo.Category.CreateInBatches(items, batchSize); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.CategoryCreateAfter(&item)
		}
	}
	return
}

func (s *CategoryService) Update(item *model.Category) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.CategoryUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.Category.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.CategoryUpdateAfter(item)
	}
	return
}

func (s *CategoryService) postCheck(item *model.Category) error {
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

func (s *CategoryService) Delete(id int) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.CategoryDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repo.Category.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.CategoryDeleteAfter(id)
	}
	return
}

func (s *CategoryService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *CategoryService) Get(id int) (res *model.Category, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.Category.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.CategoryGetAfter(res)
	}
	return
}

// GetOrCreate 不存在则创建
func (s *CategoryService) GetOrCreate(name string) (*model.Category, error) {
	res, err := s.GetByName(name)
	if err == nil && res.ID > 0 {
		return res, err
	}
	var createRes = model.Category{Name: name}
	err = s.Create(&createRes)
	return &createRes, err
}

func (s *CategoryService) GetByName(name string) (res *model.Category, err error) {
	if name == "" {
		return nil, constant.ErrNameRequired
	}
	if res, err = repo.Category.GetByName(name); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.CategoryGetAfter(res)
	}
	return
}

func (s *CategoryService) GetBySlug(slug string) (res *model.Category, err error) {
	if slug == "" {
		return nil, constant.ErrSlugRequired
	}
	if res, err = repo.Category.GetBySlug(slug); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.CategoryGetAfter(res)
	}
	return
}

func (s *CategoryService) ExistsSlug(slug string) (bool, error) {
	if slug == "" {
		return false, constant.ErrSlugRequired
	}
	id, err := repo.Category.GetIdBySlug(slug)
	return id > 0, err
}

func (s *CategoryService) ExistsName(name string) (bool, error) {
	if name == "" {
		return false, constant.ErrNameRequired
	}
	id, err := repo.Category.GetIdByName(name)
	return id > 0, err
}

func (s *CategoryService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.Category.CountByWhere(where)
}

// CountTotal 统计总数
func (s *CategoryService) CountTotal() (int64, error) {
	return repo.Category.CountTotal()
}

//////////////////////////////////////////
//////////////////////////////////////////

func (s *CategoryService) List(ctx *context.Context) (res []model.Category, err error) {
	res, err = repo.Category.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 通过ids获取列表
func (s *CategoryService) ListByIds(ctx *context.Context, ids []int) (res []model.Category, err error) {
	res, err = repo.Category.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListByParentID 通过parentID 获取列表
func (s *CategoryService) ListByParentID(ctx *context.Context, parentID int) (res []model.Category, err error) {
	return s.ListByParentIds(ctx, []int{parentID})
}

// ListByParentIds 通过parentIds 获取列表
func (s *CategoryService) ListByParentIds(ctx *context.Context, ids []int) (res []model.Category, err error) {
	res, err = repo.Category.ListByParentIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (s *CategoryService) ListAfterCreateTime(ctx *context.Context, t int64) (res []model.Category, err error) {
	res, err = repo.Category.ListAfterCreateTime(ctx, t)
	s.listAfterEvents(res)
	return
}

// PseudorandomList 伪随机列表
func (s *CategoryService) PseudorandomList(ctx *context.Context) (res []model.Category, err error) {
	maxID, err := repo.Category.MaxID()
	if err != nil {
		return
	}
	return s.ListByIds(ctx, pseudorandomIds(maxID, ctx.Limit))
}

// GetWithAncestors 获取分类和其祖先
func (s *CategoryService) GetWithAncestors(ctx *context.Context, id int) (_ []model.Category, err error) {
	if id <= 0 {
		return
	}
	all, err := s.List(ctx)
	return FindCategoryWithAncestors(id, all), err
}

// GetWithParent 获取分类和其夫分类
func (s *CategoryService) GetWithParent(id int) (res []model.Category, err error) {
	if id <= 0 {
		return
	}
	res, err = repo.Category.GetWithParent(id)
	s.listAfterEvents(res)
	return
}

// GetWithAncestorsReverse 获取分类和其祖先(祖先在前)
func (s *CategoryService) GetWithAncestorsReverse(ctx *context.Context, id int) (res []model.Category, err error) {
	res, err = s.GetWithAncestors(ctx, id)
	slice.Reverse(res)
	return
}

// ListDescendants 所有后代列表树
func (s *CategoryService) ListDescendants(ctx *context.Context, rootID int) (res []dto.CategoryTree, err error) {
	all, err := s.List(ctx)
	return MakeCategoryTree(CategoryModelListToCategoryTreeList(all), rootID), err
}

// ListChildren 子分类列表
func (s *CategoryService) ListChildren(ctx *context.Context, parentID int) (res []model.Category, err error) {
	return s.ListByParentID(ctx, parentID)
}

// ListRootWithChildren 根类目列表并包含子类目
func (s *CategoryService) ListRootWithChildren(ctx *context.Context) (res []dto.CategoryTree, err error) {
	list, err := s.ListByParentID(ctx, 0)
	if err != nil || len(list) == 0 {
		return
	}
	return s.ListByCategoriesWithChildren(ctx, list)
}

// ListByIdsWithChildren 根据ids获取列表并包含子分类
func (s *CategoryService) ListByIdsWithChildren(ctx *context.Context, ids []int) (res []dto.CategoryTree, err error) {
	if len(ids) == 0 {
		return
	}
	list, err := s.ListByIds(ctx, ids)
	if err != nil || len(list) == 0 {
		return
	}
	return s.ListByCategoriesWithChildren(ctx, list)
}

// ListByCategoriesWithChildren 给分类列表调用子分类
func (s *CategoryService) ListByCategoriesWithChildren(ctx *context.Context, list []model.Category) (res []dto.CategoryTree, err error) {
	res = CategoryModelListToCategoryTreeList(list)
	var parentIds []int
	for _, v := range list {
		parentIds = append(parentIds, v.ID)
	}
	children, err := s.ListByParentIds(ctx, parentIds)
	if err != nil {
		return
	}
	childrenTree := CategoryModelListToCategoryTreeList(children)
	for k, item := range list {
		res[k].Children = MakeCategoryTree(childrenTree, item.ID)
	}
	return
}

// BatchSetParentCategory 批量设置父分类
func (s *CategoryService) BatchSetParentCategory(parentID int, ids []int) error {
	return repo.Category.BatchSetParentCategory(parentID, ids)
}

func CategoryModelListToCategoryTreeList(items []model.Category) (res []dto.CategoryTree) {
	for _, item := range items {
		res = append(res, CategoryModelToCategoryTree(item))
	}
	return
}

func CategoryModelToCategoryTree(item model.Category) (res dto.CategoryTree) {
	res.Category = item
	return
}

func MakeCategoryTree(list []dto.CategoryTree, parentID int) (res []dto.CategoryTree) {
	for _, v := range list {
		if v.ParentID == parentID {
			var children = MakeCategoryTree(list, v.ID)
			if len(children) > 0 {
				v.Children = children
			}
			res = append(res, v)
		}
	}
	return res
}

// FindCategoryWithAncestors 查找类目和其祖先
func FindCategoryWithAncestors(id int, ranges []model.Category) (res []model.Category) {
	var fn func(int)
	var count = 0
	fn = func(cid int) {
		count++
		for _, item := range ranges {
			if item.ID == cid {
				// 如果查询了100次还在继续，防止程序进入死循环，直接break
				if count > 100 {
					log.Warn("循环超过100次", zap.Any("结果集", res))
					break
				}
				// 如果条目已经存在结果集中，说明类目即将陷入死循环，防止崩溃，直接break
				if idInCategories(item.ID, res) {
					log.Warn("已存在结果在结果集中，检查是否有循环依赖的分类", zap.Int("id", item.ID), zap.Any("结果集", res))
					break
				}
				res = append(res, item)
				if item.ParentID != 0 {
					fn(item.ParentID)
				}
			}
		}
	}
	fn(id)
	return
}

func idInCategories(id int, items []model.Category) bool {
	for _, item := range items {
		if item.ID == id {
			return true
		}
	}
	return false
}
