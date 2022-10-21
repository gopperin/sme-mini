package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/gopperin/sme-mini/gateway/bean"
)

// MariaDB 数据库相关配置
var MariaDB bean.DBConfig

// Server Server Config
var Server bean.ServerConfig

const cmdRoot = "core"

var p string

func init() {

	flag.StringVar(&p, "p", "./", "set `prefix` path")
	flag.Parse()

	err := loadLocalConfig(p)
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

	Server.Port = local.GetString("server.port")
	Server.ViewLimit = local.GetInt64("server.view.limit")
	Server.Release = local.GetString("server.release")
	Server.Version = local.GetString("server.version")
	Server.GrpcURI = local.GetString("server.grpc_uri")

	MariaDB.Dialect = local.GetString("database.dialect")
	MariaDB.Database = local.GetString("database.database")
	MariaDB.User = local.GetString("database.user")
	MariaDB.Password = local.GetString("database.password")
	MariaDB.Host = local.GetString("database.host")
	MariaDB.Port = local.GetInt("database.port")
	MariaDB.Charset = local.GetString("database.charset")
	MariaDB.MaxIdleConns = local.GetInt("database.maxIdleConns")
	MariaDB.MaxOpenConns = local.GetInt("database.maxOpenConns")
	MariaDB.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		MariaDB.User, MariaDB.Password, MariaDB.Host, MariaDB.Port, MariaDB.Database, MariaDB.Charset)

	return nil
}
