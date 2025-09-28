package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/config"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/grpc"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository"
)

var sigCh = make(chan os.Signal, 1)

// @title       商品服务 API
// @version     1.0
// @description 商品微服务相关接口
// @host        localhost:8080
// @BasePath    /product-ms/v1
func main() {
	config.Init()
	log.InitLogger()
	repository.Init()
	go grpc.Init(sigCh)
	go http.Init(sigCh)
	// listen terminage signal
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh // Block until signal is received
	log.Logger.Infof("Received signal: %v, shutting down...", sig)
}
