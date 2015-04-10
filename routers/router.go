package routers

import (
	"github.com/astaxie/beego"
	"github.com/xulei8/daqiapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/autodata_save", &controllers.DataSave{})
	beego.Router("/autodata_get", &controllers.DataGet{})
	beego.Router("/autodata_delete", &controllers.DataDelete{})
	beego.Router("/xml", &controllers.Xml{})
	beego.Router("/xmldata_save", &controllers.XmlSave{})
	beego.Router("/xmldata_get", &controllers.XmlGet{})

}
