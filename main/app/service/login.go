/*
 * @Author: Coller
 * @Date: 2023-01-05 10:29:11
 * @LastEditTime: 2023-12-27 15:25:21
 * @Desc: 登录
 */
package service

import (
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/repo/context"
)

var Login = new(LoginService)

type LoginService struct{}

func (s *LoginService) RecordCreate(item *model.LoginRecord) (err error) {
	if err = repo.LoginRecord.Create(item); err != nil {
		return
	}
	return
}

func (s *LoginService) List(ctx *context.Context) (res []model.LoginRecord, err error) {
	res, err = repo.LoginRecord.List(ctx)
	return
}
