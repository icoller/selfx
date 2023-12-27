/*
 * @Author: coller
 * @Date: 2023-12-27 18:56:00
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 18:59:01
 * @Desc:
 */
package controller

import (
	"errors"
	webServ "selfx/app/web/service"
	"selfx/constant"
	"selfx/init/log"

	"github.com/gofiber/fiber/v2"
)

func UserSignIn(ctx *fiber.Ctx) error {
	view, err := webServ.UserSignIn()
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("category controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(view))
}

func UserSignUp(ctx *fiber.Ctx) error {
	view, err := webServ.UserSignUp()
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("category controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(view))
}
