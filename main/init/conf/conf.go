/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:34:18
 * @Desc: 配置
 */
package conf

import (
	"selfx/init/command"
	"strings"

	"github.com/spf13/viper"
)

type DbDriver string

const (
	fieldAddr = "addr"
	fieldDB   = "db"
	fieldDSN  = "dsn"

	DbDriverSqlite     DbDriver = "sqlite"
	DbDriverMysql      DbDriver = "mysql"
	DbDriverPostgresql DbDriver = "postgresql"
)

var (
	Addr string   // 监听地址
	DB   DbDriver // 数据库类型
	DSN  string   // 链接数据库DSN
)

func init() {
	p := viper.New()

	// get by file
	p.SetConfigFile(command.ConfFilePath)
	p.SetDefault(fieldAddr, command.Addr)
	p.SetDefault(fieldDB, "sqlite")
	p.SetDefault(fieldDSN, "./selfx.db?_pragma=journal_mode(WAL)")

	_ = p.ReadInConfig()
	_ = p.WriteConfig()

	// get by env
	p.SetEnvPrefix("selfx")
	_ = p.BindEnv(fieldAddr)
	_ = p.BindEnv(fieldDB)
	_ = p.BindEnv(fieldDSN)

	if command.Addr != "" {
		p.Set(fieldAddr, command.Addr)
	}

	Addr = p.GetString(fieldAddr)
	DB = FormatDbDriver(p.GetString(fieldDB))
	DSN = p.GetString(fieldDSN)
}

func FormatDbDriver(val string) DbDriver {
	switch strings.ToLower(val) {
	case "sqlite", "sqlite3":
		return DbDriverSqlite
	case "mysql", "mariadb", "maria":
		return DbDriverMysql
	case "pgsql", "postgres", "postgresql":
		return DbDriverPostgresql
	default:
		return DbDriver(val)
	}
}
