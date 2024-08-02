package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"server/model"
	"server/tools"
	"time"
)

type registrationUser struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func Register(ctx *gin.Context) {
	var rUser registrationUser
	if err := ctx.ShouldBind(&rUser); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.ECode{Code: 1, Message: err.Error()})
		return
	}
	if rUser.Name == "" || rUser.Password == "" || rUser.RePassword == "" {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "账号或密码为空！"})
		return
	}
	if rUser.Password != rUser.RePassword {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "两次密码不一样！"})
		return
	}
	nameLen := len(rUser.Name)
	password := len(rUser.Password)
	if nameLen > 16 || nameLen < 8 || password > 16 || password < 8 {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "账号或密码长度不符合"})
		return
	}
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(rUser.Password) {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "账号或密码格式不符合"})
		return
	}

	if oldUser, _ := model.GetUserByName(rUser.Name); oldUser.Id > 0 {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "用户已存在!"})
		return
	}
	newUser := model.User{
		Uid:         tools.GetUID(),
		Name:        rUser.Name,
		Password:    tools.Encrypt(rUser.Password),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	if err := model.RegistrationUser(&newUser); err != nil {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "用户创建失败!"})
		return
	}
	ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "用户创建成功!"})
}
