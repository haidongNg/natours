package v1

import (
	"context"
	"natours/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

// @Summary Get all tours
// @Router GET /tours
func GetAllTours(ctx *gin.Context) {
	clientdb := db.GetDB()
	tours, err := clientdb.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("tours"),
	})

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error",
			"data":    nil,
		})
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    tours.Items,
	})
}

// @Summary Get tour by id
// @Route GET /tours/:id
func GetTour(ctx *gin.Context) {
}

// @Summary Create tour
// @Route POST /tours
func CreateTour(ctx *gin.Context) {

}

// @Summary Create tour
// @Route PATCH /tours
func UpdateTour(ctx *gin.Context) {

}

// @Summary Delete tour by Iid
// @Route DELETE /tours/:id
func DeleteTour(ctx *gin.Context) {

}
