/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 15:32:26
 * @Desc:
 */
package startup

import (
	"os"
	"selfx/app/api/service"
	"selfx/init/command"
	"selfx/plugins"

	"github.com/gookit/color"
)

func init() {
	executeCommand()
	initPlugins()
}

func executeCommand() {
	if command.AdminPath != "" {
		if err := service.AdminPathUpdate(command.AdminPath); err != nil {
			panic(err)
		}
		color.Green.Println("admin path updated successfully\n")
	}
	if command.AdminUsername != "" {
		if err := service.AdminUsernameUpdate(command.AdminUsername); err != nil {
			panic(err)
		}
		color.Green.Println("admin username updated successfully\n")
	}
	if command.AdminPassword != "" {
		if err := service.AdminPasswordUpdate(command.AdminPassword); err != nil {
			panic(err)
		}
		color.Green.Println("admin password updated successfully\n")
	}
	if command.AdminPath != "" || command.AdminUsername != "" || command.AdminPassword != "" {
		os.Exit(0)
	}
}

func initPlugins() {
	service.PluginInit(
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
