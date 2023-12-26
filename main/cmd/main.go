/*
 * @Author: coller
 * @Date: 2023-12-20 22:04:25
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 10:48:04
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
