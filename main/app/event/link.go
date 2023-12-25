/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 11:47:01
 * @Desc:
 */
package event

import "selfx/app/model"

type LinkCreateBefore interface {
	LinkCreateBefore(*model.Link) error
}

type LinkCreateAfter interface {
	LinkCreateAfter(*model.Link)
}

type LinkUpdateBefore interface {
	LinkUpdateBefore(item *model.Link) error
}

type LinkUpdateAfter interface {
	LinkUpdateAfter(item *model.Link)
}

type LinkDeleteBefore interface {
	LinkDeleteBefore(id int) error
}

type LinkDeleteAfter interface {
	LinkDeleteAfter(id int)
}

type LinkGetAfter interface {
	LinkGetAfter(*model.Link)
}

type LinkListAfter interface {
	LinkListAfter([]model.Link)
}
