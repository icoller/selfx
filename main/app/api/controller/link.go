/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:39:31
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func LinkList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Link.List(&repoCtx)))
}

func LinkCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Link.CountByWhere(&where)))
}

func LinkGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Link.Get(id)))
}

func LinkCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Link](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, service.Link.Create(obj)))
}

func LinkUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.Link](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Link.Update(obj)))
}

func LinkDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Link.Delete(id)))
}

func LinkBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Link.BatchDelete(ids)))
}

func LinkExistsURL(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Link.ExistsURL(string(ctx.Body()))))
}

func LinkLikeURL(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Link.ListLikeURL(nil, string(ctx.Body()))))
}

func LinkStatus(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	var item model.Link
	err = ctx.BodyParser(&item)
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	if item.Status {
		err = service.Link.EnableLink(id)
	} else {
		err = service.Link.DisableLink(id)
	}
	return ctx.JSON(mapper.Result(err))
}
