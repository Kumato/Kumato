package controller

import (
	"github.com/kumato/kumato/internal/logger"
	"github.com/spf13/viper"
	"time"
)

var (
	configPath = "./"
)

func ReadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".kumato") // name of config file (without extension)
	viper.SetConfigType("yml")     // REQUIRED if the config file does not have the extension in the name

	if path != "" {
		configPath = path + "/"
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Info("config file not found:", err.Error())
		} else {
			// Config file was found but another error was produced
			logger.Fatal("fail to read kumato.yml", err.Error())
		}
		return
	}

	for k, v := range viper.GetStringMapString("node") {
		client, err := registerConn(k, v)
		if err != nil {
			continue
		}
		nodes.LoadOrStore(k, client)
	}

	time.Sleep(5 * time.Second)
	go AssignTask()
}

func saveConfig(k, v string) {
	m := viper.GetStringMapString("node")
	m[k] = v
	viper.Set("node", m)
}

func writeConfig() {
	if err := viper.WriteConfigAs(configPath + "/.kumato.yml"); err != nil {
		logger.Fatal("error appeared when writing config:", err.Error())
	}
}
