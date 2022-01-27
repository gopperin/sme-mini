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

	var cmd pb.GudpUserCreateCommand

	c.BindJSON(&cmd)

	id, err := controller.CreateUser(cmd)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": id})
	return
}

// DeleteUser DeleteUser
func DeleteUser(c *gin.Context) {

	var cmd pb.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("DeleteUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
	return
}

// PutUser PutUser
func PutUser(c *gin.Context) {

	var cmd pb.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("PutUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
	return
}

// PatchUser PatchUser
func PatchUser(c *gin.Context) {

	var cmd pb.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("PatchUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
	return
}

// GetUserByUID GetUserByUID
func GetUserByUID(c *gin.Context) {

	param := c.Params.ByName("uid")
	uid, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": ""})
		return
	}

	obj, err := controller.GetUserByUID(uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": obj})
	return
}

// ListUser ListUser
func ListUser(c *gin.Context) {

	var cmd pb.GudpUserCreateCommand

	c.BindJSON(&cmd)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
	return
}
