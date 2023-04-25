package controllers

import (
	"back-hw-manage/connections"
	"back-hw-manage/helpers"
	"back-hw-manage/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	*models.User
}

func Checkuser(db *gorm.DB, userName string) bool {

	var count int64
	db.Where("user_name = ?", userName).Count(&count)
	status := false
	if count > 0 {
		status = true
	}
	return status
}

func CreateUser(c *gin.Context) {

	db, err := connections.Conection_hw()

	if err != nil {
		panic(err)
	}
	User := User{}
	c.Bind(&User)

	checkUser := Checkuser(db, User.UserName)

	if !checkUser {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "UserName repeat",
		})
		return
	}
	hashPassword, err := helpers.HashPassword(User.PassWord)
	if err != nil {
		panic(err)
	}
	user := models.User{
		FirstName: User.FirstName,
		LastName:  User.LastName,
		Email:     User.Email,
		PassWord:  hashPassword,
		UserName:  User.UserName,
	}

	if result := db.Create(&user); result.Error != nil {
		panic(result.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"status":  200,
	})
}

func GetUserAll(c *gin.Context) {

}
func GetUser(c *gin.Context) {

	User := User{}

	result := connections.DB.Find(&User)

	for _, user := range result {

	}
}
