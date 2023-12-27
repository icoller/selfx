/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:12:19
 * @Desc: tag
 */
package controller

import (
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func TagList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.List(&repoCtx)))
}

func TagCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.CountByWhere(&where)))
}

func TagGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.Get(id)))
}

func TagCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Tag](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, service.Tag.Create(obj)))
}

func TagUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Tag](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Tag.Update(obj)))
}

func TagDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiServ.DeleteTag(id)))
}

func TagBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiServ.BatchDeleteTag(ids)))
}

func TagExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Tag.ExistsSlug(string(ctx.Body()))))
}

func TagExistsName(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Tag.ExistsName(string(ctx.Body()))))
}

func TagListByArticleID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.ListByArticleID(nil, id)))
}

func TagGetByIds(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.ListByIds(nil, ids)))
}
