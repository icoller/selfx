/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 17:51:20
 * @Desc:
 */
package controller

import (
	"path/filepath"
	"selfx/app/api/mapper"
	"selfx/init/template"

	"github.com/gofiber/fiber/v2"
)

func ThemeInit(ctx *fiber.Ctx) error {
	if err := template.InitTemplate(); err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(nil))
}

func ThemeList(ctx *fiber.Ctx) error {
	list, err := template.ThemeList()
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(list, nil))
}

func ThemeScreenshot(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	if id == "" {
		return ctx.JSON(mapper.Fail("id is required"))
	}
	dir := filepath.Base(id)
	return ctx.JSON(mapper.ResultData(template.ReadThemeScreenshot(dir), nil))
}
