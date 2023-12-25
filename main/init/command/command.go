/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-24 17:02:51
 * @Desc:
 */
package command

import (
	"fmt"
	"os"
	"selfx/constant"

	"github.com/spf13/pflag"
)

var (
	Addr          string // 监听地址
	ConfFilePath  string // 配置文件路径
	ShowVersion   bool   // 显示程序版本
	AdminPath     string // 后台管理路径
	AdminUsername string // 管理员用户名
	AdminPassword string // 管理员密码

)

func init() {

	pflag.StringVarP(&Addr, "addr", "a", "", "listening address.")
	pflag.StringVarP(&ConfFilePath, "config", "c", "./conf.toml", "config file path.")
	pflag.BoolVarP(&ShowVersion, "version", "v", false, "print version information and quit.")
	pflag.StringVarP(&AdminPath, "adminpath", "", "", fmt.Sprintf("reset administration path. (default %s)", constant.DefaultAdminPath))
	pflag.StringVarP(&AdminUsername, "username", "", "", "reset administrator username.")
	pflag.StringVarP(&AdminPassword, "password", "", "", "reset administrator password.")
	pflag.Parse()

	if ShowVersion {
		fmt.Println(constant.AppName + " " + constant.AppVersion)
		os.Exit(0)
	}
}
