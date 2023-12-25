/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 15:33:58
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func CrawlList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Crawl.List(&repoCtx)))
}

func CrawlCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Crawl.CountByWhere(&where)))
}

func CrawlGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Crawl.Get(id)))
}

func CrawlCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Crawl](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, service.Crawl.Create(obj)))
}

func CrawlUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Crawl](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Crawl.Update(obj)))
}

func CrawlDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Crawl.Delete(id)))
}

func CrawlBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Crawl.BatchDelete(ids)))
}

func CrawlPost(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(apiServ.CrawlPost(id)))
}
