package controller

import (
	"store-user/persist"
	store "types/mariadb"
	pb "types/pb"
)

// CreateUser CreateUser
func CreateUser(cmd pb.GudpUserCreateCommand) error {
	var _obj store.GudpUser
	_obj.UID = cmd.Uid
	_obj.NickName = cmd.NickName
	_obj.Mobile = cmd.Mobile
	_obj.Email = cmd.Email
	_obj.Pwd = cmd.Pwd
	_obj.SecretKey = cmd.SecretKey

	return persist.GMariadb.CreateUser(_obj)
}
