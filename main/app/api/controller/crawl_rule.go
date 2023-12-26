/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 16:15:17
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func CrawlRuleList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Crawl.List(&repoCtx)))
}

func CrawlRuleGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Crawl.Get(id)))
}

func CrawlRuleAdd(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Crawl](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, service.Crawl.Create(obj)))
}

func CrawlRuleEdit(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Crawl](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Crawl.Update(obj)))
}

func CrawlRuleDel(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Crawl.Delete(id)))
}
