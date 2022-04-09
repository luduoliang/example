package middleware

import (
	"encoding/gob"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//初始化SESSION
func InitSession(r *gin.Engine, cookieName string) {
	//使用session
	session_cookie_name := cookieName
	store := sessions.NewCookieStore([]byte(session_cookie_name))
	store.Options(sessions.Options{
		MaxAge: 0, //30min
		Path:   "/",
	})
	r.Use(sessions.Sessions(session_cookie_name, store))
}

//获取
func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

//设置
func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	//序列化存储的值
	gob.Register(value)
	session.Set(key, value)
	session.Save()
}

//删除
func DelSession(c *gin.Context, key string) {
	session := sessions.Default(c)
	session.Delete(key)
	session.Save()
}

//清空
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
