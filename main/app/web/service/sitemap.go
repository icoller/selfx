package service

import (
	"selfx/app/dto"
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/config"
	configEntity "selfx/config/entity"
	"strings"
	"time"
)

// Sitemap 站点地图数据
var Sitemap = new(sitemap)

type sitemap struct{}

// ArticleList 文章站点地图数据列表
func SitemapArticleList() (res []model.ArticleBase, err error) {
	if config.Set.Sitemap.Article.Limit > 0 {
		res, err = service.Article.ListAfterCreateTime(SitemaplistOption(config.Set.Sitemap.Article))
	}
	return
}

// CategoryList 分类站点地图数据列表
func SitemapCategoryList() (res []model.Category, err error) {
	if config.Set.Sitemap.Category.Limit > 0 {
		res, err = service.Category.ListAfterCreateTime(SitemaplistOption(config.Set.Sitemap.Category))
	}
	return
}

// TagList 标签站点地图数据列表
func SitemapTagList() (res []model.Tag, err error) {
	if config.Set.Sitemap.Tag.Limit > 0 {
		res, err = service.Tag.ListAfterCreateTime(SitemaplistOption(config.Set.Sitemap.Tag))
	}
	return
}

func SitemaplistOption(opt *configEntity.SitemapOption) (ctx *context.Context, t int64) {
	if opt.InHours > 0 {
		t = time.Now().Unix() - int64(opt.InHours)*60*60
	}
	return context.NewContext(opt.Limit, "id desc"), t
}

func SitemapArticleText() (res string, err error) {
	var urls []string
	items, err := SitemapArticleList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func SitemapCategoryText() (res string, err error) {
	var urls []string
	items, err := SitemapCategoryList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func SitemapTagText() (res string, err error) {
	var urls []string
	items, err := SitemapTagList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func SitemapArticleXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := SitemapArticleList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Set.Sitemap.Article.ChangeFreq, config.Set.Sitemap.Article.Priority))
	}
	return xml.String()
}

func SitemapCategoryXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := SitemapCategoryList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Set.Sitemap.Category.ChangeFreq, config.Set.Sitemap.Category.Priority))
	}
	return xml.String()
}

func SitemapTagXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := SitemapTagList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Set.Sitemap.Tag.ChangeFreq, config.Set.Sitemap.Tag.Priority))
	}
	return xml.String()
}
