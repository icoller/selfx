/*
 * @Author: Coller
 * @Date: 2021-09-29 17:57:08
 * @LastEditTime: 2023-12-26 17:49:24
 * @Desc: 表单验证
 */
package form

import (
	"errors"
	"reflect"
	"selfx/global"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, r interface{}) (err error) {
	ref := reflect.TypeOf(r).Elem()
	if err := global.VALID.Struct(r); err != nil {
		invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
		if ok {
			return errors.New("输入参数错误：" + invalid.Error())
		}
		validationErrs := err.(validator.ValidationErrors) // 断言是ValidationErrors
		for _, validationErr := range validationErrs {
			fieldName := validationErr.Field()      //获取是哪个字段不符合格式
			field, ok := ref.FieldByName(fieldName) //通过反射获取filed
			if ok {
				errorInfo := field.Tag.Get("info") //info tag值
				if errorInfo != "" {
					return errors.New(errorInfo) //返回错误
				}
			}
			errorName := field.Tag.Get("name") //info tag值
			if errorName != "" {
				return errors.New(errorName + "校验失败")
			}
			return errors.New(fieldName + "校验失败")
		}
	}
	return nil
}
