package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// DeleteUser DeleteUser
func DeleteUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	fmt.Println("DeleteUser", _cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _cmd.Uid})
	return
}

// PutUser PutUser
func PutUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	fmt.Println("PutUser", _cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _cmd.Uid})
	return
}

// PatchUser PatchUser
func PatchUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	fmt.Println("PatchUser", _cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _cmd.Uid})
	return
}

// GetUserByUID GetUserByUID
func GetUserByUID(c *gin.Context) {

	_param := c.Params.ByName("uid")
	_uid, err := strconv.ParseInt(_param, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": ""})
		return
	}

	_obj, err := controller.GetUserByUID(_uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _obj})
	return
}

// ListUser ListUser
func ListUser(c *gin.Context) {

	var _cmd pb.GudpUserCreateCommand

	c.BindJSON(&_cmd)

	fmt.Println("ListUser", _cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": _cmd.Uid})
	return
}
