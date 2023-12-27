package service

import (
	"errors"
	"path/filepath"
	"selfx/app/model"
	"selfx/app/service"
	"selfx/config"
	"selfx/init/template"
	"strconv"
)

func RenderIndex() ([]byte, error) {
	return template.Render("template/index.html", template.Binds{
		Page: template.Page{
			Name:        "index",
			Title:       config.Config.Site.Title,
			Keywords:    config.Config.Site.Keywords,
			Description: config.Config.Site.Description,
		},
	})
}

func RenderTemplatePage(path string) ([]byte, error) {
	return template.Render(filepath.Join("page", path), template.Binds{
		Page: template.Page{
			Name: "page",
			Path: path,
		},
		Data: map[string]any{},
	})
}

func RenderArticleBySlug(slug string) (_ []byte, err error) {
	item, err := service.Article.GetBySlug(slug)
	if err != nil {
		return
	}
	return RenderArticle(item)
}

func RenderArticle(item *model.Article) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	return template.Render("template/article.html", template.Binds{
		Page: template.Page{
			Name:        "article",
			Title:       item.Title + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
		},
		Data: item,
	})
}

func RenderCategoryBySlug(slug string, page int) (_ []byte, err error) {
	item, err := service.Category.GetBySlug(slug)
	if err != nil {
		return
	}
	return RenderCategory(item, page)
}

func RenderCategory(item *model.Category, page int) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	var pageTitle string
	if page > 1 {
		pageTitle = " - " + strconv.Itoa(page)
	}
	var title = item.Name
	if item.Title != "" {
		title = item.Title
	}
	return template.Render("template/category.html", template.Binds{
		Page: template.Page{
			Name:        "category",
			Title:       title + pageTitle + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
			PageNumber:  page,
		},
		Data: item,
	})
}

func RenderTagBySlug(slug string, page int) (_ []byte, err error) {
	item, err := service.Tag.GetBySlug(slug)
	if err != nil {
		return
	}
	return RenderTag(item, page)
}

func RenderTag(item *model.Tag, page int) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	var pageTitle string
	if page > 1 {
		pageTitle = " - " + strconv.Itoa(page)
	}
	var title = item.Name
	if item.Title != "" {
		title = item.Title
	}
	return template.Render("template/tag.html", template.Binds{
		Page: template.Page{
			Name:        "tag",
			Title:       title + pageTitle + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
			PageNumber:  page,
		},
		Data: item,
	})
}
