/*
 * @Author: coller
 * @Date: 2023-12-27 12:56:32
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 17:20:43
 * @Desc: 用户
 */
package service

import (
	"errors"
	"selfx/app/api/dto"
	"selfx/app/model"
	"selfx/app/repo"
	"selfx/app/service"
	"selfx/config"
	"selfx/constant"
	"selfx/pkg/token"
	"selfx/utils"
	"selfx/utils/cryptx"
	"selfx/utils/isx"
	"strings"
	"time"
)

func UserRegister(req *dto.UserRegister) (user *model.User, err error) {
	if config.Set.System.CloseUser {
		return user, errors.New("会员系统已关闭，无法注册")
	}
	if req.Mode == constant.UserModeMobile {
		if service.User.HasMobile(req.Username) {
			return user, errors.New("手机号已注册，请直接登录")
		}
	} else if req.Mode == constant.UserModeEmail {
		if service.User.HasEmail(req.Username) {
			return user, errors.New("邮箱已注册，请直接登录")
		}
		userArr := strings.Split(req.Username, "@")
		if len(userArr) != 2 {
			req.Username = userArr[1]
			return user, errors.New("邮箱地址不正确")
		}
	} else {
		return user, errors.New("注册类型错误")
	}
	salt := utils.GetRandString(constant.DefaultSaltLen, "")
	user = &model.User{
		Username: req.Username,
		Password: cryptx.EncodePassword(req.Password),
		Salt:     salt,
		Email:    req.Username,
		Mobile:   req.Username}
	if err = service.User.Create(user); err != nil {
		return user, err
	}
	return user, nil
}

func UserLogin(req *dto.UserLogin) (data *dto.UserLoginInfo, err error) {
	if config.Set.System.CloseUser {
		return data, errors.New("会员系统已关闭，无法登录")
	}
	var user *model.User
	if req.Mode == constant.UserModeUsername {
		user, err = service.User.GetByUsername(req.Username)
		if err != nil {
			return data, errors.New("手机号未注册")
		}
		if err = service.User.CheckPassword(user.ID, user.Password, req.Password, req.Ip); err != nil {
			return data, err
		}
	} else if req.Mode == constant.UserModeEmail || req.Mode == constant.UserModeMobile { // 邮件和手机号
		if isx.IsMobile(req.Username) {
			user, err = service.User.GetByMobile(req.Username)
			if err != nil {
				return data, errors.New("手机号未注册")
			}
		} else if isx.IsEmail(req.Username) {
			user, err = service.User.GetByEmail(req.Username)
			if err != nil {
				return data, errors.New("邮箱未注册")
			}
		}
		if err = service.Verify.CheckUsernameCode(req.Username, req.Code, constant.VerifyTypeIdLogin); err != nil {
			return data, err
		}
	} else {
		return data, errors.New("登录方式错误")
	}
	var setPass int
	if user.Password != "" {
		setPass = 1
	}
	userInfo := &dto.UserInfo{Id: user.ID, Email: user.Email, Username: user.Username, Mobile: user.Mobile, Salt: user.Salt, SetPass: setPass, Setting: user.Setting}
	// 获取token
	xAuth, err := token.Create(user.ID, user.Salt, 0)
	if err != nil {
		return data, errors.New("生成token错误")
	}
	// 更新上次登录时间
	repo.User.UpdateLoginAtById(user.ID)
	// 返回登录数据
	return &dto.UserLoginInfo{
		UserAuth: dto.UserAuth{
			XAuth:     xAuth,
			ExpiresIn: time.Now().Add(time.Hour * config.Set.System.JwtExpiresTime).Unix(),
		},
		UserInfo: userInfo,
	}, nil
}
