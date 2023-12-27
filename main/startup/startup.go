/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 11:57:37
 * @Desc:
 */
package startup

import (
	"os"
	pluginServ "selfx/app/plugin/service"
	"selfx/app/service"
	"selfx/init/command"
	"selfx/plugins"

	"github.com/gookit/color"
)

func init() {
	executeCommand()
	initPlugins()
}

func executeCommand() {
	var service *service.AdminService
	if command.AdminPath != "" {
		if err := service.PathUpdate(command.AdminPath); err != nil {
			panic(err)
		}
		color.Green.Println("admin path updated successfully\n")
	}
	if command.AdminUsername != "" {
		if err := service.UsernameUpdate(command.AdminUsername); err != nil {
			panic(err)
		}
		color.Green.Println("admin username updated successfully\n")
	}
	if command.AdminPassword != "" {
		if err := service.PasswordUpdate(command.AdminPassword); err != nil {
			panic(err)
		}
		color.Green.Println("admin password updated successfully\n")
	}
	if command.AdminPath != "" || command.AdminUsername != "" || command.AdminPassword != "" {
		os.Exit(0)
	}
}

func initPlugins() {
	pluginServ.PluginInit(
		plugins.NewGenerateSlug(),
		plugins.NewArticleSanitizer(),
		plugins.NewSaveArticleImages(),
		plugins.NewDetectLinks(),
		plugins.NewGenerateDescription(),
		plugins.NewPreBuildArticleCache(),
		plugins.NewPushToBaidu(),
		plugins.NewMakeCarousel(),
		plugins.NewPostCrawl(),
		plugins.NewDidiAuto(),
	)

}
