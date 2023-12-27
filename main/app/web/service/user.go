package service

import (
	"selfx/config"
	"selfx/init/template"
)

func UserSignIn() ([]byte, error) {
	return template.Render("template/user/signIn.html", template.Binds{
		Page: template.Page{
			Name:        "index",
			Title:       config.Set.Site.Title,
			Keywords:    config.Set.Site.Keywords,
			Description: config.Set.Site.Description,
		},
	})
}

func UserSignUp() ([]byte, error) {
	return template.Render("template/user/signUp.html", template.Binds{
		Page: template.Page{
			Name:        "index",
			Title:       config.Set.Site.Title,
			Keywords:    config.Set.Site.Keywords,
			Description: config.Set.Site.Description,
		},
	})
}
