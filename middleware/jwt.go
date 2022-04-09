package middleware

import (
	"errors"
	"example/config"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//jwt验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "Authorization is nil",
				"data":   nil,
			})
			c.Abort()
			return
		}
		// 按空格划分为数组，最后提取有效部分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "请求头中auth格式有误",
				"data":   nil,
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "无效的token",
				"data":   nil,
			})
			c.Abort()
			return
		}

		//接下来这一步应该去数据库里面验证操作
		//如果验证失败说明这个token无效
		//如果验证成功将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)

		// 处理请求
		c.Next()
	}
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.Cfg.JwtTokenExpriseHour)).Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),                                                                //发放时间
			Issuer:    "my-example",                                                                     // 签发人
		},
	}
	// 使用指定的签名方法(hash)创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(config.Cfg.JwtSecretKey))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.Cfg.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
