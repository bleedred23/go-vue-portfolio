package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-tracker/entity"
	"portfolio-tracker/service"
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
	c.R.Use(CORSMiddleware())

	apiRoutes := c.R.Group("/api")
	{
		apiRoutes.GET("/txn/:id", controller.FindTransactionById)
		apiRoutes.GET("/txn", controller.FindAllTransactions)
		apiRoutes.POST("/txn/add", controller.AddTransaction)
		apiRoutes.POST("/txn/edit", controller.EditTransaction)
		apiRoutes.DELETE("/txn/delete", controller.DeleteTransaction)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (c *Controller) FindTransactionById(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
func (c *Controller) FindAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Updated successfully"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted successfully"})
}
