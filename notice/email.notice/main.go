package main

import (
	"email.notice/config"
	"email.notice/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()

	//初始化路由
	router.Init(r)

	// init config
	err := config.Init()
	if err != nil {

		return
	}

	//  监听端口
	err = r.Run(fmt.Sprintf(":%d", config.EmailConf.System.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

	// 优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("metis.statistics.server server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
