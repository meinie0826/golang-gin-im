package main

import (
	"golangim/router"
	"golangim/utils"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router() // router.Router()
	r.Run(viper.GetString("port.server"))
}
