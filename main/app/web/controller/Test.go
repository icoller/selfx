/*
 * @Author: coller
 * @Date: 2023-12-25 17:38:22
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 09:03:13
 * @Desc: 测试
 */
package controller

import "github.com/gofiber/fiber/v2"

type HotItem struct {
	Link  string
	Img   string
	Title string
	Desc  string
	Hot   int
}

func Test(ctx *fiber.Ctx) error {

	return nil
}
