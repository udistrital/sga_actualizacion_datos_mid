// @APIVersion 1.0.0
// @Title Microservicio SGA MID - Solicitudes de Evaluación
// @Description Microservcio del SGA MID para solicitudes de actualización de datos - evaluación
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/sga_actualizacion_dato_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/solicitudes",
			beego.NSInclude(
				&controllers.SolicitudEvaluacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
