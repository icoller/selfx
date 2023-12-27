/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:12:26
 * @Desc:
 */
package controller

import (
	webServ "selfx/app/web/service"

	"github.com/gofiber/fiber/v2"
)

var Sitemap = new(SitemapController)

type SitemapController struct{}

func (s SitemapController) ArticleTXT(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapArticleText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) CategoryTXT(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapCategoryText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) TagTXT(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapTagText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) ArticleXML(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapArticleXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}

func (s SitemapController) CategoryXML(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapCategoryXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}

func (s SitemapController) TagXML(ctx *fiber.Ctx) error {
	res, err := webServ.SitemapTagXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}
