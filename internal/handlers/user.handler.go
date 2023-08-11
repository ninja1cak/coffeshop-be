package handlers

import (
	"log"
	"net/http"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostDataUser(ctx *gin.Context) {
	var user = models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}
	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)

	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	response, err := h.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return

	} else {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Created",
			"data":    response,
		})
	}

}

func (h *HandlerUser) GetDataUser(ctx *gin.Context) {
	// var user models.User

	// if err := ctx.ShouldBind(&user); err != nil {

	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	response, err := h.GetUser("")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return

	} else {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Ok",
			"data":    response,
		})
	}

}

func (h *HandlerUser) UpdateDataUser(ctx *gin.Context) {

	var user models.User
	user.Photo_profile = ctx.MustGet("image").(*string)
	log.Println("handler", user.Photo_profile)

	user.Email = ctx.Query("email")

	if err := ctx.ShouldBind(&user); err != nil {
		log.Println("tes:", err)
		return
	}

	response, err := h.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":  201,
			"message": "Ok",
			"data":    response,
		})
	}

}

func (h *HandlerUser) DeleteDataUser(ctx *gin.Context) {
	var user models.User
	user.Email = ctx.Query("email")
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	response, err := h.DeleteUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return

	} else {
		ctx.JSON(200, gin.H{
			"status":  201,
			"message": "Ok",
			"data":    response,
		})
	}

}

func (h *HandlerUser) GetDataUserLogin(ctx *gin.Context) {

	var user models.User

	user.User_id = ctx.MustGet("user_id").(string)

	response, err := h.GetUser(user.User_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return

	} else {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Ok",
			"data":    response,
		})
	}

}
