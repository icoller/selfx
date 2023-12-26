/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:39:15
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	appService "selfx/app/api/service"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func CategoryList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Category.List(&repoCtx)))
}

func CategoryCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Category.CountByWhere(&where)))
}

func CategoryGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Category.Get(id)))
}

func CategoryCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Category](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, service.Category.Create(obj)))
}

func CategoryUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Category](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Category.Update(obj)))
}

func CategoryDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Category.Delete(id)))
}

func CategoryBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Category.BatchDelete(ids)))
}

func CategoryExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Category.ExistsSlug(string(ctx.Body()))))
}

func CategoryExistsName(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Category.ExistsName(string(ctx.Body()))))
}

func CategoryTree(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(appService.CategoryTree()))
}

func CategoryBatchSetParentCategory(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	parentID, err := ctx.ParamsInt("parent_id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Category.BatchSetParentCategory(parentID, ids)))
}
