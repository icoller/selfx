/*
 * @Author: coller
 * @Date: 2023-12-25 13:53:45
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 17:12:46
 * @Desc:
 */
package config

import (
	"selfx/config/aggregate"
	"selfx/config/service"
)

var Set = aggregate.NewConfig()

func init() {
	for _, item := range Set.Items() {
		if err := service.Sync(item); err != nil {
			panic(err)
		}
	}
}
