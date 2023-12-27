/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:53:45
 * @Desc:
 */
package service

import (
	"selfx/app/service"
)

// DeleteTag 删除标签
func DeleteTag(id int) error {
	if err := service.Tag.Delete(id); err != nil {
		return err
	}
	return service.Mapping.DeleteTag(id)
}

// BatchDeleteTag 批量删除
func BatchDeleteTag(ids []int) (err error) {
	for _, id := range ids {
		if err = DeleteTag(id); err != nil {
			return
		}
	}
	return
}
