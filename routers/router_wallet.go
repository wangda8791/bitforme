package routers

import (
	// "time"
	"github.com/bn_funds/api/v1"
	"github.com/bn_funds/middleware"
	"github.com/bn_funds/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init_Wallet() {

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(sessions.Sessions("my-session", store))

	// Cors: allows all origins
	router.Use(cors.Default())
	router.Use(middleware.CORSMiddleware())

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Vl

	// test := router.Group("/api/v1")
	// {
	// 	test.GET("/test", api_v1_auth.Test)
	// }

	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallet_notify", api_v1.Wallet_Notify)
	}
	// authorized := router.Group("/secure")
	// authorized.Use(middleware.AuthorizeRequest())
	// {
	// 	// authorized.GET("/", views.UserProfileView)
	// }

	// deposit := router.Group("/api/v1/deposit")
	// deposit.Use(middleware.AuthorizeRequest())
	// {
	// 	deposit.GET("/get_address", api_v1.Get_Address)
	// }

	// currencies := router.Group("/api/v1/currencies")
	// {
	// 	currencies.GET("", api_v1.Get_Currencies)
	// }

	router.Run("0.0.0.0:" + utils.GetEnv("WALLET_PORT", ""))
}
