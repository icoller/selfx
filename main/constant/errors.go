/*
 * @Author: coller
 * @Date: 2023-12-25 11:21:04
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 13:05:17
 * @Desc:
 */
package constant

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrRecordNotFound = errors.New("record not found")

	ErrIdRequired      = errors.New("id is required")
	ErrSlugRequired    = errors.New("slug is required")
	ErrNameRequired    = errors.New("name is required")
	ErrTitleRequired   = errors.New("title is required")
	ErrContentRequired = errors.New("content is required")
	ErrUrlRequired     = errors.New("url is required")

	ErrSlugStartSpaceRequired = errors.New("slug cannot start with a space")
	ErrSlugEndSpaceRequired   = errors.New("slug cannot end with a space")

	ErrIdAlreadyExists    = errors.New("id already exists")
	ErrSlugAlreadyExists  = errors.New("slug already exists")
	ErrTitleAlreadyExists = errors.New("title already exists")

	ErrUsernameRequired = errors.New("username is required")

	ErrEmailAndMobileRequired = errors.New("mobile and email is required")

	ErrEmailRequired  = errors.New("email is required")
	ErrMobileRequired = errors.New("mobile is required")
)
