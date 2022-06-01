package controller

import (
	"github.com/gin-gonic/gin"
	"web-base/dao"
	"web-base/model"
	"web-base/utils"
)

//register api
func Register(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	//create user
	err := dao.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

//login api
func Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	//get user from db
	raw_user, err := dao.GetUser(user.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	//check password
	if raw_user.Password != user.Password {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "password error",
		})
		return
	}
	//create token
	token, err := utils.CreateToken(raw_user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}

//get userinfo api
func GetUserInfo(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user, err := dao.GetUserByID(uint(user_id))
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": user,
	})
}
