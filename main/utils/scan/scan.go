/*
 * @Author: coller
 * @Date: 2023-12-25 13:23:48
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:24:02
 * @Desc:
 */
package scan

import (
	"reflect"
	"selfx/app/model"
)

// SliceToMap 切片转map
func SliceToMap[M model.ModelInterface](s []M, field string) map[any]M {
	var res = make(map[any]M)
	for _, v := range s {
		id := int(reflect.ValueOf(v).FieldByName(field).Int())
		res[id] = v
	}
	return res
}
