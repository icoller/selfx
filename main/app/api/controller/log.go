/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:39:37
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/api/service"
	"selfx/init/log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LogInit(ctx *fiber.Ctx) error {
	log.Init()
	return ctx.JSON(mapper.Result(nil))
}

func LogRead(ctx *fiber.Ctx) error {
	var page, _ = strconv.Atoi(ctx.Query("page"))
	var limit, _ = strconv.Atoi(ctx.Query("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	return ctx.JSON(mapper.ResultData(service.LogRead(ctx.Params("id"), page, limit)))
}
