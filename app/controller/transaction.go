package controller

import (
	"adamnasrudin03/challenge-transaction/app/dto"
	"adamnasrudin03/challenge-transaction/app/service"
	"adamnasrudin03/challenge-transaction/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TransactionController interface is a contract what this controller can do
type TransactionController interface {
	ListTransaction(ctx *gin.Context)
}

type transactionController struct {
	Service *service.Services
}

// NewTransactionController creates a new instance of TransactionController
func NewTransactionController(srv *service.Services) TransactionController {
	return &transactionController{
		Service: srv,
	}
}

func (c *transactionController) ListTransaction(ctx *gin.Context) {
	var (
		paramPage  uint64 = 0
		paramLimit uint64 = 10
	)

	paramPage, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse("query param page not found or invalid", http.StatusBadRequest, true, nil))
		return
	}

	paramLimit, err = strconv.ParseUint(ctx.Query("limit"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse("query param limit not found or invalid", http.StatusBadRequest, true, nil))
		return
	}
	param := dto.ParamTransactions{
		Page:  paramPage,
		Limit: paramLimit,
	}

	transactions, err := c.Service.Transaction.GetTransactions(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse("error to get transaction", http.StatusBadRequest, true, nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponse("List of transaction", http.StatusOK, false, transactions))
}
