package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
	"strconv"
	"time"
)

// GetOneSolicitudDocente is ...
func GetOneSolicitudDocente(idSolicitud string) (result []interface{}, outputError interface{}) {
	fmt.Println("id solicitud: ", idSolicitud)
	var solicitudes []map[string]interface{}
	var v []interface{}
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"/solicitud/?query=Id:"+idSolicitud, &solicitudes)
	if errSolicitud == nil && fmt.Sprintf("%v", solicitudes[0]["System"]) != "map[]" {
		if solicitudes[0]["Status"] != 404 && solicitudes[0]["Id"] != nil {

			var solicitantes []map[string]interface{}
			errSolicitante := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"/solicitante/?query=SolicitudId:"+idSolicitud, &solicitantes)
			if errSolicitante == nil && fmt.Sprintf("%v", solicitantes[0]["System"]) != "map[]" {
				if solicitantes[0]["Status"] != 404 {

					var evolucionEstado []map[string]interface{}
					errEvolucion := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"/solicitud_evolucion_estado/?limit=0&query=SolicitudId:"+idSolicitud, &evolucionEstado)
					if errEvolucion == nil && fmt.Sprintf("%v", evolucionEstado[0]["System"]) != "map[]" {
						if evolucionEstado[0]["Status"] != 404 && evolucionEstado[0]["Id"] != nil {

							var observaciones []map[string]interface{}
							errObservacion := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"/observacion/?limit=0&query=SolicitudId:"+idSolicitud, &observaciones)
							if errObservacion == nil && fmt.Sprintf("%v", observaciones[0]["System"]) != "map[]" {
								if observaciones[0]["Status"] != 404 {

									v = append(v, map[string]interface{}{
										"Id":                    solicitudes[0]["Id"],
										"EstadoTipoSolicitudId": solicitudes[0]["EstadoTipoSolicitudId"],
										"Referencia":            solicitudes[0]["Referencia"],
										"Resultado":             solicitudes[0]["Resultado"],
										"FechaRadicacion":       solicitudes[0]["FechaRadicacion"],
										"Observaciones":         observaciones,
										"Solicitantes":          solicitantes,
										"EvolucionEstado":       evolucionEstado,
									})
									return v, nil
								}
							} else {
								logs.Error(observaciones)
								return nil, errObservacion
							}
						}
					} else {
						logs.Error(evolucionEstado)
						return nil, errEvolucion
					}
				}
			} else {
				logs.Error(solicitantes)
				return nil, errSolicitante
			}
		} else {
			logs.Error(solicitudes)
			return nil, errSolicitud
		}
	} else {
		logs.Error(solicitudes)
		return nil, errSolicitud
	}
	return v, nil
}

