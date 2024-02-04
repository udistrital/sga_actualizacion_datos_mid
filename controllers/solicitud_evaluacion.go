package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_mid_actualizacion_datos/models"
	"github.com/udistrital/sga_mid_actualizacion_datos/services"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// SolicitudEvaluacionController ...
type SolicitudEvaluacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudEvaluacionController) URLMapping() {
	c.Mapping("PutSolicitudEvaluacion", c.PutSolicitudEvaluacion)
	c.Mapping("PostSolicitudActualizacionDatos", c.PostSolicitudActualizacionDatos)
	c.Mapping("GetSolicitudActualizacionDatos", c.GetSolicitudActualizacionDatos)
	c.Mapping("GetDatosSolicitud", c.GetDatosSolicitud)
	c.Mapping("GetAllSolicitudActualizacionDatos", c.GetAllSolicitudActualizacionDatos)
	c.Mapping("PostSolicitudEvolucionEstado", c.PostSolicitudEvolucionEstado)
	c.Mapping("GetDatosSolicitudById", c.GetDatosSolicitudById)
}

// GetDatosSolicitudById ...
// @Title GetDatosSolicitudById
// @Description Consultar los datos ingresados por el estudiante en su solicitud consultando por id de la solicitud
// @Param	id_solicitud	path	int	true	"Id de la solicitud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /:solicitud_id [get]
func (c *SolicitudEvaluacionController) GetDatosSolicitudById() {
	id_solicitud := c.Ctx.Input.Param(":solicitud_id")
	var Solicitud map[string]interface{}
	var TipoDocumentoGet map[string]interface{}
	var TipoDocumentoActualGet map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+id_solicitud, &Solicitud)
	if errSolicitud == nil {
		if Solicitud != nil && fmt.Sprintf("%v", Solicitud) != "map[]" {
			Referencia := Solicitud["Referencia"].(string)
			resultado["FechaSolicitud"] = Solicitud["FechaRadicacion"]
			var ReferenciaJson map[string]interface{}
			if err := json.Unmarshal([]byte(Referencia), &ReferenciaJson); err == nil {
				formatdata.JsonPrint(ReferenciaJson)
				TipoSolicitud := Solicitud["EstadoTipoSolicitudId"].(map[string]interface{})["Id"]
				TipoSolicitudId, _ := strconv.ParseInt(fmt.Sprintf("%v", TipoSolicitud), 10, 64)
				if TipoSolicitudId == 15 || TipoSolicitudId == 17 || TipoSolicitudId == 20 || TipoSolicitudId == 33 {
					TipoDocumento := fmt.Sprintf("%v", ReferenciaJson["DatosAnteriores"].(map[string]interface{})["TipoDocumentoActual"].(map[string]interface{})["Id"])
					services.ConfigurarResultadoGetSolicitudId(&resultado, &ReferenciaJson, 1)

					c.Data["json"] = services.SolicitudTipoDocGetSolicitudId(TipoDocumento, &TipoDocumentoGet, &resultado, &alerta, &alertas, &errorGetAll)

					c.Data["json"] = services.SolicitudDocActualGetSolicitudId(ReferenciaJson, &TipoDocumentoActualGet, &resultado, &alerta, &alertas, &errorGetAll)
				} else if TipoSolicitudId == 16 || TipoSolicitudId == 18 || TipoSolicitudId == 19 || TipoSolicitudId == 32 {
					services.ConfigurarResultadoGetSolicitudId(&resultado, &ReferenciaJson, 2)
				}
			}
		} else {
			services.ManejoError(&alerta, &alertas, "No data found", &errorGetAll)
			c.Data["json"] = map[string]interface{}{"Response": alerta}
		}
	} else {
		services.ManejoError(&alerta, &alertas, "", &errorGetAll, errSolicitud)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// PostSolicitudEvolucionEstado ...
// @Title PostSolicitudEvolucionEstado
// @Description Agregar una evolucion del estado a la solicitud planteada
// @Param   body        body    {}  true        "body Agregar una evolucion del estado a la solicitud planteada content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /evoluciones [post]
func (c *SolicitudEvaluacionController) PostSolicitudEvolucionEstado() {
	var Solicitud map[string]interface{}
	var SolicitudAux map[string]interface{}
	var SolicitudAuxPost map[string]interface{}
	var SolicitudEvolucionEstado []map[string]interface{}
	var EstadoTipoSolicitudId int
	var SolicitudEvolucionEstadoPost map[string]interface{}
	var ObservacionPost map[string]interface{}
	var SolicitudAprob map[string]interface{}
	var Tercero map[string]interface{}
	var TerceroPut map[string]interface{}
	var DatosIdentificacion []map[string]interface{}
	var DatosIdentificacionPut map[string]interface{}
	var DatosIdentificacionPost map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &Solicitud); err == nil {
		c.Data["json"] = services.SolicitudEstadoPostSolicitud(&resultado, &SolicitudEvolucionEstadoPost, &SolicitudAux, EstadoTipoSolicitudId, &SolicitudAuxPost, Solicitud, &ObservacionPost, &errorGetAll, &alerta, &alertas, &SolicitudAprob, &Tercero, &TerceroPut, &DatosIdentificacion, &DatosIdentificacionPut, &DatosIdentificacionPost, &SolicitudEvolucionEstado)
	} else {
		services.ManejoError(&alerta, &alertas, "", &errorGetAll, err)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// GetAllSolicitudActualizacionDatos ...
// @Title GetAllSolicitudActualizacionDatos
// @Description Consultar todas la solicitudes de actualización de datos
// @Param	id_estado_tipo_sol	path	int	true	"Id del estado tipo solicitud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estados/:tipo_estado_id [get]
func (c *SolicitudEvaluacionController) GetAllSolicitudActualizacionDatos() {
	//Consulta a tabla de solicitante la cual trae toda la info de la solicitud
	id_estado_tipo_sol := c.Ctx.Input.Param(":tipo_estado_id")
	var Solicitudes []map[string]interface{}
	var TipoSolicitud map[string]interface{}
	var Estado map[string]interface{}
	var Observacion []map[string]interface{}
	var respuesta []map[string]interface{}
	//var respuestaAux []map[string]in
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	c.Data["json"] = services.ManejoSolicitudesGetAll(&Solicitudes, Observacion, &respuesta, &TipoSolicitud, &Estado, &errorGetAll, &alertas, &alerta, id_estado_tipo_sol, &resultado)

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// GetDatosSolicitud ...
// @Title GetDatosSolicitud
// @Description Consultar los datos ingresados por el estudiante en su solicitud
// @Param	id_persona	path	int	true	"Id del estudiante"
// @Param	id_estado_tipo_solicitud	path	int	true	"Id del estado del tipo de solictud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estados/:tipo_estado_id/terceros/:tercero_id [get]
func (c *SolicitudEvaluacionController) GetDatosSolicitud() {
	id_persona := c.Ctx.Input.Param(":tercero_id")
	id_estado_tipo_solicitud := c.Ctx.Input.Param(":tipo_estado_id")
	var Solicitudes []map[string]interface{}
	var TipoDocumentoGet map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	c.Data["json"] = services.SolicitudGetDatos(&resultado, &TipoDocumentoGet, &errorGetAll, &alertas, &alerta, id_persona, id_estado_tipo_solicitud, &Solicitudes)

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// GetSolicitudActualizacionDatos ...
// @Title GetSolicitudActualizacionDatos
// @Description Consultar la solicitudes de un estudiante de actualización de datos
// @Param	id_persona	path	int	true	"Id del estudiante"
// @Success 200 {}
// @Failure 403 body is empty
// @router /terceros/:persona_id [get]
func (c *SolicitudEvaluacionController) GetSolicitudActualizacionDatos() {
	id_persona := c.Ctx.Input.Param(":persona_id")
	var Solicitudes []map[string]interface{}
	var TipoSolicitud map[string]interface{}
	var Estado map[string]interface{}
	var respuesta []map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	c.Data["json"] = services.ManejoSolicitudesGetActualizacion(&Solicitudes, id_persona, &respuesta, &TipoSolicitud, &Estado, &errorGetAll, &alertas, &alerta, &resultado)

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// PostSolicitudActualizacionDatos ...
// @Title PostSolicitudActualizacionDatos
// @Description Agregar una solicitud de actualizacion de datos(ID o nombre)
// @Param   body        body    {}  true        "body Agregar solicitud actualizacion datos content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *SolicitudEvaluacionController) PostSolicitudActualizacionDatos() {
	var Solicitud map[string]interface{}
	var SolicitudPadre map[string]interface{}
	var SolicitudPost map[string]interface{}
	var SolicitantePost map[string]interface{}
	var Referencia string
	var IdEstadoTipoSolicitud int
	var SolicitudEvolucionEstadoPost map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &Solicitud); err == nil {
		c.Data["json"] = services.ManejoSolicitudesPostActualizacion(IdEstadoTipoSolicitud, &SolicitudEvolucionEstadoPost, &resultado, &SolicitantePost, &errorGetAll, &alertas, &alerta, &SolicitudPost, Solicitud, Referencia, &SolicitudPadre)
	} else {
		services.ManejoError(&alerta, &alertas, "", &errorGetAll, err)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	if !errorGetAll {
		services.ManejoExito(&alerta, &alertas, resultado)
		c.Data["json"] = map[string]interface{}{"Response": alerta}
	}

	c.ServeJSON()
}

// PutSolicitudEvaluacion ...
// @Title PutSolicitudEvaluacion
// @Description actualiza de forma publica el estado de una solicitud tipo evaluacion
// @Success 200 {}
// @Failure 404 not found resource
// @router /:id [get]
func (c *SolicitudEvaluacionController) PutSolicitudEvaluacion() {
	//Id de la solicitud
	idSolicitud := c.Ctx.Input.Param(":id")
	//resultado resultado final
	var resultadoPutSolicitud map[string]interface{}
	resultadoRechazo := make(map[string]interface{})

	var solicitudEvaluacion map[string]interface{}
	if solicitudEvaluacionList, errGet := models.GetOneSolicitudDocente(idSolicitud); errGet == nil {
		if errorSystem, dataJson := services.ManejoSolicitudes(solicitudEvaluacion, solicitudEvaluacionList, resultadoRechazo, idSolicitud, resultadoPutSolicitud); errorSystem == nil {
			c.Data["json"] = dataJson
		} else {
			c.Data["system"] = errorSystem
			c.Abort("400")
		}
	} else {
		logs.Error(errGet)
		c.Data["system"] = resultadoPutSolicitud
		c.Abort("400")
	}
	c.ServeJSON()
}
