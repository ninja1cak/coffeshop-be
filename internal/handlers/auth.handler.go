package handlers

import (
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `db:"email" form:"email"`
	Password string `db:"password" form:"password"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	userFromDB, err := h.GetAuthData(user.Email)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(userFromDB.Password, user.Password); err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: "wrong password",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(userFromDB.User_id, user.Email, userFromDB.Role)
	token, err := jwtt.Generate()
	pkg.NewResponse(200, &config.Result{
		Data: token,
	}).Send(ctx)
	return

}
