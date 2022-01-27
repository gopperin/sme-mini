package controller

import (
	"store-user/persist"
	store "types/mariadb"
	pb "types/pb"
)

// CreateUser CreateUser
func CreateUser(cmd pb.GudpUserCreateCommand) error {
	var obj store.GudpUser
	obj.UID = cmd.Uid
	obj.NickName = cmd.NickName
	obj.Mobile = cmd.Mobile
	obj.Email = cmd.Email
	obj.Pwd = cmd.Pwd
	obj.SecretKey = cmd.SecretKey

	return persist.GMariadb.CreateUser(obj)
}
