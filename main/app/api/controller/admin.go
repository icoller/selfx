/*
 * @Author: coller
 * @Date: 2023-12-25 12:30:40
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:10:36
 * @Desc: 后台管理
 */
package controller

import (
	"selfx/app/api/dto"
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
	"selfx/constant"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AdminExists(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(apiServ.AdminExists(), nil))
}

func AdminCreate(ctx *fiber.Ctx) error {
	if time.Since(constant.AppStartTime).Minutes() > 10 {
		return ctx.JSON(mapper.Fail("please restart the application to create an administrator within 10 minutes"))
	}
	obj, err := mapper.Config.BodyToAdminInit(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiServ.AdminCreate(obj.Username, obj.Password)))
}

var loginLock sync.Mutex

func AdminLogin(ctx *fiber.Ctx) error {
	loginLock.Lock()
	defer loginLock.Unlock()
	time.Sleep(1200 * time.Millisecond)
	obj, err := mapper.Config.BodyToAdminLogin(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(apiServ.AdminLogin(obj.Username, obj.Password, obj.Captcha, obj.CaptchaID)))
}

func AdminCaptcha(ctx *fiber.Ctx) error {
	loginLock.Lock()
	defer loginLock.Unlock()
	time.Sleep(1200 * time.Millisecond)
	return ctx.JSON(mapper.ResultData(dto.NewCaptcha(apiServ.AdminCaptcha()), nil))
}
