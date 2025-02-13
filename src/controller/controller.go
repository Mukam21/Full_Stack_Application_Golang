package controller

import (
	"full_stack_application/entity"
	"full_stack_application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TransactionService service.TransactionService
}

type Config struct {
	R                  *gin.Engine
	TransactionService service.TransactionService
}

func NewController(c *Config) {
	controller := &Controller{
		TransactionService: c.TransactionService,
	}

	apiRoutes := c.R.Group("/api")
	{
		apiRoutes.GET("/txn", controller.FindAllTransactions)
		apiRoutes.POST("/txn/add", controller.AddTransaction)
		apiRoutes.POST("/txn/edit", controller.EditTransaction)
		apiRoutes.DELETE("/txn/delete", controller.DeleteTransaction)
	}
}

func (c *Controller) FindAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

func (c *Controller) AddTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TransactionService.Add(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Added successfully"})
}

func (c *Controller) EditTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TransactionService.Edit(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Edited successfully"})
}

func (c *Controller) DeleteTransaction(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TransactionService.Delete(ctx.Request.Context(), request.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted successfully"})
}
