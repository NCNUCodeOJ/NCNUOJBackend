package router

import (
	"NCNUOJBackend/user/models"
	"NCNUOJBackend/user/view"
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// SetupRouter index
func SetupRouter() *gin.Engine {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "NCNUOJ",
		SigningAlgorithm: "HS512",
		Key:              []byte(os.Getenv("SECRET_KEY")),
		MaxRefresh:       time.Hour,
		Authenticator:    view.Login,
		TimeFunc:         time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":       v.ID,
					"username": v.UserName,
					"admin":    v.Admin,
				}
			}
			return jwt.MapClaims{}
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	baseURL := "api/v1"
	r := gin.Default()
	nor := r.Group(baseURL + "/auth")
	{
		nor.POST("/register", view.Register)
		nor.POST("/login", authMiddleware.LoginHandler)
	}
	nsr := r.Group(baseURL + "/user")
	nsr.Use(authMiddleware.MiddlewareFunc())
	{
		nsr.GET("/info", view.Hello)
	}
	auth := r.Group(baseURL + "/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})
	return r
}
