package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type config struct {
	Server struct {
		Name     string `mapstructure:"NAME"`
		GRPCPort string    `mapstructure:"GRPC_PORT"`
		HTTPPort string    `mapstructure:"HTTP_PORT"`
		TempoHost string    `mapstructure:"TEMPO_HOST"`
		TempoNameSpace string    `mapstructure:"TEMPO_NAMESPACE"`
	} `mapstructure:"SERVER"`
	Database struct {
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"POST"`
		User string    `mapstructure:"USER"`
		Password string    `mapstructure:"PASSWORD"`
		DBName string    `mapstructure:"DATABASE_NAME"`
		Schema string    `mapstructure:"SCHEMA"`
		ConnMaxLifetimeSecond int    `mapstructure:"CONN_MAX_LIFETIME_SECOND"`
		MaxOpenConn int    `mapstructure:"MAX_OPEN_CONN"`
		MaxIdleConn int    `mapstructure:"MAX_IDLE_CONN"`
	} `mapstructure:"DATABASE"`
}

var C config

func ReadConfig() {
	Config := &C
	Config.Server.Name = getEnv[string]("SERVER_NAME", "canaanadvisors-test")
	Config.Server.GRPCPort = getEnv[string]("SERVER_GRPC_PORT", "8001")
	Config.Server.HTTPPort = getEnv[string]("SERVER_HTTP_PORT", "9001")
	Config.Server.TempoHost = getEnv[string]("TEMPO_HOST", "temporal:7233")
	Config.Server.TempoNameSpace = getEnv[string]("TEMPO_NAMESPACE", "canaanadvisors-test")

	Config.Database.Host = getEnv[string]("DB_HOST", "localhost")
	Config.Database.Port = getEnv[string]("DB_PORT", "5432")
	Config.Database.User = getEnv[string]("DB_USER", "canaanadvisors")
	Config.Database.Password = getEnv[string]("DB_PASSWORD", "1qazxsw2")
	Config.Database.DBName = getEnv[string]("DB_NAME", "canaanadvisors")
	Config.Database.Schema = getEnv[string]("DB_SCHEMA", "public")
	Config.Database.ConnMaxLifetimeSecond = getEnv[int]("DB_CONN_MAX_LIFETIME_SECOND", 300)
	Config.Database.MaxOpenConn = getEnv[int]("DB_MAX_OPEN_CONN", 100)
	Config.Database.MaxIdleConn = getEnv[int]("DB_MAX_IDLE_CONN", 100)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

type EnvType interface {
	string | int
}

func getEnv[V EnvType](key string, fallback V) V {
	if value, ok := os.LookupEnv(key); ok {
		var convertVal V
		log.Println(value)
		_, err := fmt.Sscanf(value, "%v", &convertVal)
		if err == nil {
			return convertVal
		}
	}
	return fallback
}

func readConf() {
	Config := &C

	viper.SetConfigName(".env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(".env not found")
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}
}