/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 12:07:37
 * @Desc:
 */
package service

import (
	"selfx/config"
)

func CacheSize() (_ int64, err error) {
	d, err := config.Config.Cache.CurrentDriver()
	if err != nil {
		return
	}
	return d.Size()
}
