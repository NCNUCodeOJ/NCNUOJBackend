package router

import (
	"NCNUOJBackend/homework/view"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func getUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(jwt.ExtractClaims(c)["id"].(string))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "系統錯誤",
				"error":   err.Error(),
			})
		} else {
			c.Set("userID", uint(id))
			c.Next()
		}
	}
}

// SetupRouter index
func SetupRouter() *gin.Engine {
	/*
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
		authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
			Realm:            "NCNUOJ",
			SigningAlgorithm: "HS512",
			Key:              []byte(os.Getenv("SECRET_KEY")),
			MaxRefresh:       time.Hour,
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
	*/
	baseURL := "api/v1"
	r := gin.Default()

	prob := r.Group(baseURL + "/prob")
	//prob.Use(authMiddleware.MiddlewareFunc())
	//prob.Use(getUserID())
	{

		prob.GET("/problem/:probid", view.GetProblemByProblemId) // 用problemid 查詢 problem
		prob.GET("/tag/:tagid", view.GetProblemByTagId)          // 查這個標籤，有哪些problem有用
		prob.GET("/list_of_problem", view.GetProblemList)        //列出所有題目id
		prob.POST("/createproblem", view.CreateProblem)          // 創建題目
		prob.GET("/send", view.StartRabbitSend)                  // send
		prob.POST("/receive", view.StartRabbitReceive)           // receive
		prob.DELETE("/delete/:probid", view.DeleteProblemById)   // 用problemid 刪掉 problem
		prob.PUT("/edit/:probid", view.EditProblemByProblemId)   // 編輯更新problem
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})
	return r
}
