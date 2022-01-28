module store-user

go 1.16

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/liftbridge-io/go-liftbridge/v2 v2.2.0
	github.com/spf13/viper v1.10.1
	go.etcd.io/bbolt v1.3.6

	types v0.0.0
)

replace types => ../types
