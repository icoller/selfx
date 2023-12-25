/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:35:09
 * @Desc:
 */
package controller

import (
	"errors"
	"selfx/app/api/mapper"
	appService "selfx/app/api/service"
	"selfx/app/service"
	"selfx/init/db"
	"time"

	"github.com/gofiber/fiber/v2"
)

var Dashboard = new(dashboard)

type dashboard struct {
}

func (d *dashboard) Controller(ctx *fiber.Ctx) (err error) {
	var data any
	switch ctx.Params("id") {
	case "systemLoad":
		data = appService.SystemLoadPercent()
	case "systemCPU":
		data, err = appService.SystemCPUPercent(time.Second)
	case "systemMemory":
		data, err = appService.SystemMemoryPercent()
	case "systemDisk":
		data, err = appService.SystemDiskPercents()
	case "appCPU":
		data, err = appService.AppCPUPercent()
	case "appMemory":
		data, err = appService.AppUsedMemory()
	case "appInfo":
		data = appService.AppInfo()
	case "database":
		data = db.GetSize()
	case "log":
		data, err = appService.LogDirSize()
	case "cache":
		data, err = appService.CacheSize()
	case "articleTotal":
		data, err = service.Article.CountTotal()
	case "articleToday":
		data, err = service.Article.CountToday()
	case "articleYesterday":
		data, err = service.Article.CountYesterday()
	case "articleLast7days":
		data, err = service.Article.CountLastFewDays(7)
	case "articleLast30days":
		data, err = service.Article.CountLastFewDays(30)
	case "storeTotal":
		data, err = service.Store.CountTotal()
	case "storeToday":
		data, err = service.Store.CountToday()
	case "storeYesterday":
		data, err = service.Store.CountYesterday()
	case "categoryTotal":
		data, err = service.Category.CountTotal()
	case "tagTotal":
		data, err = service.Tag.CountTotal()
	case "linkTotal":
		data, err = service.Link.CountTotal()
	default:
		return ctx.JSON(mapper.MessageResult(errors.New("id is undefined")))
	}
	return ctx.JSON(mapper.MessageResultData(data, err))
}
