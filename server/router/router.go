package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/logic"
	"server/middleware"
	"syscall"
	"time"
)

func NewRouter() {

	r := gin.Default()

	//发送验证码接口
	r.POST("/sendCode", logic.SendCode)
	//手机号验证码登录
	r.POST("/login/sms", logic.SmsLogin)

	//用户注册
	r.POST("/registration/register", logic.Register)
	//手机号绑定
	r.POST("/registration/bindPhone", logic.BindPhone)

	// JWT路由组
	JWTGroup := r.Group("/")
	JWTGroup.Use(middleware.JWTMiddleware())
	{
		JWTGroup.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This is a protected message",
			})
		})
	}

	server := &http.Server{Addr: ":8088", Handler: r}

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, os.Kill, syscall.SIGTERM)

	// 等待接收到信号
	<-quit
	log.Println("Shutdown server...")

	// 优雅关闭HTTP服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
