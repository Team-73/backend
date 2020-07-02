package config

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// GetDBConfig to read initial config
func GetDBConfig() (config entity.InitialConfig) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config.Username = cast.ToString(viper.Get("DB_USER"))
	config.Password = cast.ToString(viper.Get("DB_PASSWORD"))
	config.Host = cast.ToString(viper.Get("DB_HOST"))
	config.Port = cast.ToString(viper.Get("DB_PORT"))
	config.DBName = cast.ToString(viper.Get("DB_NAME"))
	config.MaxLifeInMinutes = cast.ToInt(viper.Get("MAX_LIFE_IN_MINUTES"))
	config.MaxIdleConns = cast.ToInt(viper.Get("MAX_IDLE_CONNS"))
	config.MaxOpenConns = cast.ToInt(viper.Get("MAX_OPEN_CONNS"))

	return
}
