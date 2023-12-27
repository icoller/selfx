/*
 * @Author: Coller
 * @Date: 2022-05-17 12:38:10
 * @LastEditTime: 2023-12-27 13:11:40
 * @Desc: 字符串处理
 */
package utils

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"selfx/utils/conv"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"
)

/**
 * @desc: 截取字符串
 * @param s 字符串
 * @param start 开始的位置
 * @param length 长度
 * @return {*}
 */
func CutString(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

/**
 * @desc: 随机获取字符串
 * @param l 长度
 * @return {*}
 */
func GetRandString(l int, types string) string {
	var str string
	if types == "string" {
		str = "abcdefghijklmnopqrstuvwxyz"
	} else if types == "number" {
		str = "0123456789"
	} else {
		str = "0123456789abcdefghijklmnopqrstuvwxyz"
	}
	bytes := conv.StringToByte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return conv.ByteToString(result)
}

/**
 * @desc: 根据明文密码和加盐值生成密码
 * @param password
 * @param salt 盐值
 * @return {*}
 */
func GetPassword(password string, salt string) (verify string, err error) {
	var rb []byte
	rb, err = scrypt.Key(conv.StringToByte(password), conv.StringToByte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}
	verify = hex.EncodeToString(rb)
	return
}

// 去除前后无用字符
func StringTrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

/**
 * @desc:生成随机数字
 * @param start 开始
 * @param end 结束
 * @return {*}
 */
func RandInt(start int, end int) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return strconv.Itoa(random)
}

/**
 *  @Description: 字符串补零
 *  @param str :需要操作的字符串
 *  @param resultLen 结果字符串的长度
 *  @param reverse true 为前置补零，false 为后置补零
 *  @return string
 */
func ZeroFillByStr(str string, resultLen int, reverse bool) string {
	if len(str) > resultLen || resultLen <= 0 {
		return str
	}
	if reverse {
		return fmt.Sprintf("%0*s", resultLen, str) // 不足前置补零
	}
	result := str
	for i := 0; i < resultLen-len(str); i++ {
		result += "0"
	}
	return result
}

// 去除字符串中的html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{1,}")
	src = re.ReplaceAllString(src, "")
	src = strings.Replace(src, "&nbsp;", "", 1) // 去除空格符
	return strings.TrimSpace(src)
}

/**
 * @desc: 解析grom的表名
 * @param {string} str
 * @return {*}
 */
func GetGromTag(str string) string {
	if str == "" || str == "-" {
		return ""
	}
	if strings.Contains(str, "column") {
		names := strings.Split(str, ";")
		if names[0] != "" {
			column := strings.Split(names[0], ":")
			if column[1] != "" {
				return column[1]
			}
		}
	}
	return ""
}

func GetDupList(list []string) []string {
	dupFre := make(map[string]int)
	var dep []string
	for _, item := range list {
		// 检查重复频率map中是否存在项目/元素
		_, exist := dupFre[item]
		if exist {
			//如果已经在map中，则将计数器增加1
			dupFre[item] += 1
			dep = append(dep, item)
		} else {
			//从1开始计数
			dupFre[item] = 1
		}
	}
	return dep
}

func MoneyToFormatFloat(old string) float64 {
	s := strings.Replace(old, "$", "", 1)
	s = strings.Replace(s, ",", "", -1)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

/**
 * @desc: 按字节截取字符串
 * @param {string} str
 * @param {int} length
 * @return {*}
 */
func SubStrByByte(str string, length int) string {
	if len(str) <= length {
		return str
	}
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

func ParseHideSection(str string) string {
	len := len(str)
	if len <= 3 {
		return str
	} else if len <= 9 {
		return str[:3] + "**" + str[len-3:]
	} else if len == 11 {
		return str[:3] + "****" + str[7:]
	} else if len >= 30 {
		return str[:9] + "******" + str[len-12:]
	} else {
		return str[:3] + "****" + str[len-9:]
	}
}
