/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 11:47:06
 * @Desc:
 */
package event

import "selfx/app/model"

type TagCreateBefore interface {
	TagCreateBefore(*model.Tag) error
}

type TagCreateAfter interface {
	TagCreateAfter(*model.Tag)
}

type TagUpdateBefore interface {
	TagUpdateBefore(item *model.Tag) error
}

type TagUpdateAfter interface {
	TagUpdateAfter(item *model.Tag)
}

type TagDeleteBefore interface {
	TagDeleteBefore(id int) error
}

type TagDeleteAfter interface {
	TagDeleteAfter(id int)
}

type TagGetAfter interface {
	TagGetAfter(*model.Tag)
}

type TagListAfter interface {
	TagListAfter([]model.Tag)
}
