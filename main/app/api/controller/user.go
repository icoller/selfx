/*
 * @Author: coller
 * @Date: 2023-12-25 12:30:40
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 17:51:56
 * @Desc: 管理元
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/api/service"
	"selfx/constant"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 注册用户
func UserRegister(ctx *fiber.Ctx) error {
	if time.Since(constant.AppStartTime).Minutes() > 10 {
		return ctx.JSON(mapper.Fail("please restart the application to create an administrator within 10 minutes"))
	}
	obj, err := mapper.Config.BodyToAdminInit(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.AdminCreate(obj.Username, obj.Password)))
}
