package main

import (
	"context"
	"fmt"
	"golangstudy/jike/awesomeProject/controllers"
	"golangstudy/jike/awesomeProject/dao/mysql"
	"golangstudy/jike/awesomeProject/dao/redis"
	"golangstudy/jike/awesomeProject/logger"
	snowflake "golangstudy/jike/awesomeProject/pkg/snowflake"
	"golangstudy/jike/awesomeProject/routers"
	"golangstudy/jike/awesomeProject/setttings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"
)

//go web 脚手架模板

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	// 加载配置(配置文件加载 远程加载)
	if err := setttings.Init(); err != nil {
		fmt.Println("init settings failed", err)
		return
	}
	// 初始化日志  大型项目必须使用日志
	if err := logger.Init(setttings.Conf.LogConfig, setttings.Conf.Mode); err != nil {
		fmt.Println("init logger failed", err)
	}
	defer zap.L().Sync() //缓冲区日志 追加到日志文件
	zap.L().Debug("logger init success")
	// 初始化MySQL连接
	if err := mysql.Init(setttings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed", err)
	}
	defer mysql.Close()
	// 初始化Redis连接
	if err := redis.Init(setttings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed", err)
	}
	defer redis.Close()
	if err := snowflake.Init(setttings.Conf.StartTime, setttings.Conf.MachineID); err != nil {
		fmt.Println("init snowfalke failed", err)
		return
	}
	//初始化框架编译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Println("init validator trans failed", err)
		return
	}
	// 注册路由
	r := routers.Setup(setttings.Conf.Mode)
	// 启动服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(setttings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞z
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
