package persist

import (
	store "types/mariadb"
)

// CreateUser CreateUser Persist
func (maria *Mariadb) CreateUser(obj store.GudpUser) error {
	return maria.db.Create(&obj).Error
}
