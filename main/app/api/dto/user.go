/*
 * @Author: coller
 * @Date: 2023-12-26 17:44:21
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 17:44:38
 * @Desc: 注册
 */
package dto

type UserRegister struct { // 注册
	Mobile   string `json:"mobile" validate:"required" info:"请输入正确的手机号"` // 手机号
	Password string `json:"password" validate:"required" info:"请输入密码"`   // 密码
	Code     string `json:"code"`                                        // 验证码
	Os       string `json:"os"`                                          // 系统
	Browser  string `json:"browser"`                                     // 浏览器
}
