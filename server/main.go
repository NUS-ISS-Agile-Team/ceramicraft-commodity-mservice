package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/config"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/grpc"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/metrics"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/mq"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-user-mservice/common/utils"
)

var sigCh = make(chan os.Signal, 1)

// @title       商品服务 API
// @version     1.0
// @description 商品微服务相关接口
// @BasePath    /product-ms/v1
func main() {
	config.Init()
	log.InitLogger()
	metrics.RegisterMetrics()
	repository.Init()
	utils.InitJwtSecret()
	mq.Init()
	go grpc.Init(sigCh)
	go http.Init(sigCh)
	// listen terminage signal
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh // Block until signal is received
	log.Logger.Infof("Received signal: %v, shutting down...", sig)
}
