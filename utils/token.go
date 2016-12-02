package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
)

var SecretKey string = "hello world"

func CreateToken(userid string, cusId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userid,
		"cusId":  cusId,
		"exp":    time.Now().Add(time.Minute * 60).Unix(),
	})
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

func ParseToken(tokenString string) (bool, string, string) {
	if tokenString == "" {
		return false, "", ""
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		logs.Error(err.Error())
		return false, "", ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		vlu, isOk := claims["exp"].(float64)

		if !isOk {
			return false, "", ""
		}

		if int64(vlu) < time.Now().Unix() {
			return false, "", ""
		}

		return true, fmt.Sprintf("%s", claims["userid"]), fmt.Sprintf("%s", claims["cusId"])
		//		value, ok := claims["userid"].(float64)
		//		if !ok {
		//			return false, -2
		//		}

		//		return true, int64(value)
	} else {
		if err != nil {
			fmt.Println(err.Error())
		}
		return false, "", ""
	}
}
