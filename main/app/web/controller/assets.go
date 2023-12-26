/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 17:47:56
 * @Desc:
 */
package controller

import (
	"encoding/base64"
	"selfx/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AssetsRobotsTxt(ctx *fiber.Ctx) error {
	if config.Config.Template.RobotsTxt == "" {
		return ctx.Next()
	}
	return ctx.SendString(config.Config.Template.RobotsTxt)
}

func AssetsAdsTxt(ctx *fiber.Ctx) error {
	if config.Config.Template.AdsTxt == "" {
		return ctx.Next()
	}
	return ctx.SendString(config.Config.Template.AdsTxt)
}

func FaviconIco(ctx *fiber.Ctx) error {
	if config.Config.Template.FaviconIco == "" {
		return ctx.Next()
	}
	var bs64 = config.Config.Template.FaviconIco
	i := strings.Index(config.Config.Template.FaviconIco, ",")
	if i > 0 {
		bs64 = config.Config.Template.FaviconIco[i+1:]
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(bs64))
	return ctx.Type("ico").SendStream(dec)
}

func Logo(ctx *fiber.Ctx) error {
	if config.Config.Template.Logo == "" {
		return ctx.Next()
	}
	var bs64 = config.Config.Template.Logo
	i := strings.Index(config.Config.Template.Logo, ",")
	if i > 0 {
		bs64 = config.Config.Template.Logo[i+1:]
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(bs64))
	return ctx.Type("png").SendStream(dec)
}
