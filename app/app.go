package app

import "github.com/leeroyakbar/messaging-app-db/app/configs"

func RunApp() {
	configs.ConnectDB()
}
