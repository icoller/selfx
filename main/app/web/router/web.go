/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 19:19:28
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
	// 网页路由
	route.Get("/robots.txt", controller.AssetsRobotsTxt)
	route.Get("/ads.txt", controller.AssetsAdsTxt)
	route.Get("/favicon.ico", controller.FaviconIco)
	route.Get(constant.LogoFilePath, controller.Logo)

	route.Get("/user/signIn", middleware.WebAuth, controller.UserSignIn)
	route.Get("/user/signUp", middleware.WebAuth, controller.UserSignUp)

	// sitemap
	sitemap := route.Group(config.Set.Router.GetSitemapPath())
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
	route.Get(config.Set.Router.GetCategoryPageRule(), middleware.Cache, middleware.MinifyCode, controller.HomeCategory).Name("category")
	route.Get(config.Set.Router.GetCategoryRule(), middleware.Cache, middleware.MinifyCode, controller.HomeCategory).Name("category")

	// tag
	route.Get(config.Set.Router.GetTagPageRule(), middleware.Cache, middleware.MinifyCode, controller.HomeTag).Name("tag")
	route.Get(config.Set.Router.GetTagRule(), middleware.Cache, middleware.MinifyCode, controller.HomeTag).Name("tag")

	// article
	articleRule := config.Set.Router.GetArticleRule()
	route.Get(articleRule, middleware.Cache, middleware.MinifyCode, controller.HomeArticle).Name("article")
	route.Put(articleRule, controller.HomeArticleViews)

	// 后端访问地址
	route.Use(middleware.ReplaceBodyContent(
		map[string]string{
			// 注意：静态页面打包时dir带前后斜杠，所以这里去掉斜杠
			"{{__DIR__}}": strings.Trim(config.Set.Router.GetAdminPath(), "/"),
		},
		// pwa文件 manifest.webmanifest 的类型是 application/octet-stream
		[]string{"text/html", "javascript", "text/css", "application/octet-stream"}))
	route.Use(config.Set.Router.GetAdminPath(), filesystem.New(filesystem.Config{
		Root:         http.FS(resources.Admin()),
		Index:        "index.html",
		NotFoundFile: "index.html",
		// Next 如果访问地址非管理路径，则执行Next
		Next: func(ctx *fiber.Ctx) bool {
			path := ctx.Path()
			dir := config.Set.Router.GetAdminPath()
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

	// not found
	route.All("*", middleware.MinifyCode, controller.HomeNotFound)
}

func auth() any {
	return middleware.Auth("token", func(token string) (string, bool) {
		if config.Set.Admin.VerifyJwtToken(token) {
			return "administrator", true
		}
		if config.Set.API.Enable && token == config.Set.API.SecretKey {
			return "api", true
		}
		return "", false
	})
}
