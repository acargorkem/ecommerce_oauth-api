package app

import (
	"github.com/acargorkem/ecommerce_oauth-api/src/http"
	"github.com/acargorkem/ecommerce_oauth-api/src/repository/db"
	"github.com/acargorkem/ecommerce_oauth-api/src/repository/rest"
	"github.com/acargorkem/ecommerce_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	usersRepository := rest.NewRepository()
	dbRepository := db.NewRepository()
	atService := access_token.NewService(usersRepository, dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
