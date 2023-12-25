/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:16:39
 * @Desc:
 */
package service

import (
	"selfx/app/dto"
	"selfx/app/model"
	"selfx/app/repo/context"
	"selfx/app/service"
	appService "selfx/app/service"
)

// CategoryTree 获取所有的分类树
func CategoryTree() ([]dto.CategoryTree, error) {
	items, err := service.Category.List(context.NewContext(1000, "")) // 限制最大值
	if err != nil {
		return nil, err
	}
	return appService.MakeCategoryTree(appService.CategoryModelListToCategoryTreeList(items), 0), nil
}

// CategoryGetOrCreate 获取或创建类目
func CategoryGetOrCreate(name string) (*model.Category, error) {
	res, err := service.Category.GetOrCreate(name)
	if err != nil { // 如果出现错误，有可能是并发请求造成的，再次尝试获取就可以得到正确的结果
		res, err = service.Category.GetOrCreate(name)
	}
	return res, err
}
