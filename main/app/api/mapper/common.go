/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 16:35:29
 * @Desc: 映射
 */
package mapper

import (
	"encoding/json"
	"selfx/app/api/dto"
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/config"
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
	ctx.FastOffset = config.Config.More.FastOffsetMinPage > 0 && ctx.Page > config.Config.More.FastOffsetMinPage // 加速分页查询
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
