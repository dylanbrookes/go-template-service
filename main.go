package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/template-service/pkg/controllers"
	"github.com/template-service/pkg/middlewares"
	"github.com/template-service/pkg/utils"
)

var port string

func init() {
	fmt.Println("--- Template Service ---")
	utils.LoadEnv()
	port = os.Getenv("PORT")
	fmt.Println("Port configured to: ", port)
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/", "/live"),
		gin.Recovery(),
	)

	api := router.Group("/api", middlewares.Authenticate())
	{
		v1 := api.Group("/v1")
		example := v1.Group("/example")
		{
			example.GET("/", controllers.GetExample)
		}

	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("failed to start server! ", err)
	}

}
