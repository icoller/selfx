/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-26 10:51:20
 * @Desc: 打包资源
 */
package resources

import (
	"embed"
	"io/fs"
)

var (
	//go:embed app
	App embed.FS
	//go:embed admin
	admin embed.FS
	//go:embed themes
	Themes        embed.FS
	ThemesDirName = "themes"
)

func Admin() fs.FS {
	s, _ := fs.Sub(admin, "admin")
	return s
}
