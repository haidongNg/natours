package v1

import (
	"context"
	"natours/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Router GET /users
func GetAllUsers(ctx *gin.Context) {

}

// @Summary Get user by id
// @Route GET /users/:id
func GetUser(ctx *gin.Context) {

}

type CreateUserForm struct {
	Name            string `json:"name" valid:"Required"`
	Email           string `json:"email" valid:"Required"`
	Photo           string `json:"photo"`
	Password        string `json:"password" valid:"Required"`
	PasswordConfirm string `json:"passwordConfirm" valid:"Required"`
}

// @Summary Create user
// @Route POST /users
func CreateUser(ctx *gin.Context) {
	var form CreateUserForm
	clientdb := db.GetDB()
	ctx.BindJSON(&form)
	newUser, err := clientdb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item: map[string]types.AttributeValue{
			"id":       &types.AttributeValueMemberS{Value: "A123123123BC"},
			"email":    &types.AttributeValueMemberS{Value: "ABC"},
			"photo":    &types.AttributeValueMemberS{Value: "ABC"},
			"password": &types.AttributeValueMemberS{Value: "ABC"},
		},
	})

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error",
			"errors":  err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    newUser.Attributes,
	})
}

// @Summary Create tour
// @Route PATCH /users
func UpdateUser(ctx *gin.Context) {

}

// @Summary Delete tour by id
// @Route DELETE /users/:id
func DeleteUser(ctx *gin.Context) {

}
