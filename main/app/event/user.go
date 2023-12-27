/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:38:08
 * @Desc: 用户事件
 */
package event

import "selfx/app/model"

type UserCreateBefore interface {
	UserCreateBefore(*model.User) error
}

type UserCreateAfter interface {
	UserCreateAfter(*model.User)
}

type UserUpdateBefore interface {
	UserUpdateBefore(item *model.User) error
}

type UserUpdateAfter interface {
	UserUpdateAfter(item *model.User)
}

type UserDeleteBefore interface {
	UserDeleteBefore(id uint) error
}

type UserDeleteAfter interface {
	UserDeleteAfter(id uint)
}

type UserGetAfter interface {
	UserGetAfter(*model.User)
}

type UserListAfter interface {
	UserListAfter([]model.User)
}
