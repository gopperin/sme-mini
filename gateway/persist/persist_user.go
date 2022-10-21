package persist

import (
	"errors"

	mystore "github.com/gopperin/sme-mini/types/mariadb"
)

// GetUserByUID GetUserByUID Persist
func (maria *Mariadb) GetUserByUID(uid int64) (mystore.GudpUserBase, error) {
	var obj mystore.GudpUserBase
	err := maria.db.Table("gudp_users").Where("uid = ?", uid).First(&obj).Error
	if err != nil {
		return obj, errors.New("用户不存在")
	}
	return obj, nil
}
