package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetSolicitudActualizacionDatos",
            Router: "/personas/:id_persona/solicitudes",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitud",
            Router: "/personas/:id_persona/solicitudes/estados/:id_estado_tipo_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PostSolicitudActualizacionDatos",
            Router: "/solicitudes",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitudById",
            Router: "/solicitudes/:id_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PutSolicitudEvaluacion",
            Router: "/solicitudes/:id_solicitud/estado",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAllSolicitudActualizacionDatos",
            Router: "/solicitudes/estados/:id_estado_tipo_solicitud",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_actualizacion_datos/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "PostSolicitudEvolucionEstado",
            Router: "/solicitudes/evoluciones",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
