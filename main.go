package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stormentt/zpass-lib/crypt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	initConfig()
	initLogging()
	initCrypt()
	gui.NewQGuiApplication(len(os.Args), os.Args)

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".github.com/stormentt/zpass-client" (without extension).
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName("zpass-client")
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogging() {
	log.SetLevel(log.DebugLevel)
}

func initCrypt() {
	viper.SetDefault("Hasher", "sha512")
	viper.SetDefault("Crypter", "xsalsa20")

	crypt.ConfigHasher = viper.GetString("Hasher")
	crypt.ConfigCrypter = viper.GetString("Crypter")
}
