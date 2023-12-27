/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 10:20:56
 * @Desc: http 日志
 */
package middleware

import (
	"selfx/config"
	"selfx/init/log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HttpLog(ctx *fiber.Ctx) error {
	next := ctx.Next()
	if log.Visitor.IsClosed() && log.Spider.IsClosed() {
		return next
	}
	log.Client.InvokePoolHTTP(log.HttpData{
		RequestTime: ctx.Context().Time(),
		Status:      ctx.Context().Response.StatusCode(),
		Depth:       ctx.Context().ConnRequestNum(),
		IP:          getRequestIP(ctx),
		Method:      ctx.Method(),
		URL:         string(ctx.Context().URI().FullURI()),
		Referer:     string(ctx.Context().Referer()),
		UserAgent:   string(ctx.Context().UserAgent()),
		Headers:     string(ctx.Request().Header.RawHeaders()),
		Path:        ctx.Path(),
	})
	return next
}

func getRequestIP(ctx *fiber.Ctx) (ip string) {
	for _, v := range config.Set.Router.ProxyHeader {
		if ip = ctx.Get(v); ip != "" {
			arr := strings.Split(ip, ",")
			if len(arr) == 0 {
				continue
			}
			if arr[0] != "" {
				return arr[0]
			}
		}
	}
	return ctx.IP()
}
