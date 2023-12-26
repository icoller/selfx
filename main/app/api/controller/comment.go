/*
 * @Author: coller
 * @Date: 2023-12-25 12:30:40
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 17:32:23
 * @Desc: 管理元
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/api/service"

	"github.com/gofiber/fiber/v2"
)

func CommentAdd(ctx *fiber.Ctx) error {
	obj, err := mapper.Config.BodyToAdminInit(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.AdminCreate(obj.Username, obj.Password)))
}
