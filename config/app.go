package config

import (
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
}

var C config

func ReadConfig() {
	Config := &C
	Config.Server.Name = getEnv("SERVER_NAME", "canaanadvisors-test")
	Config.Server.GRPCPort = getEnv("SERVER_GRPC_PORT", "8001")
	Config.Server.HTTPPort = getEnv("SERVER_HTTP_PORT", "9001")
	Config.Server.TempoHost = getEnv("TEMPO_HOST", "temporal:7233")
	Config.Server.TempoNameSpace = getEnv("TEMPO_NAMESPACE", "canaanadvisors-test")
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func readConf() {
	Config := &C

	viper.SetConfigName("./config/.env")
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