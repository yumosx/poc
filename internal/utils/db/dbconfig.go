package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config 数据库信息配置
// 1. 通过 function option 模式配置
// 2. 通过 toml 文件进行配置
type Config struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
}

type ConfigOption interface {
	Option(config *Config)
}

type ConfigOptionFunc func(config *Config)

func (fn ConfigOptionFunc) Option(config *Config) {
	fn(config)
}

func WithUserName(userName string) ConfigOptionFunc {
	return ConfigOptionFunc(func(config *Config) {
		config.UserName = userName
	})
}

func WithPassword(password string) ConfigOptionFunc {
	return ConfigOptionFunc(func(config *Config) {
		config.Password = password
	})
}

func WithHost(host string) ConfigOptionFunc {
	return ConfigOptionFunc(func(config *Config) {
		config.Host = host
	})
}

func WithPort(port string) ConfigOptionFunc {
	return ConfigOptionFunc(func(config *Config) {
		config.Port = port
	})
}

func WithDBName(db string) ConfigOptionFunc {
	return ConfigOptionFunc(func(config *Config) {
		config.DBName = db
	})
}

func NewConfig(options ...ConfigOption) *Config {
	config := &Config{}
	for _, opt := range options {
		opt.Option(config)
	}

	return config
}

// NewDB 根据参数建立对应的数据库连接
func NewDB(config *Config) (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func TearTables(db *gorm.DB, tables ...string) error {
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s", table)
		err := db.Exec(query).Error
		if err != nil {
			return err
		}
	}
	return nil
}
