package persist

import (
	store "github.com/gopperin/sme-mini/types/mariadb"
)

// CreateUser CreateUser Persist
func (maria *Mariadb) CreateUser(obj store.GudpUser) error {
	return maria.db.Create(&obj).Error
}
