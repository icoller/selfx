/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:41:11
 * @Desc: 初始化插件
 */
package service

import (
	"selfx/app/dto"
	"selfx/app/mapper"
	"selfx/app/plugin/entity"
	"selfx/app/plugin/factory"
	"selfx/app/plugin/service"
	"selfx/init/log"
)

func PluginInit(items ...entity.PluginEntry) {
	for _, item := range items {
		if err := service.Plugin.Init(item); err != nil {
			log.Error("plugin loaded failed", log.Any("info", item.Info()), log.Err(err))
		} else {
			log.Debug("plugin loaded successfully", log.Any("info", item.Info()))
		}
	}
}

func PluginList() []dto.PluginList {
	return mapper.PluginItemsToPluginInfoList(service.Plugin.Items)
}

func PluginLogList(id string, page, limit int) (any, error) {
	filePath := factory.GetPluginLogFilePath(id)
	return logRead(filePath, page, limit)
}
