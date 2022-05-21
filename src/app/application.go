package app

import (
	"github.com/acargorkem/ecommerce_oauth-api/src/clients/cassandra"
	"github.com/acargorkem/ecommerce_oauth-api/src/domain/access_token"
	"github.com/acargorkem/ecommerce_oauth-api/src/http"
	"github.com/acargorkem/ecommerce_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
