package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gopperin/sme-mini/gateway/controller"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
	"github.com/gopperin/sme-mini/types/proto"
)

type GudpUserCreate struct {
	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	NickName  string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Mobile    string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Pwd       string `protobuf:"bytes,5,opt,name=pwd,proto3" json:"pwd,omitempty"`
	SecretKey string `protobuf:"bytes,6,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

// CreateUser CreateUser
func CreateUser(c *gin.Context) {

	var cmd mystore.GudpUserBase

	c.BindJSON(&cmd)

	id, err := controller.CreateUser(cmd)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{"flag": 2, "msg": controller.GetLangContent("", "", err.Error()), "data": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": id})
}

// DeleteUser DeleteUser
func DeleteUser(c *gin.Context) {

	var cmd proto.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("DeleteUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
}

// PutUser PutUser
func PutUser(c *gin.Context) {

	var cmd proto.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("PutUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
}

// PatchUser PatchUser
func PatchUser(c *gin.Context) {

	var cmd proto.GudpUserCreateCommand

	c.BindJSON(&cmd)

	fmt.Println("PatchUser", cmd.Uid)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
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
}

// ListUser ListUser
func ListUser(c *gin.Context) {

	var cmd proto.GudpUserCreateCommand

	c.BindJSON(&cmd)

	c.JSON(http.StatusOK, gin.H{"flag": 1, "msg": controller.GetLangContent("", "", "成功"), "data": cmd.Uid})
}
