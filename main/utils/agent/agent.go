/*
 * @Author: Coller
 * @Date: 2022-01-04 20:15:51
 * @LastEditTime: 2023-12-27 14:10:40
 * @Desc: 头处理
 */
package agent

import (
	"fmt"
	"regexp"
	"selfx/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetIp(ctx *fiber.Ctx) (ip string) {
	for _, v := range config.Set.Router.ProxyHeader {
		if ip = ctx.Get(v); ip != "" {
			arr := strings.Split(ip, ",")
			if len(arr) == 0 {
				continue
			}
			if arr[0] != "" {
				return arr[0]
			}
		}
	}
	return ctx.IP()
}

// 获取OS
func GetOs(userAgent string) string {
	osName := "Unknown"
	if userAgent == "" {
		return osName
	}

	strRe, _ := regexp.Compile("(?i:\\((.*?)\\))")
	userAgent = strRe.FindString(userAgent)

	levelNames := ":micromessenger:dart:Windows NT:Windows Mobile:Windows Phone:Windows Phone OS:Macintosh|Macintosh:Mac OS:CrOS|CrOS:iPhone OS:iPad|iPad:OS:Android:Linux:blackberry:hpwOS:Series:Symbian:PalmOS:SymbianOS:J2ME:Sailfish:Bada:MeeGo:webOS|hpwOS:Maemo:"
	var regStrArr []string
	namesArr := strings.Split(strings.Trim(levelNames, ":"), ":")
	for _, name := range namesArr {
		regStrArr = append(regStrArr, fmt.Sprintf("(%s[\\s?\\/XxSs0-9_.]+)", name))
	}
	regexpStr := fmt.Sprintf("(?i:%s)", strings.Join(regStrArr, "|"))
	nameRe, _ := regexp.Compile(regexpStr)

	names := nameRe.FindAllString(userAgent, -1)
	name := ""
	for _, s := range names {
		if name == "" {
			name = strings.TrimSpace(s)
		} else if len(name) > 0 {
			if strings.Contains(name, "Macintosh") && s != "" {
				name = strings.TrimSpace(s)
			} else if strings.Contains(name, s) {
				name = strings.TrimSpace(s)
			} else if !strings.Contains(s, name) {
				if strings.Contains(name, "iPhone") ||
					strings.Contains(name, "iPad") {
					s = strings.Trim(s, "Mac OS X")
				}

				if s != "" {
					name += " " + strings.TrimSpace(s)
				}
			}
			break
		}

		// if strings.Contains(name, "Windows NT") {
		// 	name = getWinOsNameWithWinNT(name)
		// 	break
		// }
	}

	if name != "" {
		osName = name
	}

	return osName
}

// 获取浏览器
func GetBrowser(userAgent string) string {
	deviceName := "Unknown"
	levelNames := ":VivoBrowser:QQDownload:QQBrowser:QQ:MQQBrowser:MicroMessenger:TencentTraveler:LBBROWSER:TaoBrowser:BrowserNG:UCWEB:TwonkyBeamBrowser:NokiaBrowser:OviBrowser:NF-Browser:OneBrowser:Obigo:DiigoBrowser:baidubrowser:baiduboxapp:xiaomi:Redmi:MI:Lumia:Micromax:MSIEMobile:IEMobile:EdgiOS:Yandex:Mercury:Openwave:TouchPad:UBrowser:Presto:Maxthon:MetaSr:Trident:Opera:IEMobile:Edge:Chrome:Chromium:OPR:CriOS:Firefox:FxiOS:fennec:CrMo:Safari:Nexus One:Nexus S:Nexus:Blazer:teashark:bolt:HTC:Dell:Motorola:Samsung:LG:Sony:SonyST:SonyLT:SonyEricsson:Asus:Palm:Vertu:Pantech:Fly:Wiko:i-mobile:Alcatel:Nintendo:Amoi:INQ:ONEPLUS:Tapatalk:PDA:Novarra-Vision:NetFront:Minimo:FlyFlow:Dolfin:Nokia:Series:AppleWebKit:Mobile:Mozilla:Version:"

	var regStrArr []string
	namesArr := strings.Split(strings.Trim(levelNames, ":"), ":")
	for _, name := range namesArr {
		regStrArr = append(regStrArr, fmt.Sprintf("(%s[\\s?\\/0-9.]+)", name))
	}
	regexpStr := fmt.Sprintf("(?i:%s)", strings.Join(regStrArr, "|"))
	nameRe, _ := regexp.Compile(regexpStr)
	names := nameRe.FindAllString(userAgent, -1)

	level := 0
	for _, name := range names {
		replaceRe, _ := regexp.Compile("(?i:[\\s?\\/0-9.]+)")
		n := replaceRe.ReplaceAllString(name, "")
		l := strings.Index(levelNames, fmt.Sprintf(":%s:", n))
		if level == 0 {
			deviceName = strings.TrimSpace(name)
		}

		if l >= 0 && (level == 0 || level > l) {
			level = l
			deviceName = strings.TrimSpace(name)
		}
	}
	return deviceName
}
