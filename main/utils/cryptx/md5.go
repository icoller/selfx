/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2023-12-27 13:17:33
 * @Desc: md5
 */
package cryptx

import (
	"crypto/md5"
	"encoding/hex"
	"selfx/utils/conv"
)

/**
 * @desc: md5 encryption
 * @return {*}
 */
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write(conv.StringToByte(value))

	return hex.EncodeToString(m.Sum(nil))
}
