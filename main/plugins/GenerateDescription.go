/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:21:30
 * @Desc:
 */
package plugins

import (
	"selfx/app/model"
	"selfx/app/plugin/entity"
	"selfx/app/service"
	"selfx/utils/htmlx"
)

type GenerateDescription struct {
	Enable bool `json:"enable"` // 启用
	Length int  `json:"length"` // 长度
}

func NewGenerateDescription() *GenerateDescription {
	return &GenerateDescription{Enable: true, Length: 150}
}

func (d *GenerateDescription) Info() *entity.PluginInfo {
	return &entity.PluginInfo{
		ID:    "GenerateDescription",
		About: "generate description when created",
	}
}

func (d *GenerateDescription) Load(ctx *entity.Plugin) error {
	service.Article.AddCreateBeforeEvents(d)
	return nil
}

func (d *GenerateDescription) ArticleCreateBefore(item *model.Article) (err error) {
	if !d.Enable || d.Length <= 0 || item.Description != "" || item.Content == "" {
		return
	}

	text := htmlx.GetTextFromHTML(item.Content)
	if text == "" {
		return
	}
	textRune := []rune(text)
	if len(textRune) > d.Length {
		item.Description = string(textRune[0:d.Length])
	} else {
		item.Description = string(textRune[0:])
	}
	return
}

func (d *GenerateDescription) Run(ctx *entity.Plugin) (err error) {
	return nil
}
