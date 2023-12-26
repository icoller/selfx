/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 10:53:53
 * @Desc: 路由
 */
package router

import (
	"net/http"
	"path/filepath"
	"selfx/app/middleware"
	"selfx/app/web/controller"
	"selfx/config"
	"selfx/constant"
	"selfx/init/template"
	"selfx/resources"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func Register(route fiber.Router) {
	route.Get("/test", controller.Test)
	// 后端访问地址
	route.Use(middleware.ReplaceBodyContent(
		map[string]string{
			// 注意：静态页面打包时dir带前后斜杠，所以这里去掉斜杠
			"{{__DIR__}}": strings.Trim(config.Config.Router.GetAdminPath(), "/"),
		},
		// pwa文件 manifest.webmanifest 的类型是 application/octet-stream
		[]string{"text/html", "javascript", "text/css", "application/octet-stream"}))

	route.Use(config.Config.Router.GetAdminPath(), filesystem.New(filesystem.Config{
		Root:         http.FS(resources.Admin()),
		Index:        "index.html",
		NotFoundFile: "index.html",
		// Next 如果访问地址非管理路径，则执行Next
		Next: func(ctx *fiber.Ctx) bool {
			path := ctx.Path()
			dir := config.Config.Router.GetAdminPath()
			if dir == path {
				return false
			}
			if !strings.HasSuffix(dir, "/") {
				dir = dir + "/"
			}
			if strings.HasPrefix(path, dir) {
				return false
			}
			return true
		},
	}))
	// 网页路由
	route.Get("/robots.txt", controller.AssetsRobotsTxt)
	route.Get("/ads.txt", controller.AssetsAdsTxt)
	route.Get("/favicon.ico", controller.FaviconIco)
	route.Get(constant.LogoFilePath, controller.Logo)

	// sitemap
	sitemap := route.Group(config.Config.Router.GetSitemapPath())
	sitemap.Get("/article.xml", middleware.Cache, controller.Sitemap.ArticleXML).Name("sitemap")
	sitemap.Get("/article.txt", middleware.Cache, controller.Sitemap.ArticleTXT).Name("sitemap")
	sitemap.Get("/category.xml", middleware.Cache, controller.Sitemap.CategoryXML).Name("sitemap")
	sitemap.Get("/category.txt", middleware.Cache, controller.Sitemap.CategoryTXT).Name("sitemap")
	sitemap.Get("/tag.xml", middleware.Cache, controller.Sitemap.TagXML).Name("sitemap")
	sitemap.Get("/tag.txt", middleware.Cache, controller.Sitemap.TagTXT).Name("sitemap")

	// home
	route.Get("/", middleware.Cache, middleware.MinifyCode, controller.HomeIndex).Name("home")

	// static路由应当放到  template page路由前面
	// 否则不能正确响应文件的content-Type
	// template public
	if currentThemePath, err := template.CurrentThemePath(); err == nil {
		route.Static("/", filepath.Join(currentThemePath, "public"))
	}
	// public
	route.Static("/", constant.PublicDir)

	// template page
	route.Get("/*", middleware.Cache, middleware.MinifyCode, controller.TemplatePage).Name("page")

	// category
	route.Get(config.Config.Router.GetCategoryPageRule(), middleware.Cache, middleware.MinifyCode, controller.HomeCategory).Name("category")
	route.Get(config.Config.Router.GetCategoryRule(), middleware.Cache, middleware.MinifyCode, controller.HomeCategory).Name("category")

	// tag
	route.Get(config.Config.Router.GetTagPageRule(), middleware.Cache, middleware.MinifyCode, controller.HomeTag).Name("tag")
	route.Get(config.Config.Router.GetTagRule(), middleware.Cache, middleware.MinifyCode, controller.HomeTag).Name("tag")

	// article
	articleRule := config.Config.Router.GetArticleRule()
	route.Get(articleRule, middleware.Cache, middleware.MinifyCode, controller.HomeArticle).Name("article")
	route.Put(articleRule, controller.HomeArticleViews)

	// not found
	route.All("*", middleware.MinifyCode, controller.HomeNotFound)
}
