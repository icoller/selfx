/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:02:46
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/api/service"
	"selfx/config"

	"github.com/gofiber/fiber/v2"
)

func ConfigList(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(mapper.Config.ConfigListToInfoList(config.Config.Items()), nil))
}

func ConfigGet(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(config.Config.Get(ctx.Params("id"))))
}

func ConfigUpdate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "admin" { // post管理员配置需要加密密码，所以单独执行
		return ConfigUpdateAdmin(ctx)
	}
	return ctx.JSON(mapper.MessageResult(config.Config.Save(id, ctx.Body())))
}

func ConfigUpdateAdmin(ctx *fiber.Ctx) error {
	obj, err := mapper.Config.BodyToAdminPost(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.AdminUpdate(obj.Username, obj.Password, obj.LoginExpire)))
}
