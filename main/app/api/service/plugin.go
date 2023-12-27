/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:07:50
 * @Desc: 初始化插件
 */
package service

import (
	"selfx/app/dto"
	"selfx/app/mapper"
	"selfx/app/plugin/factory"
	"selfx/app/plugin/service"
)

func PluginList() []dto.PluginList {
	return mapper.PluginItemsToPluginInfoList(service.Plugin.Items)
}

func PluginLogList(id string, page, limit int) (any, error) {
	filePath := factory.GetPluginLogFilePath(id)
	return logRead(filePath, page, limit)
}
