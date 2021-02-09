package persist

import (
	"errors"

	mystore "types/mariadb"
)

// GetUserByUID GetUserByUID Persist
func (maria *Mariadb) GetUserByUID(uid int64) (mystore.GudpUserBase, error) {
	var _obj mystore.GudpUserBase
	err := maria.db.Table("gudp_users").Where("uid = ?", uid).First(&_obj).Error
	if err != nil {
		return _obj, errors.New("用户不存在")
	}
	return _obj, nil
}
