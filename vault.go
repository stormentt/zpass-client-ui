package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stormentt/zpass-client/keyvault"
	"github.com/therecipe/qt/core"
)

func init() {
	Vault_QmlRegisterType2("Vault", 1, 0, "Vault")
}

type Vault struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) bool `slot:"unlock"`
	_ func()            `signal:"unlocked"`
}

func (v *Vault) init() {
	v.ConnectUnlock(v.unlock)
}

func (v *Vault) unlock(pass string) bool {
	path := viper.GetString("keyvault-path")
	log.WithFields(log.Fields{
		"path": path,
	}).Debug("Unlocking Keyvault")

	err := keyvault.Init(path, []byte(pass))
	if err != nil {
		return false
	} else {
		v.Unlocked()
		return true
	}
}
