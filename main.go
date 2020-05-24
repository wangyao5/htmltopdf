package main

import (
	"flag"
	"fmt"
	"html/pdf/controller"
	"html/pdf/filter"

	"github.com/e421083458/golang_common/log"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 9000, "server port")
	flag.Parse()
	router := initRouter()
	log.Info("initialized with port(s): %d (http)", *port)

	//start server
	if err := router.Run(fmt.Sprintf(":%d", *port)); err != nil {
		log.Error("err:%s", err)
	}

}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(filter.Cors())
	router.GET("/download/pdf", controller.DownloadPdf)
	return router
}
