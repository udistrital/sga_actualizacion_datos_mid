package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/sga_actualizacion_dato_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
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
// @router /:id_solicitud [get]
func (c *SolicitudEvaluacionController) GetDatosSolicitudById() {
	defer errorhandler.HandlePanic(&c.Controller)

	id_solicitud := c.Ctx.Input.Param(":id_solicitud")

	respuesta := services.DatosSolicitud(id_solicitud)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
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
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	respuesta := services.SolicitudEvolucion(data)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// GetAllSolicitudActualizacionDatos ...
// @Title GetAllSolicitudActualizacionDatos
// @Description Consultar todas la solicitudes de actualización de datos
// @Param	id_estado_tipo_solicitud	path	int	true	"Id del estado tipo solicitud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estados/:id_estado_tipo_solicitud [get]
func (c *SolicitudEvaluacionController) GetAllSolicitudActualizacionDatos() {
	//Consulta a tabla de solicitante la cual trae toda la info de la solicitud
	defer errorhandler.HandlePanic(&c.Controller)

	id_estado_tipo_sol := c.Ctx.Input.Param(":id_estado_tipo_solicitud")

	respuesta := services.SolicitudActualizacionDatos(id_estado_tipo_sol)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// GetDatosSolicitud ...
// @Title GetDatosSolicitud
// @Description Consultar los datos ingresados por el estudiante en su solicitud
// @Param	id_tercero	path	int	true	"Id del estudiante"
// @Param	id_estado_tipo_solicitud	path	int	true	"Id del estado del tipo de solictud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estudiantes/:id_tercero/estados/:id_estado_tipo_solicitud [get]
func (c *SolicitudEvaluacionController) GetDatosSolicitud() {
	defer errorhandler.HandlePanic(&c.Controller)
	// id_persona --> id del tercero
	id_persona := c.Ctx.Input.Param(":id_tercero")
	id_estado_tipo_solicitud := c.Ctx.Input.Param(":id_estado_tipo_solicitud")

	respuesta := services.GetDatosSolicitud(id_persona, id_estado_tipo_solicitud)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta

	c.ServeJSON()
}

// GetSolicitudActualizacionDatos ...
// @Title GetSolicitudActualizacionDatos
// @Description Consultar la solicitudes de un estudiante de actualización de datos
// @Param	id_tercero	path	int	true	"Id del estudiante"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estudiantes/:id_tercero [get]
func (c *SolicitudEvaluacionController) GetSolicitudActualizacionDatos() {
	defer errorhandler.HandlePanic(&c.Controller)

	id_persona := c.Ctx.Input.Param(":id_tercero")

	respuesta := services.GetSolictudActualizacion(id_persona)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
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
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	respuesta := services.ActualizacionDatosPost(data)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// PutSolicitudEvaluacion ...
// @Title PutSolicitudEvaluacion
// @Description actualiza de forma publica el estado de una solicitud tipo evaluacion
// @Param	id	path	int	true	"Id de la solicitud"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:id/evaluacion [get]
func (c *SolicitudEvaluacionController) PutSolicitudEvaluacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	//Id de la solicitud
	idSolicitud := c.Ctx.Input.Param(":id")

	respuesta := services.SolicitudEvaluacionPut(idSolicitud)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// PutSolicitud...
// @Title PutSolicitud
// @Description Modifica una solicitud existente
// @Param   body        body    {}  true        "body Modificar solicitud content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /:id_solicitud [put]
func (c *SolicitudEvaluacionController) PutSolicitudReferencia() {
	defer errorhandler.HandlePanic(&c.Controller)

	idSolicitud := c.Ctx.Input.Param(":id_solicitud")

	var referencia map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &referencia); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato de la solicitud")
		c.ServeJSON()
		return
	}

	respuesta := services.PutSolicitudReferencia(idSolicitud, referencia)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}
