package controller

import (
	"errors"
	"selfx/app/service"
	webServ "selfx/app/web/service"
	"selfx/config"
	"selfx/constant"
	"selfx/init/log"
	"selfx/init/template"

	"github.com/gofiber/fiber/v2"
	"github.com/panjf2000/ants/v2"
)

func HomeIndex(ctx *fiber.Ctx) error {
	b, err := webServ.RenderIndex()
	if err != nil {
		log.Error("index controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeCategory(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	page, _ := ctx.ParamsInt("page", 1)
	if page == 0 {
		page = 1
	}
	// 超出最大页数限制
	if config.Set.Template.CategoryPageList.MaxPage > 0 && page > config.Set.Template.CategoryPageList.MaxPage {
		return ctx.SendStatus(404)
	}
	b, err := webServ.RenderCategoryBySlug(slug, page)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("category controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeTag(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	page, _ := ctx.ParamsInt("page", 1)
	if page == 0 {
		page = 1
	}
	// 限制最大页数
	if config.Set.Template.TagPageList.MaxPage > 0 && page > config.Set.Template.TagPageList.MaxPage {
		return ctx.SendStatus(404)
	}
	b, err := webServ.RenderTagBySlug(slug, page)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("tag controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}

	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeArticle(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	b, err := webServ.RenderArticleBySlug(slug)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("article controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeArticleViews(ctx *fiber.Ctx) error {
	if !ctx.XHR() {
		return ctx.SendStatus(200)
	}
	if config.Set.More.ArticleViewsPool == 0 || viewsPool.Free() == 0 {
		return ctx.SendStatus(200)
	}
	if err := viewsPool.Invoke(ctx.Params("slug")); err != nil {
		log.Warn("article views put in pool failed", log.Err(err))
	}
	return ctx.SendStatus(200)
}

var viewsPool, _ = ants.NewPoolWithFunc(config.Set.More.ArticleViewsPool, articleViewUpdate)

func articleViewUpdate(val any) {
	slug, ok := val.(string)
	if !ok {
		log.Warn("article slug transform error in views update")
		return
	}
	if err := service.Article.UpdateViewsBySlug(slug, 1); err != nil {
		log.Warn("article views update error", log.Err(err))
	}
}

func HomeNotFound(ctx *fiber.Ctx) error {
	b, err := template.Render("template/notFound.html", template.Binds{
		Page: template.Page{
			Name: "notFound",
			Path: ctx.Path(),
		},
	})
	if err != nil {
		return ctx.SendStatus(404)
	}
	return ctx.Type("html", "utf-8").Status(404).SendString(string(b))
}
