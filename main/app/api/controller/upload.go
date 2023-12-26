/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:32:01
 * @Desc:
 */
package controller

import (
	"mime/multipart"
	"selfx/app/api/mapper"
	"selfx/init/storage"
	"selfx/init/upload"

	"github.com/gofiber/fiber/v2"
)

func UploadInit(ctx *fiber.Ctx) error {
	if err := upload.Init(); err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	return ctx.JSON(mapper.Result(nil))
}

func Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(mapper.Result(err))
	}
	var urls []string
	for _, header := range form.File {
		for _, v := range header {
			res, err := uploadFunc(v)
			if err != nil {
				return ctx.JSON(mapper.Result(err))
			}
			urls = append(urls, res.URL)
		}
	}
	return ctx.JSON(mapper.ResultData(urls, nil))
}

func uploadFunc(h *multipart.FileHeader) (res *upload.Result, err error) {
	f, err := h.Open()
	if err != nil {
		return
	}
	defer f.Close()
	return upload.Upload(h.Filename, "", storage.NewSetValue(f))
}
