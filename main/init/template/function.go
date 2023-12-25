/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:01:02
 * @Desc:
 */
package template

import (
	"selfx/config"
	"selfx/config/entity"
	"selfx/init/template/query"
	"selfx/init/template/utils"
	"selfx/init/template/widget"
)

func initEngineFn(e Engine) {
	e.AddFunc("Config", newConfigHandle())
	e.AddFunc("Widget", widget.New())
	e.AddFunc("Utils", utils.New())

	e.AddFunc("Article", query.NewArticle)
	e.AddFunc("Category", query.NewCategory)
	e.AddFunc("Tag", query.NewTag)
	e.AddFunc("Link", query.NewLink)
}

type configHandle struct {
	Site     *entity.Site
	Router   *entity.Router
	Upload   *entity.Upload
	Cache    *entity.Cache
	Theme    *entity.Theme
	Template *entity.Template
	Sitemap  *entity.Sitemap
}

func newConfigHandle() *configHandle {
	return &configHandle{
		Site:     config.Config.Site,
		Router:   config.Config.Router,
		Upload:   config.Config.Upload,
		Cache:    config.Config.Cache,
		Theme:    config.Config.Theme,
		Template: config.Config.Template,
		Sitemap:  config.Config.Sitemap,
	}
}
