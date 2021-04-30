package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Thiti-Dev/tawb-service-v1/database"
	"github.com/Thiti-Dev/tawb-service-v1/database/collections"
	"github.com/Thiti-Dev/tawb-service-v1/helpers"
	"github.com/Thiti-Dev/tawb-service-v1/validator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//NOTE We removed the binding out and use validate only instead for validation phase
type RegisterUserBody struct {
	Username string `json:"username" xml:"username" validate:"required"`
	Password string `json:"password" xml:"password" validate:"required"`
}

func RegisterUser(c *gin.Context) error {
	registerBody := RegisterUserBody{}
	if errA := c.ShouldBind(&registerBody); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed","message":"invalid request body not a JSON format"})
	}else{

		// // ─── VALIDATE STRUCT ─────────────────────────────────────────────
		isValid, errorsData := validator.ValidateStructAndGetErrorMsg(registerBody)
		if !isValid{
			return helpers.ResponseMsg(c, 400, "Validation Errors", errorsData) 
		}
		// ─────────────────────────────────────────────────────────────────

		// ─── DB & CTX DECLARATION ─────────────────────────────────────────────
		ctx := context.Background()
		col := database.GetDatabaseInstance().Collection("User")
		// ────────────────────────────────────────────────────────────────────────────────


		// check if username is already existed or not
		q := bson.M{"username": registerBody.Username}
		result := col.FindOne(ctx,q)
		if result.Err() == nil{
			return helpers.ResponseMsg(c, 400, "This username is already existed", nil) 
		}
		// ─────────────────────────────────────────────────────────────────


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