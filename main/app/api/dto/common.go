/*
 * @Author: coller
 * @Date: 2023-12-25 12:35:10
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:35:11
 * @Desc:
 */
package dto

type MessageResult struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Captcha struct {
	Base64 string `json:"base64"`
	ID     string `json:"id"`
}

func NewCaptcha(base64, id string) *Captcha {
	return &Captcha{Base64: base64, ID: id}
}
