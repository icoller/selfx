/*
 * @Author: coller
 * @Date: 2023-12-25 13:24:47
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:26:56
 * @Desc:
 */
package seq

import (
	"selfx/app/model"
	"selfx/utils/scan"
)

// SortByIds 根据ids排序
func SortByIds[M model.ModelInterface](items []M, ids []int) (res []M) {
	if len(items) < 2 || len(ids) < 2 {
		return items
	}
	m := scan.SliceToMap(items, "ID")
	for _, id := range ids {
		item, ok := m[id]
		if !ok {
			continue
		}
		res = append(res, item)
	}
	return
}
