package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"blog/controller"
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/logger"
	"blog/pkg/snowflake"
	"blog/router"
	"blog/setting"

	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("need config field.eg: blog config.yaml")
		return
	}
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("init setting failed, err:%v", err)
		return
	}
	defer zap.L().Sync()
	
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v", err)
		return
	}
	defer zap.L().Sync()

	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
	}

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init InitTrans failed, err:%v\n", err)
		return
	}

	r := routes.Setup(setting.Conf.Mode)
	fmt.Println(setting.Conf.Port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}