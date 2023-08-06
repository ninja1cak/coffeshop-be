package handlers

import (
	"log"
	"net/http"
	"ninja1cak/coffeshop-be/internal/models"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostDataUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
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

	response, err := h.GetUser()
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

	// file, err := ctx.FormFile("file")
	// log.Println("tesssssssssssss")
	// if err != nil {
	// 	log.Println("err: ", err)
	// 	return
	// } else {
	// 	log.Println("file: ", file.Filename)

	// }

	var user models.User

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
