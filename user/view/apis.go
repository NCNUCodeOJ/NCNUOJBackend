package view

import (
	"NCNUOJBackend/user/models"
	"NCNUOJBackend/user/pkg"

	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// StoreData default jwt data
type StoreData struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Admin    bool   `json:"admin"`
}

// Hello hello test
func Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": claims["username"],
		"admin":    claims["admin"],
	})
}

// Login login
func Login(c *gin.Context) (interface{}, error) {
	d := struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.BindJSON(&d); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	u, err := models.UserDetailByUserName(d.Name)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	if pkg.Compare(u.Password, d.Password) == nil {
		return &u, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// Register 註冊
func Register(c *gin.Context) {
	data := struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}{}
	var user models.User
	c.BindJSON(&data)
	if u, err := models.UserDetailByUserName(data.Name); err == nil && u.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":  "此 username 已被使用",
			"username": data.Name,
		})
		return
	}
	pwd, err := pkg.Encrypt(data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	user.UserName = data.Name
	user.Password = pwd
	models.AddUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增成功",
	})
}
