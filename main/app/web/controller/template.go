/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:37:47
 * @Desc:
 */
package controller

import (
	"selfx/app/api/service"
	"selfx/init/log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// TemplatePage 模板自定义页面匹配
func TemplatePage(ctx *fiber.Ctx) error {
	var _path = ctx.Path()
	if _path == "" || _path == "/" {
		return ctx.Next()
	}
	// 不判断后缀
	// if path.Ext(_path) == "" {
	// 	_path = filepath.Join(_path, "index.html")
	// }
	b, err := service.Render.TemplatePage(_path)
	if err != nil {
		if strings.HasPrefix(err.Error(), "template not found:") {
			return ctx.Next()
		}
		log.Error("page controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}
