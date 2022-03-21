package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//define database struct
type DBConfig struct {
	DbServer              string        `json:"db_server"`
	DbPort                int           `json:"db_port"`
	Username              string        `json:"db_username"`
	Password              string        `json:"db_password"`
	Name                  string        `json:"db_name"`
	Dbprovider            string        `json:"db_provider"`
	Dbfailover            string        `json:"db_failover"`
	MaxOpenConnections    int           `json:"maxOpenConnections"`
	MaxIdleConnections    int           `json:"maxIdleConnections"`
	ConnectionMaxLifetime time.Duration `json:"connectionMaxLifetime"`
	ConnectionMaxIdleTime time.Duration `json:"connectionMaxIdleTime"`
}

var db *sql.DB

func Get() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}

func InitDatabase(ctx *gin.Context, Config DBConfig) error {

	//initialize error variable
	var err error

	//init database connection
	log.Info(ctx).Interface(DatabaseConfigKey, Config).Msg("initializing database")

	//get ParseTimer Value
	// ParseTime := "?parseTime=" + configs.Get().GetStringD(constants.ApplicationConfig, constants.ParseTimeKey, "")

	//open a db connection
	db, err = sql.Open(Config.Dbprovider, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.Username, Config.Password, Config.DbServer, Config.DbPort, Config.Name))
	if err != nil {
		return err
	}

	// try to ping
	err = db.Ping()
	if err != nil {
		return err
	}

	// now set the configurations
	db.SetMaxOpenConns(Config.MaxOpenConnections)
	db.SetMaxIdleConns(Config.MaxIdleConnections)
	db.SetConnMaxIdleTime(Config.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(Config.ConnectionMaxLifetime)

	return nil
}
func CloseDatabse(ctx *gin.Context) error {
	err := Close()
	if err != nil {
		return err
	}
	return nil
}
