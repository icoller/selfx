package controller

import (
	"selfx/app/api/mapper"
	apiService "selfx/app/api/service"
	"selfx/app/model"
	"selfx/app/service"

	"github.com/gofiber/fiber/v2"
)

func ArticleList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Article.List(&repoCtx)))
}

func ArticleCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Article.CountByWhere(&where)))
}

func ArticleGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Article.Get(id)))
}

func ArticleCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.ArticlePost](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, apiService.ArticlePost("create", obj)))
}

func ArticleUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[model.ArticlePost](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(obj, apiService.ArticlePost("update", obj)))
}

func ArticleDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.DeleteArticle(id)))
}

func ArticleBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.BatchDeleteArticle(ids)))
}

func ArticleExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Article.ExistsSlug(string(ctx.Body()))))
}

func ArticleExistsTitle(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.ResultData(service.Article.ExistsTitle(string(ctx.Body()))))
}

func ArticleGetTags(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.ResultData(service.Tag.ListByArticleID(nil, id)))
}

// ArticleCreateTag 创建文章标签
func ArticleCreateTag(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.CreateArticleTagByName(id, string(ctx.Body()))))
}

// ArticleCreateTagByNameList 创建文章标签通过name列表
func ArticleCreateTagByNameList(ctx *fiber.Ctx) error {
	tagNameList, err := mapper.BodyToStrSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.CreateArticleTagsByNameList(id, tagNameList)))
}

// ArticleDeleteTagByName 删除文章标签
func ArticleDeleteTagByName(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.DeleteArticleTagByName(id, string(ctx.Body()))))
}

// ArticleDeleteTagByIds 删除文章标签
func ArticleDeleteTagByIds(ctx *fiber.Ctx) error {
	tagIds, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(apiService.DeleteArticleTagByIds(id, tagIds)))
}

// ArticleBatchSetCategory 文章批量设置分类
func ArticleBatchSetCategory(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	categoryID, err := ctx.ParamsInt("category_id")
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(service.Article.BatchSetCategory(categoryID, ids)))
}
