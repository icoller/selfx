/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:07:05
 * @Desc:
 */
package middleware

import (
	"selfx/config"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TLS(ctx *fiber.Ctx) error {

	if !config.Set.TLS.Enable {
		return ctx.Next()
	}

	// 强制http跳转到https
	if config.Set.TLS.ForceHTTPS && ctx.Protocol() == "http" {
		domain := strings.Split(ctx.Hostname(), ":")[0]
		port := config.Set.TLS.ListenPort()
		var portStr string
		if port != 443 {
			portStr = ":" + strconv.Itoa(port)
		}
		tlsURL := "https://" + domain + portStr + ctx.OriginalURL()
		return ctx.Redirect(tlsURL, config.Set.TLS.GetRedirectStatus())
	}

	return ctx.Next()

}
