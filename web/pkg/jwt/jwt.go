package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

type UserToken struct {
	jwt.StandardClaims
	NickName string `json:"nick_name"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
var expaire =2*time.Hour
var secreint =[]byte("你猜答案是什么")


func GetToken(nickname string)(string,error){
	user:= UserToken{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expaire).UnixNano(),
			Issuer: "web_micro_",
		},
		nickname,
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodES512,user)

	fmt.Println("anser:token")
	//token.SigningString()
	return token.SignedString(secreint)

}

func GenToken1( username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt64("auth.jwt_expire"))*time.Hour).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secreint)
}

func AuthToken(tokenString string) (*UserToken, error){
	// 解析token
	token,err := jwt.ParseWithClaims(tokenString,&UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secreint,nil
	})

	if err != nil {
		return nil,err
	}

	clasims,is_ok := token.Claims.(*UserToken)

	// 验证token
	if is_ok && token.Valid { // 正常的
		return clasims,nil
	}

	return nil,errors.New("token valid err")

}




