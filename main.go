package main

import (
	"adamnasrudin03/challenge-transaction/app"
	"adamnasrudin03/challenge-transaction/app/configs"
	"adamnasrudin03/challenge-transaction/app/controller"
	routers "adamnasrudin03/challenge-transaction/app/router"
	"adamnasrudin03/challenge-transaction/pkg/database"
	"adamnasrudin03/challenge-transaction/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDbConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	transactionController controller.TransactionController = controller.NewTransactionController(services)
)

func main() {
	defer database.CloseDbConnection(db)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("welcome its server", http.StatusOK, false, nil, nil))
	})

	// Route here
	routers.TransactionRouter(router, transactionController)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("page not found", http.StatusNotFound, true, nil, nil))
	})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
