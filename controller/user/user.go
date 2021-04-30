package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Thiti-Dev/tawb-service-v1/database"
	"github.com/Thiti-Dev/tawb-service-v1/database/collections"
	"github.com/Thiti-Dev/tawb-service-v1/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserBody struct {
	Username string `json:"username" xml:"username" binding:"required"`
	Password string `json:"password" xml:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) error {
	registerBody := RegisterUserBody{}
	if errA := c.ShouldBind(&registerBody); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed","message":"invalid request body"})
	}else{
		ctx := context.Background()
		col := database.GetDatabaseInstance().Collection("User")
		user := collections.User{}
		user.ID = primitive.NewObjectID()
		user.Username = registerBody.Username
		user.Password = registerBody.Password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		_, err := col.InsertOne(ctx,user)
		if err != nil {
			fmt.Println(err)
			return helpers.ResponseMsg(c, 500, "Inserted data unsuccesfully", err.Error())
		}
		// ID of the inserted document.
		//objectID := result.InsertedID.(primitive.ObjectID)
		return helpers.ResponseMsg(c,200,"",user)
	}
	return nil
}