package server

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/ginx"
	"github.com/goliajp/http-api-gin/env"
	log "github.com/sirupsen/logrus"
	"os"
)

func RunHttpServer(stop chan os.Signal) {
	gin.SetMode(gin.ReleaseMode)
	engine := withRouters(ginx.New())
	srv := ginx.NewServer(
		engine, &ginx.HttpConfig{
			Addr:              env.HttpListenAddr,
			Port:              env.HttpListenPort,
			ReadTimeout:       env.HttpReadTimeout,
			ReadHeaderTimeout: env.HttpReadHeaderTimeout,
			WriteTimeout:      env.HttpWriteTimeout,
			IdleTimeout:       env.HttpIdleTimeout,
		},
	)
	if err := ginx.GracefulServe(stop, srv); err != nil {
		log.Fatalf("http server graceful serve failed: %v", err)
	}
}
