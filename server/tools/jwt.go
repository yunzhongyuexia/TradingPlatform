package tools

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("my_secret_key")

// Claims 定义 JWT 声明结构
type Claims struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码参数token
func GenerateToken(uid int64, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 30).Unix() //30分钟

	claims := Claims{
		Uid:  uid,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,       // 过期时间
			Issuer:    "yunzhongyuexia", //指定发行人
		},
	}
	// 该方法内部生成签名字符串，再用于获取完整、已签名的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息(进而获取其中的用户名和密码)
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // Valid()验证基于时间的声明
			return claims, nil
		}
	}
	return nil, err
}
