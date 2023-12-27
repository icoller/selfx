package controller

import (
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
	"selfx/app/plugin/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PluginList(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(apiServ.PluginList(), nil))
}

func PluginOptions(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Plugin.GetOptions(ctx.Params("id"))))
}

func PluginSaveOptions(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.Result(service.Plugin.UpdateOptions(ctx.Params("id"), ctx.Body())))
}

func PluginRun(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.Result(service.Plugin.Run(ctx.Params("id"))))
}

func PluginCronStart(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.Result(service.Plugin.UpdateCronStart(ctx.Params("id"), true)))
}

func PluginCronStop(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.Result(service.Plugin.UpdateCronStart(ctx.Params("id"), false)))
}

func PluginUpdateCronExp(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.Result(service.Plugin.UpdateCronExp(ctx.Params("id"), string(ctx.Body()))))
}

func PluginLogList(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "100"))
	return ctx.JSON(mapper.ResultData(apiServ.PluginLogList(ctx.Params("id"), page, limit)))
}
