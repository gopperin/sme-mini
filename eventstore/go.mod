module eventstore

go 1.13

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/liftbridge-io/go-liftbridge/v2 v2.0.1
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.35.0

	types v0.0.0
)

replace types => ../types
