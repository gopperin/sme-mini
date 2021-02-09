module gateway

go 1.13

require (
	github.com/Eric-GreenComb/contrib v0.0.0-20201231170135-6ed7a5914415
	github.com/btnguyen2k/consu/olaf v0.1.3
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/hashicorp/golang-lru v0.5.1
	github.com/jinzhu/gorm v1.9.16
	github.com/liftbridge-io/go-liftbridge/v2 v2.0.1
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.7.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.35.0

	types v0.0.0
)

replace types => ../types
