package main

import (
	"fmt"
	"sing-box-telegram/service/builder"
	"sing-box-telegram/service/donate"
	"sing-box-telegram/service/execute"
	"sing-box-telegram/service/subscribe"
	"sing-box-telegram/service/telegram"
)

func main() {

	executeInstance := execute.NewExecute()
	executeInstance.ExecuteCommand("./reinstall.sh")

	fmt.Println("read setting file...")

	builderInstance := builder.NewBuilder().
		SetServerIP().
		SetSettingsFile().
		SetPublicKeyAndPrivateKey().
		SetConfigurations().
		SetBlock().
		Save()

	sub := subscribe.NewSubscribe(builderInstance, executeInstance)

	tel := telegram.NewTelegramClient(builderInstance, sub)

	donate.NewDonate(builderInstance, sub)

	tel.SendVNstat()

}
