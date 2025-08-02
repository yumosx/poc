package ioc

import (
	"github.com/spf13/viper"
	"github.com/yumosx/poc/server/internal/repo/dao"
	"github.com/yumosx/poc/server/internal/utils/db"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	type Config struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"DBName"`
		UserName string `yaml:"userName"`
		Password string `yaml:"password"`
	}
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "13306")
	viper.SetDefault("DB_NAME", "ai_platform")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "root")

	config := &Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		DBName:   viper.GetString("DB_NAME"),
		UserName: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
	}

	newConfig := db.NewConfig(
		db.WithHost(config.Host),
		db.WithPort(config.Port),
		db.WithDBName(config.DBName),
		db.WithUserName(config.UserName),
		db.WithPassword(config.Password),
	)

	newDB, err := db.NewDB(newConfig)
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(newDB)
	if err != nil {
		panic(err)
	}
	return newDB
}
