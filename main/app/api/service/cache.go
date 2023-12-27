/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 17:14:17
 * @Desc:
 */
package service

import (
	"selfx/config"
)

func CacheSize() (_ int64, err error) {
	d, err := config.Set.Cache.CurrentDriver()
	if err != nil {
		return
	}
	return d.Size()
}
