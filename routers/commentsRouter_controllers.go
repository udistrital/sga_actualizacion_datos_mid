package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PostSolicitudActualizacionDatos",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PutSolicitudEvaluacion",
            Router: "/:id/evaluacion",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitudById",
            Router: "/:id_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAllSolicitudActualizacionDatos",
            Router: "/estados/:id_estado_tipo_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetSolicitudActualizacionDatos",
            Router: "/estudiantes/:id_tercero",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitud",
            Router: "/estudiantes/:id_tercero/estados/:id_estado_tipo_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PostSolicitudEvolucionEstado",
            Router: "/evoluciones",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
