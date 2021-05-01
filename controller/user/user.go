package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Thiti-Dev/tawb-service-v1/database"
	"github.com/Thiti-Dev/tawb-service-v1/database/collections"
	"github.com/Thiti-Dev/tawb-service-v1/helpers"
	models "github.com/Thiti-Dev/tawb-service-v1/models/jwt"
	"github.com/Thiti-Dev/tawb-service-v1/models/user/user_request_bodies"
	"github.com/Thiti-Dev/tawb-service-v1/packages/encryptor"
	"github.com/Thiti-Dev/tawb-service-v1/packages/jwt"
	"github.com/Thiti-Dev/tawb-service-v1/validator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c *gin.Context) error {
	registerBody := user_request_bodies.RegisterUserBody{}
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

		//hashing
		hashPwd, _ := encryptor.HashPassword(registerBody.Password)
		// ─────────────────────────────────────────────────────────────────

		user := collections.User{}
		user.ID = primitive.NewObjectID()
		user.Username = registerBody.Username
		user.Password = hashPwd
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

func SignInWithCredential(c *gin.Context) error {
	credential := new(user_request_bodies.LoginUserBody)
	if err := c.ShouldBind(credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed","message":"invalid request body not a JSON format"})
		return nil
	}
	if validator.ValidateStructAndSendResponseMsgIfThereIsError(c,credential){
		return errors.New("Validation errors")
	}

	// ─── DB & CTX DECLARATION ─────────────────────────────────────────────
	ctx := context.Background()
	col := database.GetDatabaseInstance().Collection("User")
	// ────────────────────────────────────────────────────────────────────────────────

	// check if username is already existed or not
	q := bson.M{"username": credential.Username}
	result := col.FindOne(ctx,q)
	if result.Err() != nil{
		return helpers.ResponseMsg(c, 400, "email or password is incorrect", nil) 
	}
	// ─────────────────────────────────────────────────────────────────

	userData := new(collections.User)
	result.Decode(userData)

	// checking if password are the same between plain&crypted
	if !encryptor.CheckPasswordHash(credential.Password,userData.Password) {
		return helpers.ResponseMsg(c, 400, "email or password is incorrect", nil) // doesnt give any hint xD
	}
	// ─── SIGN THE TOKEN FOR USER ────────────────────────────────────────────────────
	signedToken := jwt.GetSignedTokenFromData(models.RequiredDataToClaims{
		Username: userData.Username,
		ID: userData.ID,
	})
	// ────────────────────────────────────────────────────────────────────────────────

	c.JSON(200,gin.H{"status":"success","accessToken":signedToken})

	return nil
}

func GetUserCredential(c *gin.Context) error {
	_, userData := helpers.GetUserDataFromContext(c)
	c.JSON(http.StatusOK,gin.H{"success":true,"data":userData})
	return nil
}