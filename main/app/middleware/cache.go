/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:03:05
 * @Desc:
 */
package middleware

import (
	"selfx/config"
	"selfx/init/cache"
	"selfx/init/log"

	"github.com/gofiber/fiber/v2"
)

func Cache(ctx *fiber.Ctx) error {

	if !config.Set.Cache.Enable || ctx.Method() != "GET" {
		return ctx.Next()
	}

	name := ctx.Route().Name
	key := ctx.Path()
	option := config.Set.Cache.GetOption(name)

	if option == nil || !option.Enable {
		return ctx.Next()
	}

	if key == "" || key == "/" {
		key = "default"
	}

	// 默认不打印错误，否则找不到文件错误会爆满
	if val, err := cache.Get(name, key); err == nil {
		return ctx.Type("html").Send(val)
	}

	next := ctx.Next()

	if ctx.Response().StatusCode() == 200 {
		if err := cache.Set(name, key, ctx.Response().Body(), option.TTL.Duration()); err != nil {
			log.Warn("set cache error", log.Err(err))
		}
	}

	return next
}
