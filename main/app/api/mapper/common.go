/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 23:22:16
 * @Desc: 映射
 */
package mapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"selfx/app/api/dto"
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/config"
	"selfx/utils"

	"github.com/go-playground/validator/v10"
)

func Result(err error) *dto.Result {
	if err == nil {
		return &dto.Result{Succ: true}
	}
	return &dto.Result{Msg: err.Error()}
}

func ResultData(data any, err error) *dto.Result {
	if err == nil {
		return &dto.Result{Succ: true, Data: data}
	}
	return &dto.Result{Msg: err.Error()}
}

func Fail(msg string) *dto.Result {
	return &dto.Result{Msg: msg}
}

func Succ(msg string) *dto.Result {
	return &dto.Result{Succ: true, Msg: msg}
}

func BodyParser(body []byte, ptr any) error {
	return json.Unmarshal(body, ptr)
}

func BodyToContext(body []byte) (ctx context.Context, err error) {
	if len(body) == 0 {
		return
	}
	err = BodyParser(body, &ctx)
	ctx.FastOffset = config.Set.More.FastOffsetMinPage > 0 && ctx.Page > config.Set.More.FastOffsetMinPage // 加速分页查询
	if ctx.Limit == 0 {
		ctx.Limit = 20 // 限制调取数量
	}
	return
}

func BodyToWhere(body []byte) (res context.Where, err error) {
	if len(body) == 0 {
		return
	}
	err = BodyParser(body, &res)
	return
}

type curdModel interface {
	model.Article | model.Category | model.Tag | model.Link | model.ArticlePost | model.Crawl
}

func BodyToCurdModel[M curdModel](body []byte) (_ *M, err error) {
	var obj M
	err = BodyParser(body, &obj)
	return &obj, err
}

func BodyToStrSet(body []byte) (res []string, err error) {
	err = BodyParser(body, &res)
	return
}

func BodyToIntSet(body []byte) (res []int, err error) {
	err = BodyParser(body, &res)
	return
}

func BodyToModelCheck[M any](body []byte) (_ *M, err error) {
	var obj M
	ref := reflect.TypeOf(&obj).Elem()
	if err = BodyParser(body, &obj); err != nil {
		return nil, errors.New("数据解析错误")
	}

	code := utils.RandInt(1000, 9999)
	err = service.Verify.Create(&model.Verify{
		Username: "coller@139.com",
		Code:     code,
		TypeId:   10,
		IP:       "192.168.1.1",
	})

	fmt.Println(err, "err")

	if err := validator.New().Struct(&obj); err != nil {
		// 如果是输入参数无效，则直接返回输入参数错误
		invalid, ok := err.(*validator.InvalidValidationError)
		if ok {
			return nil, errors.New("输入参数错误：" + invalid.Error())
		}
		validationErrs := err.(validator.ValidationErrors) // 断言是ValidationErrors
		for _, validationErr := range validationErrs {
			fieldName := validationErr.Field()      //获取是哪个字段不符合格式
			field, ok := ref.FieldByName(fieldName) //通过反射获取filed
			if ok {
				errorInfo := field.Tag.Get("info") //info tag值
				if errorInfo != "" {
					return nil, errors.New(errorInfo) //返回错误
				}
			}
			errorName := field.Tag.Get("name") //info tag值
			if errorName != "" {
				return nil, errors.New(errorName + "校验失败")
			}
			return nil, errors.New(fieldName + "校验失败")
		}
	}
	return &obj, err
}
