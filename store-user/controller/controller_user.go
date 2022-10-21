package controller

import (
	"fmt"

	"github.com/gopperin/sme-mini/store-user/persist"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
)

// CreateUser CreateUser
func CreateUser(obj mystore.GudpUser) error {
	err := persist.GMariadb.CreateUser(obj)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
