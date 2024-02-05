// @APIVersion 1.0.0
// @Title Microservicio SGA MID - Solicitudes de Evaluación
// @Description Microservcio del SGA MID para solicitudes de actualización de datos - evaluación
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_mid_actualizacion_datos/controllers"
	"github.com/udistrital/utils_oas/errorhandler"
)

func init() {

	beego.ErrorController(&errorhandler.ErrorHandlerController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/actualizaciones-datos",
			beego.NSInclude(
				&controllers.SolicitudEvaluacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
