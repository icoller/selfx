/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:01:57
 * @Desc:
 */
package controller

import (
	"selfx/app/api/mapper"
	"selfx/config"
	"selfx/init/cache"

	"github.com/gofiber/fiber/v2"
)

func CacheInit(ctx *fiber.Ctx) error {
	if err := cache.Init(); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}

func CacheClear(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	// 执行之前，先查询一下是否存在设置的前缀，防止提交其他目录字符串导致安全问题
	if opt := config.Config.Cache.GetOption(name); opt == nil {
		return ctx.JSON(mapper.MessageFail("option not found"))
	}
	err := cache.ClearBucket(name)
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}
