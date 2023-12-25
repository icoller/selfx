/*
 * @Author: coller
 * @Date: 2023-12-25 13:16:27
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 15:32:11
 * @Desc:
 */
package model

type ModelInterface interface {
	Article | Category | Tag | Link | Crawl
}
