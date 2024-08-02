package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/model"
	"server/tools"
)

// SmsLogin 验证码登录
func SmsLogin(ctx *gin.Context) {
	var pc data
	if err := ctx.ShouldBind(&pc); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.ECode{Code: 1, Message: err.Error()})
		return
	}

	// 验证手机号格式
	if !tools.CheckPhone(pc.Phone) {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "手机号格式不正确！"})
		return
	}

	// 从缓存中获取验证码
	sessionCode := db.GetSessionCode(ctx)
	code, ok := sessionCode["code"].(string)
	fmt.Println("session code:", code)
	if !ok {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证码未发送或已过期！"})
		return
	}

	// 验证验证码是否匹配
	if pc.Code == code {
		user := model.GetUserByPhone(pc.Phone)
		//设置Session到Redis
		_ = db.SetSessionLogin(ctx, user.Name, user.Uid)
		ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "登录成功！"})
	} else {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证码错误！"})
	}
}
