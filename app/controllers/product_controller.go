package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/database"
	"github.com/mstomar698/GoMall/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

var productCollection *mongo.Collection = database.GetCollection(database.DB, "users")

// create product controller
func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product models.CreateProduct
		defer cancel()

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		resultInsertionNumber, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr != nil {
			msg := fmt.Sprintf("product item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
