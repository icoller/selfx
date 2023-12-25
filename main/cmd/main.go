/*
 * @Author: coller
 * @Date: 2023-12-20 22:04:25
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 12:28:26
 * @Desc: 入口
 */
package main

import (
	"selfx/app/router"
	_ "selfx/startup"
)

func main() {

	err := router.New().Run()
	if err != nil {
		panic(err)
	}
}
