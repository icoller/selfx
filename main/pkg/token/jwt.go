/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2023-12-27 15:27:19
 * @Desc: 权鉴token的生成与解析
 */
package token

import (
	"errors"
	"selfx/config"
	"selfx/utils/conv"
	"selfx/utils/cryptx"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId uint   `json:"userId"`
	SaltId string `json:"saltId"`
	AppId  uint   `json:"appId"`
	jwt.RegisteredClaims
}

/**
 * @desc: 生成Token
 * @param entId 企业ID
 * @param userId 用户ID
 * @param saltId 盐值
 * @param hour 时间
 * @return {*}
 */
func Create(userId uint, saltId string, hour time.Duration) (string, error) {
	if hour == 0 {
		hour = config.Set.System.JwtExpiresTime
	}
	claimsData := CustomClaims{
		UserId: userId,
		SaltId: saltId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(hour * time.Hour)),
			Issuer:    "selfx",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claimsData)
	token, err := tokenClaims.SignedString(conv.StringToByte(config.Set.System.JwtSigningKey))
	if err != nil {
		return "", err
	}
	return strings.Replace(cryptx.AesEncrypt(token, ""), "+", "_", -1), nil
}

/**
 * @desc: 解析 token
 * @return {*}
 */
func Parse(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(cryptx.AesDecrypt(strings.Replace(tokenStr, "_", "+", -1), ""), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return conv.StringToByte(config.Set.System.JwtSigningKey), nil
	})
	if err != nil || token == nil {
		return nil, errors.New("无法处理此令牌")
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无法处理此令牌")
}
