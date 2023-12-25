/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:19:41
 * @Desc:
 */
package mapper

import (
	"selfx/app/dto"
	"selfx/app/model"
	"selfx/app/plugin/entity"
	"selfx/app/plugin/service"
)

func PluginItemsToPluginInfoList(items []*entity.Plugin) (res []dto.PluginList) {
	for _, item := range items {
		var runErr string
		if item.RunError != nil {
			runErr = item.RunError.Error()
		}
		res = append(res, dto.PluginList{
			PluginInfo:  *item.Info,
			RunTime:     item.RunTime.Unix(),
			RunError:    runErr,
			RunCount:    item.RunCount,
			RunDuration: item.RunDuration.Milliseconds(),
			NextRunTime: service.Plugin.NextRunTime(item.CronID).Unix(),
		})
	}
	return
}

func StoreToArticlePost(item *model.Store) *model.ArticlePost {
	return &model.ArticlePost{
		Article: model.Article{
			ArticleBase: model.ArticleBase{
				Title:       item.Title,
				Slug:        item.Slug,
				CreateTime:  item.CreateTime,
				CategoryID:  item.CategoryID,
				Views:       item.Views,
				Thumbnail:   item.Thumbnail,
				Description: item.Description,
			},
			ArticleDetail: model.ArticleDetail{
				Keywords: item.Keywords,
				Content:  item.Content,
				Extends:  item.Extends,
			},
		},
		Tags:         item.Tags,
		CategoryName: item.CategoryName,
	}
}
