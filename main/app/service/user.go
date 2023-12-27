/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:20:38
 * @Desc: 用户
 */
package service

import (
	"errors"
	"selfx/app/event"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
	"selfx/constant"
	"selfx/utils/cryptx"
	"strconv"
	"time"
)

var User = new(UserService)

type UserService struct {
	CreateBeforeEvents []event.UserCreateBefore
	CreateAfterEvents  []event.UserCreateAfter
	UpdateBeforeEvents []event.UserUpdateBefore
	UpdateAfterEvents  []event.UserUpdateAfter
	DeleteBeforeEvents []event.UserDeleteBefore
	DeleteAfterEvents  []event.UserDeleteAfter
	GetAfterEvents     []event.UserGetAfter
	ListAfterEvents    []event.UserListAfter
}

func (s *UserService) AddCreateBeforeEvents(ev ...event.UserCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *UserService) AddCreateAfterEvents(ev ...event.UserCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *UserService) AddUpdateBeforeEvents(ev ...event.UserUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *UserService) AddUpdateAfterEvents(ev ...event.UserUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *UserService) AddDeleteBeforeEvents(ev ...event.UserDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *UserService) AddDeleteAfterEvents(ev ...event.UserDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *UserService) AddGetAfterEvents(ev ...event.UserGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *UserService) AddListAfterEvents(ev ...event.UserListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *UserService) listAfterEvents(list []model.User) {
	for _, e := range s.ListAfterEvents {
		e.UserListAfter(list)
	}
}

// getAfterEvents
func (s *UserService) getAfterEvents(item *model.User) {
	for _, e := range s.GetAfterEvents {
		e.UserGetAfter(item)
	}
}

func (s *UserService) Get(id uint) (res *model.User, err error) {
	if id == 0 {
		return nil, constant.ErrIdRequired
	}
	if res, err = repo.User.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *UserService) GetByMobile(mobile string) (res *model.User, err error) {
	if mobile == "" {
		return nil, constant.ErrMobileRequired
	}
	if res, err = repo.User.GetByMobile(mobile); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *UserService) GetByEmail(email string) (res *model.User, err error) {
	if email == "" {
		return nil, constant.ErrEmailRequired
	}
	if res, err = repo.User.GetByEmail(email); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *UserService) GetByUsername(username string) (res *model.User, err error) {
	if username == "" {
		return nil, constant.ErrUsernameRequired
	}
	if res, err = repo.User.GetByUsername(username); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, constant.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *UserService) Save(item *model.User) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *UserService) Create(item *model.User) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.UserCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.User.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.UserCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *UserService) CreateInBatches(items []model.User, batchSize int) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.UserCreateBefore(&items[k]); err != nil {
				return
			}
		}
		if err = s.postCheck(&items[k]); err != nil {
			return
		}
	}
	if err = repo.User.CreateInBatches(items, batchSize); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.UserCreateAfter(&item)
		}
	}
	return
}

func (s *UserService) Update(item *model.User) (err error) {
	if item.ID == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.UserUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repo.User.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.UserUpdateAfter(item)
	}
	return
}

func (s *UserService) postCheck(item *model.User) error {
	if item.Username == "" {
		return constant.ErrUsernameRequired
	}
	if item.Email == "" && item.Mobile == "" {
		return constant.ErrEmailAndMobileRequired
	}
	return nil
}

func (s *UserService) Delete(id uint) (err error) {
	if id == 0 {
		return constant.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.UserDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repo.User.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.UserDeleteAfter(id)
	}
	return
}

func (s *UserService) HasUsername(username string) (bool, error) {
	if username == "" {
		return false, constant.ErrUsernameRequired
	}
	id, err := repo.User.GetIdByUsername(username)
	return id > 0, err
}

func (s *UserService) HasEmail(email string) bool {
	if email == "" {
		return false
	}
	id, err := repo.User.GetIdByEmail(email)
	if err != nil {
		return false
	}
	return id > 0
}

func (s *UserService) HasMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	id, err := repo.User.GetIdByMobile(mobile)
	if err != nil {
		return false
	}
	return id > 0
}

func (s *UserService) List(ctx *context.Context) (res []model.User, err error) {
	res, err = repo.User.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 根据id调用文章列表
func (s *UserService) ListByIds(ctx *context.Context, ids []int) (res []model.User, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repo.User.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

func (s *UserService) CountByWhere(where *context.Where) (res int64, err error) {
	return repo.User.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *UserService) CountTotal() (int64, error) {
	return repo.User.CountTotal()
}

// CountToday 统计今日添加数量
func (s *UserService) CountToday() (int64, error) {
	return repo.User.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *UserService) CountYesterday() (int64, error) {
	return repo.User.CountYesterday()
}

// CountLastFewDays 统计最近几日的
func (s *UserService) CountLastFewDays(n int) (int64, error) {
	return repo.User.CountLastFewDays(n)
}

// 验证密码
func (s *UserService) CheckPassword(userId uint, password, inputPassword string, ip string) error {
	loginFail, err := repo.LoginFail.GetByUserIdIp(userId, ip)
	var count int
	if err != nil {
		count = 0
	} else {
		count = loginFail.Count
	}
	var timeCheck bool
	// 就算密码通过了，依然拦截
	if count >= 5 {
		// 判断拦截时间是否失效
		if (loginFail.UpdatedAt.Unix() + 300000) < time.Now().Unix() {
			timeCheck = true
		} else {
			return errors.New("密码错误次数过多")
		}
	}
	// 验证密码是否通过
	if !cryptx.ValidatePassword(password, inputPassword) {
		// 记录密码错误
		if count > 0 { // 要更新
			newCount := count + 1
			if timeCheck { // 失效后重新记录
				newCount = 1
			}
			repo.LoginFail.UpdateCountByUserIdIp(newCount, userId, ip)
			count = newCount
		} else {
			repo.LoginFail.Create(&model.LoginFail{
				UserID: userId,
				Ip:     ip,
			})
			count = 1
		}
		if count >= 5 {
			return errors.New("密码错误次数过多")
		}
		return errors.New("密码错误，剩余尝试次数" + strconv.Itoa(constant.UserPasswordErrCount-count))
	}
	return nil
}
