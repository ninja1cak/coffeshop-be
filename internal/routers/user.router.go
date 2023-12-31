package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/user")

	//dependency injection
	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	router.POST("/", handler.PostDataUser)
	router.PATCH("/", handler.UpdateDataUser)
	router.DELETE("/", handler.DeleteDataUser)
	router.GET("/", handler.GetDataUser)

}
