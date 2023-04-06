package router

import (
	"adamnasrudin03/challenge-transaction/app/controller"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(e *gin.Engine, transactionController controller.TransactionController) {
	transactionRoutes := e.Group("/api/v1/transaction")
	{
		transactionRoutes.GET("/", transactionController.ListTransaction)
	}
}
