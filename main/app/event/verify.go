/*
 * @Author: coller
 * @Date: 2023-12-27 22:49:15
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 22:49:50
 * @Desc: 验证
 */
package event

import "selfx/app/model"

type VerifyCreateBefore interface {
	VerifyCreateBefore(*model.Verify) error
}
