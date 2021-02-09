package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"eventstore/bean"
)

// Server Server
var Server bean.ServerConfig

// MariaDB 数据库相关配置
var MariaDB bean.DBConfig

// Lift Lift
var Lift bean.LiftConfig

const cmdRoot = "core"

var p string

func init() {

	flag.StringVar(&p, "p", "/root/union/config", "set `prefix` path")
	flag.Parse()

	err := loadRemoteConfig(p)
	if err != nil {
		fmt.Println("load remote config error:", err.Error())
		os.Exit(0)
	}

	err = loadLocalConfig("./")
	if err != nil {
		fmt.Println("load local config error:", err.Error())
		os.Exit(0)
	}

}

func loadLocalConfig(path string) error {
	local := viper.New()
	local.SetEnvPrefix(cmdRoot)
	local.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	local.SetEnvKeyReplacer(replacer)
	local.SetConfigName(cmdRoot)
	local.AddConfigPath(path)

	err := local.ReadInConfig()
	if err != nil {
		return err
	}

	Server.GrpcPort = local.GetString("server.grpc.port")

	Lift.Addrs = local.GetStringSlice("lift.addrs")
	Lift.Partition = local.GetInt32("lift.partition")
	Lift.Subjects = local.GetStringMap("lift.subject")

	return nil
}

func loadRemoteConfig(path string) error {
	remote := viper.New()
	remote.SetEnvPrefix(cmdRoot)
	remote.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	remote.SetEnvKeyReplacer(replacer)
	remote.SetConfigName(cmdRoot)
	remote.AddConfigPath(path)

	err := remote.ReadInConfig()
	if err != nil {
		return err
	}

	MariaDB.Dialect = remote.GetString("database.dialect")
	MariaDB.Database = remote.GetString("database.database")
	MariaDB.User = remote.GetString("database.user")
	MariaDB.Password = remote.GetString("database.password")
	MariaDB.Host = remote.GetString("database.host")
	MariaDB.Port = remote.GetInt("database.port")
	MariaDB.Charset = remote.GetString("database.charset")
	MariaDB.MaxIdleConns = remote.GetInt("database.maxIdleConns")
	MariaDB.MaxOpenConns = remote.GetInt("database.maxOpenConns")
	MariaDB.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		MariaDB.User, MariaDB.Password, MariaDB.Host, MariaDB.Port, MariaDB.Database, MariaDB.Charset)

	return nil
}
