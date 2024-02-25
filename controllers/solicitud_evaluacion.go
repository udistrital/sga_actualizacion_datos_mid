package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_actualizacion_dato_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
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
	defer errorhandler.HandlePanic(&c.Controller)

	id_solicitud := c.Ctx.Input.Param(":solicitud_id")

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
// @Param	id_estado_tipo_sol	path	int	true	"Id del estado tipo solicitud"
// @Success 200 {}
// @Failure 403 body is empty
// @router /estados/:tipo_estado_id [get]
func (c *SolicitudEvaluacionController) GetAllSolicitudActualizacionDatos() {
	//Consulta a tabla de solicitante la cual trae toda la info de la solicitud
	defer errorhandler.HandlePanic(&c.Controller)

	id_estado_tipo_sol := c.Ctx.Input.Param(":tipo_estado_id")

	respuesta := services.SolicitudActualizacionDatos(id_estado_tipo_sol)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
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
	defer errorhandler.HandlePanic(&c.Controller)

	id_persona := c.Ctx.Input.Param(":tercero_id")
	id_estado_tipo_solicitud := c.Ctx.Input.Param(":tipo_estado_id")

	respuesta := services.GetDatosSolicitud(id_persona, id_estado_tipo_solicitud)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta

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
	defer errorhandler.HandlePanic(&c.Controller)

	id_persona := c.Ctx.Input.Param(":persona_id")

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
// @Success 200 {}
// @Failure 404 not found resource
// @router /:id [get]
func (c *SolicitudEvaluacionController) PutSolicitudEvaluacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	//Id de la solicitud
	idSolicitud := c.Ctx.Input.Param(":id")

	respuesta := services.SolicitudEvaluacionPut(idSolicitud)

	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}
