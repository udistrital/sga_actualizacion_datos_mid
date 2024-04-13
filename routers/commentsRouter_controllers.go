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
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitudById",
            Router: "/:solicitud_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAllSolicitudActualizacionDatos",
            Router: "/estados/:tipo_estado_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetDatosSolicitud",
            Router: "/estados/:tipo_estado_id/terceros/:tercero_id",
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

    beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_actualizacion_dato_mid/controllers:SolicitudEvaluacionController"],
        beego.ControllerComments{
            Method: "GetSolicitudActualizacionDatos",
            Router: "/terceros/:persona_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
