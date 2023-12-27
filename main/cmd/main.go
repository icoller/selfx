/*
 * @Author: coller
 * @Date: 2023-12-20 22:04:25
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 20:58:01
 * @Desc: 入口
 */
package main

import (
	"selfx/app/router"
	_ "selfx/startup"
)

func main() {
	if err := router.New().Run(); err != nil {
		panic(err)
	}
}
