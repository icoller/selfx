/*
 * @Author: coller
 * @Date: 2023-12-25 13:53:45
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 13:56:12
 * @Desc:
 */
package config

import (
	"selfx/config/aggregate"
	"selfx/config/service"
)

var Config = aggregate.NewConfig()

func init() {
	for _, item := range Config.Items() {
		if err := service.Sync(item); err != nil {
			panic(err)
		}
	}
}
