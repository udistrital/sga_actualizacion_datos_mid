// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/sga_mid_actualizacion_datos/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.ErrorController(&controllers.ErrorHandlerController{})
	
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/actualizaciones-datos",
			beego.NSInclude(
				&controllers.SolicitudEvaluacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

