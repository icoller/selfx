/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:11:09
 * @Desc:
 */
package controller

import (
	"errors"
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
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
		data = apiServ.SystemLoadPercent()
	case "systemCPU":
		data, err = apiServ.SystemCPUPercent(time.Second)
	case "systemMemory":
		data, err = apiServ.SystemMemoryPercent()
	case "systemDisk":
		data, err = apiServ.SystemDiskPercents()
	case "appCPU":
		data, err = apiServ.AppCPUPercent()
	case "appMemory":
		data, err = apiServ.AppUsedMemory()
	case "appInfo":
		data = apiServ.AppInfo()
	case "database":
		data = db.GetSize()
	case "log":
		data, err = apiServ.LogDirSize()
	case "cache":
		data, err = apiServ.CacheSize()
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
	case "crawlTotal":
		data, err = service.Crawl.CountTotal()
	case "crawlToday":
		data, err = service.Crawl.CountToday()
	case "crawlYesterday":
		data, err = service.Crawl.CountYesterday()
	case "categoryTotal":
		data, err = service.Category.CountTotal()
	case "tagTotal":
		data, err = service.Tag.CountTotal()
	case "linkTotal":
		data, err = service.Link.CountTotal()
	default:
		return ctx.JSON(mapper.Result(errors.New("id is undefined")))
	}
	return ctx.JSON(mapper.ResultData(data, err))
}
