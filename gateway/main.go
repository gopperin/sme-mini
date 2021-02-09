package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/sync/errgroup"

	"gateway/config"
	"gateway/grpc"
	"gateway/persist"
	myrouter "gateway/router"
	"gateway/snowflake"
)

var (
	g errgroup.Group
)

func main() {

	err := persist.GMariadb.Init()
	if err != nil {
		fmt.Println("*** mariadb error : ", err.Error())
		return
	}
	fmt.Println("====== mariadb init ======")
	defer persist.Close()

	err = snowflake.Olaf.Init(1024, 1575043200000)
	if err != nil {
		fmt.Println("*** flake error : ", err.Error())
		return
	}
	fmt.Println("====== flake init ======")

	err = grpc.Init()
	if err != nil {
		fmt.Println("*** grpc client conn error : ", err.Error())
		return
	}
	fmt.Println("====== grpc client conn init ======")
	defer grpc.Close()

	router := gin.Default()

	router.Use(Cors())

	/* api base */
	myrouter.SetupBaseRouter(router)

	// User User
	myrouter.SetupUserRouter(router)

	server := &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	gracehttp.Serve(server)
}
