package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gateway/controller"
	"types/pb"
)

// CreateUser CreateUser
func CreateUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	_id, err := controller.CreateUser(_cmd)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": _id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _id})
	return
}

// GetUser GetUser
func GetUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	_obj, err := controller.GetUserByUID(_cmd.Uid)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _obj})
	return
}
