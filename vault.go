package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

func init() {
	Vault_QmlRegisterType2("Vault", 1, 0, "Vault")
}

type Vault struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) string `slot:"unlock"`
	_ func()              `signal:"unlocked"`
}

func (v *Vault) init() {
	v.ConnectUnlock(v.unlock)
}

func (v *Vault) unlock(pass string) string {
	fmt.Println(pass)
	v.Unlocked()
	return "hello world"
}
