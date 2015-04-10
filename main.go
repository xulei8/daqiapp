package main

import (
	"github.com/astaxie/beego"
	m "github.com/xulei8/daqiapp/models"
	_ "github.com/xulei8/daqiapp/routers"
	"os"
)

func main() {
	//m.Test()
	//os.Exit(0)
	beego.InitModXML()
	beego.ModTest()
	initArgs()
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "--newdb" || v == "new" || v == "newdb" {
			m.Syncdb()
			os.Exit(0)

		}
	}
}
