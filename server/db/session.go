package db

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var sessionCode = "session-code"

func SetSessionCode(ctx *gin.Context, phone string, code string) error {
	store := NewRedisStoreCode(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionCode)
	session.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	session.Values["phone"] = phone
	session.Values["code"] = code
	return session.Save(ctx.Request, ctx.Writer)
}

func GetSessionCode(ctx *gin.Context) map[interface{}]interface{} {
	store := NewRedisStoreCode(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionCode)
	return session.Values
}
