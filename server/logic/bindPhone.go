package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/model"
	"server/tools"
)

type data struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// BindPhone 调sendCode接口
func BindPhone(ctx *gin.Context) {
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

	// 获取登录态
	sessionLogin := db.GetSessionLogin(ctx)
	uid, ok2 := sessionLogin["uid"].(int64)
	fmt.Println("session uid:", uid)
	if !ok2 {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "登录已过期！"})
		return
	}

	// 获取验证码
	sessionCode := db.GetSessionCode(ctx)
	code, ok1 := sessionCode["code"].(string)
	fmt.Println("session code:", code)
	if !ok1 {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证码未发送或已过期！"})
		return
	}

	// 验证验证码是否匹配
	if pc.Code == code {
		if err := model.UpdateUserPhone(pc.Phone, uid); err != nil {
			ctx.JSON(http.StatusInternalServerError, tools.ECode{Code: 1, Message: "手机绑定失败！"})
			return
		}
		ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "手机号绑定成功！"})
	} else {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证码错误！"})
	}
}
