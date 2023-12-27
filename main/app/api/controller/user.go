/*
 * @Author: coller
 * @Date: 2023-12-25 12:30:40
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 15:49:41
 * @Desc: 管理元
 */
package controller

import (
	"selfx/app/api/dto"
	"selfx/app/api/mapper"
	apiServ "selfx/app/api/service"
	"selfx/app/model"
	"selfx/app/service"
	"selfx/constant"
	"selfx/init/ip2region"
	"selfx/utils/agent"

	"github.com/gofiber/fiber/v2"
)

// 注册用户
func UserRegister(c *fiber.Ctx) error {
	req, err := mapper.BodyToModelCheck[dto.UserRegister](c.Body())
	if err != nil {
		return c.JSON(mapper.Result(err))
	}
	if err := service.Verify.CheckUsernameCode(req.Username, req.Code, constant.VerifyTypeIdRegister); err != nil {
		return c.JSON(mapper.Result(err))
	}
	// // 走注册流程
	user, err := apiServ.UserRegister(req)
	if err != nil {
		return c.JSON(mapper.Result(err))
	}
	var loginReq dto.UserLogin
	ip := agent.GetIp(c)
	userInfo, err := apiServ.UserLogin(&dto.UserLogin{Ip: ip, Username: req.Username, Password: req.Password, Mode: constant.UserModeRegister})
	if err != nil {
		return c.JSON(mapper.Result(err))
	}
	userAgent := string(c.Context().UserAgent())
	service.Login.RecordCreate(&model.LoginRecord{UserId: user.ID, Ip: loginReq.Ip, Mode: req.Mode, Region: ip2region.Region(ip), Browser: agent.GetBrowser(userAgent), Os: agent.GetOs(userAgent), Remark: userAgent})
	return c.JSON(mapper.ResultData(userInfo, nil))
}
