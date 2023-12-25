/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 11:46:47
 * @Desc:
 */
package event

import "selfx/app/model"

type CategoryCreateBefore interface {
	CategoryCreateBefore(*model.Category) error
}

type CategoryCreateAfter interface {
	CategoryCreateAfter(*model.Category)
}

type CategoryUpdateBefore interface {
	CategoryUpdateBefore(item *model.Category) error
}

type CategoryUpdateAfter interface {
	CategoryUpdateAfter(item *model.Category)
}

type CategoryDeleteBefore interface {
	CategoryDeleteBefore(id int) error
}

type CategoryDeleteAfter interface {
	CategoryDeleteAfter(id int)
}

type CategoryGetAfter interface {
	CategoryGetAfter(*model.Category)
}

type CategoryListAfter interface {
	CategoryListAfter([]model.Category)
}
