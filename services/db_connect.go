package services

import (
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	yaml "gopkg.in/yaml.v2"
)

// Configuration ->
var Configuration *Config

// Config ->
type Config struct {
	Port string `yaml:"port"`
	IP   string `yaml:"ip"`
	DB   DBConf `yaml:"db"`
}

// DBConf ->
type DBConf struct {
	DBName    string `yaml:"name"`
	DBUser    string `yaml:"user"`
	DBPass    string `yaml:"pass"`
	DBAddress string `yaml:"address"`
}

// Configs ->

// NewConfigurer ->
func NewConfigurer(path string) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening configuration file", err.Error())
	}

	err = yaml.Unmarshal(data, &Configuration)
	if err != nil {
		fmt.Println("Error unmarshaling configuration", err.Error())
	}

}

// Access ->
var Access *AccessCtrl

var (

	// SQL
	sqlDriver   = "mysql"
	sqlProtocol = "tcp"

	// MySQL services port
	sqlPort = "3306"
)

// AccessCtrl ->
type AccessCtrl struct {
	SQLDB *sqlx.DB
}

// NewAccess -> DB access constructor
func NewAccess(conf DBConf) {
	accessControl := new(AccessCtrl)

	if conf.DBAddress != "" {
		var err error
		fmt.Println(conf.DBUser + ":" + conf.DBPass + "@" + sqlProtocol + "(" + "database" + ":" + sqlPort + ")" + "/" + conf.DBName)
		accessControl.SQLDB, err = sqlx.Connect(sqlDriver, conf.DBUser+":"+conf.DBPass+"@"+sqlProtocol+"("+"database"+":"+sqlPort+")"+"/"+conf.DBName)

		if err != nil {
			panic("Error connecting to DB (" + err.Error() + ")")
		} else {
			fmt.Println("Connected to DB")
		}
	} else {
		panic("Could not connect to DB, no DB address")
	}

	Access = accessControl
}

// CreateDB ...
func (a AccessCtrl) CreateDB(dbname string) error {
	_, err := a.SQLDB.Exec("CREATE DATABASE " + dbname)

	if err == nil {
		_, err = a.SQLDB.Exec("USE " + dbname)
	}

	return err
}

// GetDB grants you pointer for DB query execution
func (a AccessCtrl) GetDB() *sqlx.DB {
	return a.SQLDB
}