// PutSolicitudDocente is ...
func PutSolicitudDocente(SolicitudDocente map[string]interface{}, idStr string) (result map[string]interface{}, outputError interface{}) {
	date := time_bogota.TiempoBogotaFormato()
	//resultado experiencia
	var resultado map[string]interface{}
	SolicitudDocentePut := make(map[string]interface{})
	fechaRadicacion := time_bogota.TiempoCorreccionFormato(fmt.Sprintf("%v", SolicitudDocente["FechaRadicacion"]))
	yesterday, _ := strconv.Atoi(fmt.Sprintf("%v", SolicitudDocente["EstadoTipoSolicitudId"].(map[string]interface{})["Id"]))
	if yesterday == 1 {
		SolicitudDocentePut["Solicitud"] = map[string]interface{}{
			"Resultado":             SolicitudDocente["Resultado"],
			"Referencia":            SolicitudDocente["Referencia"],
			"FechaRadicacion":       date,
			"EstadoTipoSolicitudId": SolicitudDocente["EstadoTipoSolicitudId"],
			"FechaModificacion":     date,
		}
	} else {
		SolicitudDocentePut["Solicitud"] = map[string]interface{}{
			"Resultado":             SolicitudDocente["Resultado"],
			"Referencia":            SolicitudDocente["Referencia"],
			"FechaRadicacion":       fechaRadicacion,
			"EstadoTipoSolicitudId": SolicitudDocente["EstadoTipoSolicitudId"],
			"FechaModificacion":     date,
		}
	}
	var EstadoTipoSolicitudID interface{}
	for _, evolucionEstadoTemp := range SolicitudDocente["EvolucionEstado"].([]interface{}) {
		evolucionEstado := evolucionEstadoTemp.(map[string]interface{})
		EstadoTipoSolicitudID = evolucionEstado["EstadoTipoSolicitudId"]
	}

	var solicitudesEvolucionEstado []map[string]interface{}
	solicitudesEvolucionEstado = append(solicitudesEvolucionEstado, map[string]interface{}{
		"TerceroId":                     SolicitudDocente["TerceroId"],
		"SolicitudId":                   map[string]interface{}{"Id": 0},
		"EstadoTipoSolicitudId":         SolicitudDocente["EstadoTipoSolicitudId"],
		"EstadoTipoSolicitudIdAnterior": EstadoTipoSolicitudID,
		"Activo":                        true,
		"FechaLimite":                   CalcularFecha(SolicitudDocente["EstadoTipoSolicitudId"].(map[string]interface{})),
		"FechaCreacion":                 date,
		"FechaModificacion":             date,
	})

	var observaciones []map[string]interface{}
	for _, observacionTemp := range SolicitudDocente["Observaciones"].([]interface{}) {
		observacion := observacionTemp.(map[string]interface{})
		if observacion["Id"] == nil && observacion["Titulo"] != nil {
			observaciones = append(observaciones, map[string]interface{}{
				"TipoObservacionId": observacion["TipoObservacionId"],
				"SolicitudId":       map[string]interface{}{"Id": 0},
				"TerceroId":         observacion["TerceroId"],
				"Titulo":            observacion["Titulo"],
				"Valor":             observacion["Valor"],
				"FechaCreacion":     date,
				"FechaModificacion": date,
				"Activo":            true,
			})
		} else if observacion["Id"] != nil {
			observaciones = append(observaciones, map[string]interface{}{
				"Id":                observacion["Id"],
				"TipoObservacionId": observacion["TipoObservacionId"],
				"SolicitudId":       observacion["SolicitudId"],
				"TerceroId":         observacion["TerceroId"],
				"Titulo":            observacion["Titulo"],
				"Valor":             observacion["Valor"],
				"Activo":            true,
			})
		}
	}
	if len(observaciones) == 0 {
		observaciones = append(observaciones, map[string]interface{}{})
	}

	SolicitudDocentePut["Solicitantes"] = nil
	SolicitudDocentePut["EvolucionesEstado"] = solicitudesEvolucionEstado
	SolicitudDocentePut["Observaciones"] = observaciones

	var resultadoSolicitudDocente map[string]interface{}
	errSolicitudPut := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"/tr_solicitud/"+idStr, "PUT", &resultadoSolicitudDocente, SolicitudDocentePut)
	if errSolicitudPut == nil && fmt.Sprintf("%v", resultadoSolicitudDocente["System"]) != "map[]" {
		if resultadoSolicitudDocente["Status"] != 400 {
			resultado = SolicitudDocente
			return resultado, nil
		}
	} else {
		logs.Error(errSolicitudPut)
		return nil, errSolicitudPut
	}
	return resultado, nil
}

// CalcularFecha is ...
func CalcularFecha(EstadoTipoSolicitud map[string]interface{}) (result string) {
	numDias, _ := strconv.Atoi(fmt.Sprintf("%v", EstadoTipoSolicitud["NumeroDias"]))
	var tiempoBogota time.Time
	tiempoBogota = time.Now()

	tiempoBogota = tiempoBogota.AddDate(0, 0, numDias)

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		fmt.Println(err)
	}
	tiempoBogota = tiempoBogota.In(loc)

	var tiempoBogotaStr = tiempoBogota.Format(time.RFC3339Nano)
	return tiempoBogotaStr
}
