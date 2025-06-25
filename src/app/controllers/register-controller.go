package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	userRepository "github.com/bytesfield/golang-gin-auth-service/src/app/repositories"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) Register(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	userRepo := userRepository.New(&user)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	userCreated, err := userRepo.SaveUser(server.DB)

	if err != nil {
		responses.ServerError(ctx, "Validation Error", err.Error())

		return
	}

	ctx.Header("Location", fmt.Sprintf("%s%s/%d", ctx.Request.Host, ctx.Request.RequestURI, userCreated.ID))

	responses.CreatedResponse(ctx, "Registration successfully", userCreated)

}
