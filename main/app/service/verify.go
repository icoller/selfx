/*
 * @Author: Coller
 * @Date: 2022-01-04 19:49:04
 * @LastEditTime: 2023-12-27 11:03:23
 * @Desc: 公用方法
 */
package service

import (
	"errors"
	"selfx/app/repo"
	"selfx/constant"
	"sync"
	"time"
)

var Verify = new(VerifyService)

type VerifyService struct {
	ID uint
	sync.Mutex
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
