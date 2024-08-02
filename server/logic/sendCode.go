package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/tools"
)

type data struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// SendCode 发送验证码
func SendCode(ctx *gin.Context) {
	var phone data
	if err := ctx.ShouldBind(&phone); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.ECode{Code: 1, Message: err.Error()})
		return
	}
	if !tools.CheckPhone(phone.Phone) {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "手机号格式不正确！"})
		return
	}
	//存入session
	code := tools.SmsVerify(phone.Phone)
	fmt.Println("phone code:", code)
	_ = db.SetSessionCode(ctx, phone.Phone, code)

	ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "验证码已发送！"})
}
