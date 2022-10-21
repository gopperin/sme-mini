package controller

import (
	"github.com/gopperin/sme-mini/store-user/persist"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
)

// CreateUser CreateUser
func CreateUser(obj mystore.GudpUser) error {
	return persist.GMariadb.CreateUser(obj)
}
