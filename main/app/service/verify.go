/*
 * @Author: Coller
 * @Date: 2022-01-04 19:49:04
 * @LastEditTime: 2023-12-27 23:08:35
 * @Desc: 公用方法
 */
package service

import (
	"errors"
	"selfx/app/event"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/constant"
	"sync"
	"time"
)

var Verify = new(VerifyService)

type VerifyService struct {
	CreateBeforeEvents []event.VerifyCreateBefore
	sync.Mutex
}

func (s *VerifyService) AddCreateBeforeEvents(ev ...event.VerifyCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}

func (s *VerifyService) Create(item *model.Verify) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.VerifyCreateBefore(item); err != nil {
			return
		}
	}
	if err = repo.Verify.Create(item); err != nil {
		return
	}
	return
}

func (s *VerifyService) CheckUsernameCode(username string, code string, typeId uint) (err error) {
	verify, err := repo.Verify.GetByCodeUsernameTypeIdStatus(username, code, typeId, constant.VerifyStatusWait)
	if err != nil {
		return errors.New("验证码不正确")
	}
	if verify.ExpiredAt.Unix() < time.Now().Unix() {
		return errors.New("验证码已过期，请重新获取")
	}
	repo.Verify.UpdateVerifiedById(verify.ID)
	return nil
}
