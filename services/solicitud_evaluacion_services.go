package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_mid_actualizacion_datos/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
)

// FUNCIONES QUE SE USAN EN GET DATOS SOLICITUD BY ID

func ConfigurarResultadoGetSolicitudId(resultado *map[string]interface{}, ReferenciaJson *map[string]interface{}, tipo int) {
	if tipo == 1 {
		(*resultado)["NumeroActual"] = (*ReferenciaJson)["DatosAnteriores"].(map[string]interface{})["NumeroActual"]
		(*resultado)["FechaExpedicionActual"] = (*ReferenciaJson)["DatosAnteriores"].(map[string]interface{})["FechaExpedicionActual"]
		(*resultado)["NumeroNuevo"] = (*ReferenciaJson)["DatosNuevos"].(map[string]interface{})["NumeroNuevo"]
		(*resultado)["FechaExpedicionNuevo"] = (*ReferenciaJson)["DatosNuevos"].(map[string]interface{})["FechaExpedicionNuevo"]
		(*resultado)["Documento"] = (*ReferenciaJson)["DocumentoId"]
	} else if tipo == 2 {
		(*resultado)["NombreActual"] = (*ReferenciaJson)["DatosAnteriores"].(map[string]interface{})["NombreActual"]
		(*resultado)["ApellidoActual"] = (*ReferenciaJson)["DatosAnteriores"].(map[string]interface{})["ApellidoActual"]
		(*resultado)["NombreNuevo"] = (*ReferenciaJson)["DatosNuevos"].(map[string]interface{})["NombreNuevo"]
		(*resultado)["ApellidoNuevo"] = (*ReferenciaJson)["DatosNuevos"].(map[string]interface{})["ApellidoNuevo"]
		(*resultado)["Documento"] = (*ReferenciaJson)["DocumentoId"]
	}
}

func SolicitudDocActualGetSolicitudId(ReferenciaJson map[string]interface{}, TipoDocumentoActualGet *map[string]interface{}, resultado *map[string]interface{}, alerta *models.Alert, alertas *[]interface{}, errorGetAll *bool) interface{} {
	TipoDocumentoAux := fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["TipoDocumentoNuevo"].(map[string]interface{})["Id"])
	errTipoDocumentoActual := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"tipo_documento/"+TipoDocumentoAux, TipoDocumentoActualGet)
	if errTipoDocumentoActual == nil {
		if *TipoDocumentoActualGet != nil && fmt.Sprintf("%v", *TipoDocumentoActualGet) != "map[]" {
			(*resultado)["TipoDocumentoNuevo"] = map[string]interface{}{
				"Id":     TipoDocumentoAux,
				"Nombre": (*TipoDocumentoActualGet)["Nombre"],
			}
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errTipoDocumentoActual)
		return map[string]interface{}{"Response": *alerta}
	}
}

func SolicitudTipoDocGetSolicitudId(TipoDocumento string, TipoDocumentoGet *map[string]interface{}, resultado *map[string]interface{}, alerta *models.Alert, alertas *[]interface{}, errorGetAll *bool) interface{} {
	errTipoDocumento := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"tipo_documento/"+TipoDocumento, TipoDocumentoGet)
	if errTipoDocumento == nil {
		if *TipoDocumentoGet != nil && fmt.Sprintf("%v", *TipoDocumentoGet) != "map[]" {
			(*resultado)["TipoDocumentoActual"] = map[string]interface{}{
				"Id":     TipoDocumento,
				"Nombre": (*TipoDocumentoGet)["Nombre"],
			}
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errTipoDocumento)
		return map[string]interface{}{"Response": *alerta}
	}
}

// FUNCIONES QUE SE USAN EN POST SOLICITUD EVOLUCION ESTADO

func solicitudPutTerceroPostSolicitud(TerceroId interface{}, TerceroPut *map[string]interface{}, Tercero map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	errTerceroPut := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"tercero/"+fmt.Sprintf("%v", TerceroId), "PUT", TerceroPut, Tercero)
	if errTerceroPut == nil {
		if *TerceroPut != nil && fmt.Sprintf("%v", *TerceroPut) != "map[]" {
			formatdata.JsonPrint(*TerceroPut)
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errTerceroPut)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudGetTerceroPostSolicitud(TerceroId interface{}, Tercero *map[string]interface{}, ReferenciaJson map[string]interface{}, TerceroPut *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	errTercero := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"tercero/"+fmt.Sprintf("%v", TerceroId), Tercero)
	if errTercero == nil {
		if *Tercero != nil && fmt.Sprintf("%v", *Tercero) != "map[]" {
			(*Tercero)["NombreCompleto"] = (fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["NombreNuevo"]) + " " + fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["ApellidoNuevo"]))
			Nombres := strings.SplitAfter(fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["NombreNuevo"]), " ")
			Apellidos := strings.SplitAfter(fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["ApellidoNuevo"]), " ")
			//Se actualiza el primer y segundo nombre (si lo tiene)
			if len(Nombres) > 1 {
				(*Tercero)["PrimerNombre"] = Nombres[0]
				(*Tercero)["SegundoNombre"] = Nombres[1]
			} else {
				(*Tercero)["PrimerNombre"] = Nombres[0]
				(*Tercero)["SegundoNombre"] = ""
			}
			//Se actualiza el primer y segundo apellido (si lo tiene)
			if len(Apellidos) > 1 {
				(*Tercero)["PrimerApellido"] = Apellidos[0]
				(*Tercero)["SegundoApellido"] = Apellidos[1]
			} else {
				(*Tercero)["PrimerApellido"] = Apellidos[0]
				(*Tercero)["SegundoApellido"] = ""
			}
			return solicitudPutTerceroPostSolicitud(TerceroId, TerceroPut, *Tercero, errorGetAll, alerta, alertas)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errTercero)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudDatosIdentificacionPostSolicitud(ReferenciaJson map[string]interface{}, TerceroId interface{}, DatosIdentificacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	DatosIdentificacionNuevo := map[string]interface{}{
		"TipoDocumentoId": map[string]interface{}{
			"Id": ReferenciaJson["DatosNuevos"].(map[string]interface{})["TipoDocumentoNuevo"].(map[string]interface{})["Id"],
		},
		"TerceroId": map[string]interface{}{
			"Id": TerceroId,
		},
		"Numero":          ReferenciaJson["DatosNuevos"].(map[string]interface{})["NumeroNuevo"],
		"FechaExpedicion": time_bogota.TiempoCorreccionFormato(ReferenciaJson["DatosNuevos"].(map[string]interface{})["FechaExpedicionNuevo"].(string)),
		"Activo":          true,
	}
	errDatosIDNuevo := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion", "POST", DatosIdentificacionPost, DatosIdentificacionNuevo)
	if errDatosIDNuevo == nil {
		if *DatosIdentificacionPost != nil && fmt.Sprintf("%v", *DatosIdentificacionPost) != "map[]" {
			formatdata.JsonPrint(*DatosIdentificacionPost)
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errDatosIDNuevo)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudPutDatosIdentificacionPostSolicitud(DatosIdentificacion []map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, ReferenciaJson map[string]interface{}, TerceroId interface{}, DatosIdentificacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	DatosIdentificacion[0]["Activo"] = false
	errDatosID := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion/"+fmt.Sprintf("%v", DatosIdentificacion[0]["Id"]), "PUT", DatosIdentificacionPut, DatosIdentificacion[0])
	if errDatosID == nil {
		if *DatosIdentificacionPut != nil && fmt.Sprintf("%v", *DatosIdentificacionPut) != "map[]" {
			//POST de los nuevos datos del terceros
			return solicitudDatosIdentificacionPostSolicitud(ReferenciaJson, TerceroId, DatosIdentificacionPost, errorGetAll, alerta, alertas)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errDatosID)
		return map[string]interface{}{"Response": *alerta}
	}
}

func manejoDatosIdentificacionPostSolicitud(TerceroId interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, ReferenciaJson map[string]interface{}, DatosIdentificacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	errTercero := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion?query=Activo:true,TerceroId__Id:"+fmt.Sprintf("%v", TerceroId)+"&sortby=Id&order=desc&limit=0", DatosIdentificacion)
	if errTercero == nil {
		if *DatosIdentificacion != nil && fmt.Sprintf("%v", (*DatosIdentificacion)[0]) != "map[]" {
			//Se cambia el estado de true a false en los datos_identificación antiguos
			return solicitudPutDatosIdentificacionPostSolicitud(*DatosIdentificacion, DatosIdentificacionPut, ReferenciaJson, TerceroId, DatosIdentificacionPost, errorGetAll, alerta, alertas)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errTercero)
		return map[string]interface{}{"Response": *alerta}
	}
}

func manejoTipoSolicitudPostSolicitud(SolicitudId string, SolicitudAprob *map[string]interface{}, EstadoTipoSolicitudId int, Tercero *map[string]interface{}, TerceroPut *map[string]interface{}, TerceroId interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, DatosIdentificacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+SolicitudId, SolicitudAprob)
	if errSolicitud == nil {
		if *SolicitudAprob != nil && fmt.Sprintf("%v", *SolicitudAprob) != "map[]" {
			Referencia := (*SolicitudAprob)["Referencia"].(string)
			var ReferenciaJson map[string]interface{}
			if err := json.Unmarshal([]byte(Referencia), &ReferenciaJson); err == nil {
				if EstadoTipoSolicitudId == 17 {
					//POST a terceros, a la tabla datos_identificacion por cambio de identificación
					return manejoDatosIdentificacionPostSolicitud(TerceroId, DatosIdentificacion, DatosIdentificacionPut, ReferenciaJson, DatosIdentificacionPost, errorGetAll, alerta, alertas)
				} else if EstadoTipoSolicitudId == 18 {
					//PUT a terceros, a la tabla tercero por cambio de nombre(s)
					return solicitudGetTerceroPostSolicitud(TerceroId, Tercero, ReferenciaJson, TerceroPut, errorGetAll, alerta, alertas)
				}
			}
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
	return nil
}

func solicitudObservacionPostSolicitud(Solicitud map[string]interface{}, TerceroId interface{}, Observacion interface{}, ObservacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}) interface{} {
	ObservacionAux := map[string]interface{}{
		"TipoObservacionId": map[string]interface{}{
			"Id": 1,
		},
		"SolicitudId": map[string]interface{}{
			"Id": Solicitud["SolicitudId"],
		},
		"TerceroId": TerceroId,
		"Valor":     Observacion,
		"Activo":    true,
	}

	errObservacion := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"observacion", "POST", ObservacionPost, ObservacionAux)
	if errObservacion == nil {
		if *ObservacionPost != nil && fmt.Sprintf("%v", *ObservacionPost) != "map[]" {
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errObservacion)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudAuxPostSolicitud(resultado *map[string]interface{}, SolicitudEvolucionEstadoPost map[string]interface{}, SolicitudAux map[string]interface{}, EstadoTipoSolicitudId int, SolicitudId string, SolicitudAuxPost *map[string]interface{}, Solicitud map[string]interface{}, TerceroId interface{}, Observacion interface{}, ObservacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}, SolicitudAprob *map[string]interface{}, Tercero *map[string]interface{}, TerceroPut *map[string]interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, DatosIdentificacionPost *map[string]interface{}) interface{} {
	var data interface{}
	SolicitudAux["EstadoTipoSolicitudId"].(map[string]interface{})["Id"] = EstadoTipoSolicitudId
	errSolicitudAux := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+SolicitudId, "PUT", SolicitudAuxPost, SolicitudAux)
	if errSolicitudAux == nil {
		if *SolicitudAuxPost != nil && fmt.Sprintf("%v", *SolicitudAuxPost) != "map[]" {
			//POST a observación (si hay alguna)
			if Observacion != "" {
				data = solicitudObservacionPostSolicitud(Solicitud, TerceroId, Observacion, ObservacionPost, errorGetAll, alerta, alertas)
			}

			// En caso de que la solicitud sea aprobada se traen los datos a cambiar y se hace POST a la respectiva tabla
			if EstadoTipoSolicitudId == 17 || EstadoTipoSolicitudId == 18 {
				data = manejoTipoSolicitudPostSolicitud(SolicitudId, SolicitudAprob, EstadoTipoSolicitudId, Tercero, TerceroPut, TerceroId, DatosIdentificacion, DatosIdentificacionPut, DatosIdentificacionPost, errorGetAll, alerta, alertas)
			}

			*resultado = SolicitudEvolucionEstadoPost
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "No data found", errorGetAll, errSolicitudAux)
		return map[string]interface{}{"Response": *alerta}
	}
	return data
}

func solicitudTablaPostSolicitud(resultado *map[string]interface{}, SolicitudEvolucionEstadoPost map[string]interface{}, SolicitudAux *map[string]interface{}, EstadoTipoSolicitudId int, SolicitudId string, SolicitudAuxPost *map[string]interface{}, Solicitud map[string]interface{}, TerceroId interface{}, Observacion interface{}, ObservacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}, SolicitudAprob *map[string]interface{}, Tercero *map[string]interface{}, TerceroPut *map[string]interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, DatosIdentificacionPost *map[string]interface{}) interface{} {
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+SolicitudId, SolicitudAux)
	if errSolicitud == nil {
		if *SolicitudAux != nil && fmt.Sprintf("%v", *SolicitudAux) != "map[]" {
			//Se reemplaza el estado de la solicitud anterior por la actual
			return solicitudAuxPostSolicitud(resultado, SolicitudEvolucionEstadoPost, *SolicitudAux, EstadoTipoSolicitudId, SolicitudId, SolicitudAuxPost, Solicitud, TerceroId, Observacion, ObservacionPost, errorGetAll, alerta, alertas, SolicitudAprob, Tercero, TerceroPut, DatosIdentificacion, DatosIdentificacionPut, DatosIdentificacionPost)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudEstadoNuevoPostSolicitud(resultado *map[string]interface{}, SolicitudEvolucionEstadoPost *map[string]interface{}, SolicitudAux *map[string]interface{}, EstadoTipoSolicitudId int, SolicitudId string, SolicitudAuxPost *map[string]interface{}, Solicitud map[string]interface{}, TerceroId interface{}, Observacion interface{}, ObservacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}, SolicitudAprob *map[string]interface{}, Tercero *map[string]interface{}, TerceroPut *map[string]interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, DatosIdentificacionPost *map[string]interface{}, EstadoTipoSolicitudIdAnterior interface{}, FechaLimite interface{}) interface{} {
	SolicitudEvolucionEstadoNuevo := map[string]interface{}{
		"TerceroId": TerceroId,
		"SolicitudId": map[string]interface{}{
			"Id": Solicitud["SolicitudId"],
		},
		"EstadoTipoSolicitudIdAnterior": map[string]interface{}{
			"Id": EstadoTipoSolicitudIdAnterior,
		},
		"EstadoTipoSolicitudId": map[string]interface{}{
			"Id": EstadoTipoSolicitudId,
		},
		"FechaLimite": FechaLimite,
		"Activo":      true,
	}

	//Se registra el nuevo estado de la solicitud en el historico
	errSolicitudEvolucionEstado := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud_evolucion_estado", "POST", SolicitudEvolucionEstadoPost, SolicitudEvolucionEstadoNuevo)
	if errSolicitudEvolucionEstado == nil {
		if *SolicitudEvolucionEstadoPost != nil && fmt.Sprintf("%v", *SolicitudEvolucionEstadoPost) != "map[]" {
			// GET a la tabla solicitud
			return solicitudTablaPostSolicitud(resultado, *SolicitudEvolucionEstadoPost, SolicitudAux, EstadoTipoSolicitudId, SolicitudId, SolicitudAuxPost, Solicitud, TerceroId, Observacion, ObservacionPost, errorGetAll, alerta, alertas, SolicitudAprob, Tercero, TerceroPut, DatosIdentificacion, DatosIdentificacionPut, DatosIdentificacionPost)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "No data found", errorGetAll, errSolicitudEvolucionEstado)
		return map[string]interface{}{"Response": *alerta}
	}
}

func SolicitudEstadoPostSolicitud(resultado *map[string]interface{}, SolicitudEvolucionEstadoPost *map[string]interface{}, SolicitudAux *map[string]interface{}, EstadoTipoSolicitudId int, SolicitudAuxPost *map[string]interface{}, Solicitud map[string]interface{}, ObservacionPost *map[string]interface{}, errorGetAll *bool, alerta *models.Alert, alertas *[]interface{}, SolicitudAprob *map[string]interface{}, Tercero *map[string]interface{}, TerceroPut *map[string]interface{}, DatosIdentificacion *[]map[string]interface{}, DatosIdentificacionPut *map[string]interface{}, DatosIdentificacionPost *map[string]interface{}, SolicitudEvolucionEstado *[]map[string]interface{}) interface{} {
	SolicitudId := fmt.Sprintf("%v", Solicitud["SolicitudId"])
	Estado := Solicitud["Estado"]
	Observacion := Solicitud["Observacion"]
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud_evolucion_estado?query=SolicitudId.Id:"+SolicitudId+"&sortby:Id&order:desc&limit=0", SolicitudEvolucionEstado)
	if errSolicitud == nil {
		if *SolicitudEvolucionEstado != nil && fmt.Sprintf("%v", (*SolicitudEvolucionEstado)[0]) != "map[]" {
			TerceroId := (*SolicitudEvolucionEstado)[0]["TerceroId"]
			EstadoTipoSolicitudIdAnterior := (*SolicitudEvolucionEstado)[0]["EstadoTipoSolicitudId"].(map[string]interface{})["Id"]
			FechaLimite := (*SolicitudEvolucionEstado)[0]["FechaLimite"]
			TipoSolicitudIdAux := (*SolicitudEvolucionEstado)[0]["SolicitudId"].(map[string]interface{})["EstadoTipoSolicitudId"].(map[string]interface{})["TipoSolicitud"].(map[string]interface{})["Id"]
			//Verifica si la solicitud es de actualización de identificación o de nombre
			TipoSolicitudId, _ := strconv.ParseInt(fmt.Sprintf("%v", TipoSolicitudIdAux), 10, 64)
			if TipoSolicitudId == 3 {
				//El tipo de solicitud es de cambio de identificación
				switch Estado {
				case 9.0: // Aprobado
					EstadoTipoSolicitudId = 17
				case 11.0: // Rechazado
					EstadoTipoSolicitudId = 20
				case 14.0: // Rectificar -> Modificada
					EstadoTipoSolicitudId = 33
				}
			} else if TipoSolicitudId == 4 {
				//El tipo de solicitud es de cambio de nombre
				switch Estado {
				case 9.0: // Aprobado
					EstadoTipoSolicitudId = 18
				case 11.0: // Rechazado
					EstadoTipoSolicitudId = 19
				case 14.0: // Rectificar -> Modificada
					EstadoTipoSolicitudId = 32
				}
			}

			//JSON de la nueva evolución del estado de la solicitud
			return solicitudEstadoNuevoPostSolicitud(resultado, SolicitudEvolucionEstadoPost, SolicitudAux, EstadoTipoSolicitudId, SolicitudId, SolicitudAuxPost, Solicitud, TerceroId, Observacion, ObservacionPost, errorGetAll, alerta, alertas, SolicitudAprob, Tercero, TerceroPut, DatosIdentificacion, DatosIdentificacionPut, DatosIdentificacionPost, EstadoTipoSolicitudIdAnterior, FechaLimite)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "No data found", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

// FUNCIONES QUE SE USAN EN GET ALL SOLICITUD ACTUALIZACION DE DATOS

func solicitudObservacionGetAll(Solicitudes []map[string]interface{}, i int, Observacion []map[string]interface{}, respuesta *[]map[string]interface{}, TipoSolicitud map[string]interface{}, Estado map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, errEstado error) interface{} {
	IdSolicitud := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"])
	errObservacion := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"observacion?query=SolicitudId:"+IdSolicitud, &Observacion)
	if errObservacion == nil {
		if Observacion != nil && fmt.Sprintf("%v", Observacion[0]) != "map[]" {
			(*respuesta)[i] = map[string]interface{}{
				"Numero":      Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"],
				"Fecha":       Solicitudes[i]["SolicitudId"].(map[string]interface{})["FechaRadicacion"],
				"Tipo":        TipoSolicitud["Data"].(map[string]interface{})["Nombre"],
				"Estado":      Estado["Data"].(map[string]interface{})["Nombre"],
				"Observacion": Observacion[0]["Valor"],
				"TerceroId":   Solicitudes[i]["TerceroId"],
			}
		} else {
			(*respuesta)[i] = map[string]interface{}{
				"Numero":      Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"],
				"Fecha":       Solicitudes[i]["SolicitudId"].(map[string]interface{})["FechaRadicacion"],
				"Tipo":        TipoSolicitud["Data"].(map[string]interface{})["Nombre"],
				"Estado":      Estado["Data"].(map[string]interface{})["Nombre"],
				"Observacion": "",
				"TerceroId":   Solicitudes[i]["TerceroId"],
			}
		}
		return nil
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errEstado)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudEstadoGetAll(Solicitudes []map[string]interface{}, i int, Observacion []map[string]interface{}, respuesta *[]map[string]interface{}, TipoSolicitud map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert) interface{} {
	IdEstado := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["EstadoTipoSolicitudId"].(map[string]interface{})["EstadoId"].(map[string]interface{})["Id"])
	//Nombre estado de la solicitud
	errEstado := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"estado/"+IdEstado, Estado)
	if errEstado == nil {
		if *Estado != nil && fmt.Sprintf("%v", *Estado) != "map[]" {
			// Observacion (Si la hay) sobre la solicitud
			return solicitudObservacionGetAll(Solicitudes, i, Observacion, respuesta, TipoSolicitud, *Estado, errorGetAll, alertas, alerta, errEstado)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errEstado)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudTipoGetAll(Solicitudes []map[string]interface{}, i int, Observacion []map[string]interface{}, respuesta *[]map[string]interface{}, TipoSolicitud *map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert) interface{} {
	IdTipoSolicitud := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["EstadoTipoSolicitudId"].(map[string]interface{})["TipoSolicitud"].(map[string]interface{})["Id"])
	//Nombre tipo solicitud
	errTipoSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"tipo_solicitud/"+IdTipoSolicitud, TipoSolicitud)
	if errTipoSolicitud == nil {
		if *TipoSolicitud != nil && fmt.Sprintf("%v", *TipoSolicitud) != "map[]" {
			return solicitudEstadoGetAll(Solicitudes, i, Observacion, respuesta, *TipoSolicitud, Estado, errorGetAll, alertas, alerta)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "No data found", errorGetAll, errTipoSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

func ManejoSolicitudesGetAll(Solicitudes *[]map[string]interface{}, Observacion []map[string]interface{}, respuesta *[]map[string]interface{}, TipoSolicitud *map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, id_estado_tipo_sol interface{}, resultado *map[string]interface{}) interface{} {
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitante?query=SolicitudId.EstadoTipoSolicitudId.Id:"+fmt.Sprintf("%v", id_estado_tipo_sol)+"&sortby:Id&order:asc&limit=0", Solicitudes)
	if errSolicitud == nil {
		if *Solicitudes != nil && fmt.Sprintf("%v", (*Solicitudes)[0]) != "map[]" {
			var data interface{}
			*respuesta = make([]map[string]interface{}, len(*Solicitudes))
			for i := 0; i < len(*Solicitudes); i++ {
				data = solicitudTipoGetAll(*Solicitudes, i, Observacion, respuesta, TipoSolicitud, Estado, errorGetAll, alertas, alerta)
				if data != nil {
					return data
				}
			}

			(*resultado)["Data"] = *respuesta
			return data
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

// FUNCIONES QUE SE USAN EN GET DATOS SOLICITUD

func asignacionresultadoSolicitud15GetDatos(resultado *map[string]interface{}, ReferenciaJson map[string]interface{}, TipoDocumentoGet *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, errSolicitud error) interface{} {
	(*resultado)["Documento"] = ReferenciaJson["DocumentoId"]
	(*resultado)["FechaExpedicionNuevo"] = ReferenciaJson["DatosNuevos"].(map[string]interface{})["FechaExpedicionNuevo"]
	(*resultado)["NumeroNuevo"] = ReferenciaJson["DatosNuevos"].(map[string]interface{})["NumeroNuevo"]
	TipoDocumento := fmt.Sprintf("%v", ReferenciaJson["DatosNuevos"].(map[string]interface{})["TipoDocumentoNuevo"].(map[string]interface{})["Id"])
	errTipoDocumento := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"tipo_documento/"+TipoDocumento, TipoDocumentoGet)
	if errTipoDocumento == nil {
		if *TipoDocumentoGet != nil && fmt.Sprintf("%v", *TipoDocumentoGet) != "map[]" {
			(*resultado)["TipoDocumentoNuevo"] = map[string]interface{}{
				"Id":     TipoDocumento,
				"Nombre": (*TipoDocumentoGet)["Nombre"],
			}
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

func asignacionresultadoSolicitud16GetDatos(resultado *map[string]interface{}, ReferenciaJson map[string]interface{}) {
	(*resultado)["ApellidoNuevo"] = ReferenciaJson["DatosNuevos"].(map[string]interface{})["ApellidoNuevo"]
	(*resultado)["NombreNuevo"] = ReferenciaJson["DatosNuevos"].(map[string]interface{})["NombreNuevo"]
	(*resultado)["Documento"] = ReferenciaJson["DocumentoId"]
}

func SolicitudGetDatos(resultado *map[string]interface{}, TipoDocumentoGet *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, id_persona string, id_estado_tipo_solicitud string, Solicitudes *[]map[string]interface{}) interface{} {
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitante?query=TerceroId:"+id_persona+",SolicitudId.EstadoTipoSolicitudId.Id:"+id_estado_tipo_solicitud+"&limit=0", Solicitudes)
	if errSolicitud == nil {
		if *Solicitudes != nil && fmt.Sprintf("%v", (*Solicitudes)[0]) != "map[]" {
			Referencia := (*Solicitudes)[0]["SolicitudId"].(map[string]interface{})["Referencia"].(string)
			var ReferenciaJson map[string]interface{}
			if err := json.Unmarshal([]byte(Referencia), &ReferenciaJson); err == nil {
				if id_estado_tipo_solicitud == "15" {
					return asignacionresultadoSolicitud15GetDatos(resultado, ReferenciaJson, TipoDocumentoGet, errorGetAll, alertas, alerta, errSolicitud)
				} else if id_estado_tipo_solicitud == "16" {
					asignacionresultadoSolicitud16GetDatos(resultado, ReferenciaJson)
				}
			}
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
	return nil
}

// FUNCIONES QUE SE USAN EN GET SOLICITUD ACTUALIZACION DATOS

func solicitudObservacionGetActualizacion(Solicitudes []map[string]interface{}, i int, id_persona string, respuesta *[]map[string]interface{}, TipoSolicitud map[string]interface{}, Estado map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, errEstado error) interface{} {
	IdSolicitud := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"])
	var Observacion []map[string]interface{}
	errObservacion := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"observacion?query=SolicitudId:"+IdSolicitud+",TerceroId:"+id_persona, &Observacion)
	if errObservacion == nil {
		if Observacion != nil && fmt.Sprintf("%v", Observacion[0]) != "map[]" {
			(*respuesta)[i] = map[string]interface{}{
				"Numero":      Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"],
				"Fecha":       Solicitudes[i]["SolicitudId"].(map[string]interface{})["FechaRadicacion"],
				"Tipo":        TipoSolicitud["Data"].(map[string]interface{})["Nombre"],
				"Estado":      Estado["Data"].(map[string]interface{})["Nombre"],
				"Observacion": Observacion[0]["Valor"],
				"TerceroId":   id_persona,
			}
		} else {
			(*respuesta)[i] = map[string]interface{}{
				"Numero":      Solicitudes[i]["SolicitudId"].(map[string]interface{})["Id"],
				"Fecha":       Solicitudes[i]["SolicitudId"].(map[string]interface{})["FechaRadicacion"],
				"Tipo":        TipoSolicitud["Data"].(map[string]interface{})["Nombre"],
				"Estado":      Estado["Data"].(map[string]interface{})["Nombre"],
				"Observacion": "",
				"TerceroId":   id_persona,
			}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errEstado)
		return map[string]interface{}{"Response": *alerta}
	}
	return nil
}

func solicitudEstadoGetActualizacion(Solicitudes []map[string]interface{}, i int, id_persona string, respuesta *[]map[string]interface{}, TipoSolicitud map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert) interface{} {
	IdEstado := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["EstadoTipoSolicitudId"].(map[string]interface{})["EstadoId"].(map[string]interface{})["Id"])
	//Nombre estado de la solicitud
	errEstado := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"estado/"+IdEstado, Estado)
	if errEstado == nil {
		if *Estado != nil && fmt.Sprintf("%v", *Estado) != "map[]" {
			// Observacion (Si la hay) sobre la solicitud
			return solicitudObservacionGetActualizacion(Solicitudes, i, id_persona, respuesta, TipoSolicitud, *Estado, errorGetAll, alertas, alerta, errEstado)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errEstado)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudTipoGetActualizacion(Solicitudes []map[string]interface{}, i int, id_persona string, respuesta *[]map[string]interface{}, TipoSolicitud *map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert) interface{} {
	IdTipoSolicitud := fmt.Sprintf("%v", Solicitudes[i]["SolicitudId"].(map[string]interface{})["EstadoTipoSolicitudId"].(map[string]interface{})["TipoSolicitud"].(map[string]interface{})["Id"])
	//Nombre tipo solicitud
	errTipoSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"tipo_solicitud/"+IdTipoSolicitud, TipoSolicitud)
	if errTipoSolicitud == nil {
		if *TipoSolicitud != nil && fmt.Sprintf("%v", *TipoSolicitud) != "map[]" {
			return solicitudEstadoGetActualizacion(Solicitudes, i, id_persona, respuesta, *TipoSolicitud, Estado, errorGetAll, alertas, alerta)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "No data found", errorGetAll, errTipoSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

func ManejoSolicitudesGetActualizacion(Solicitudes *[]map[string]interface{}, id_persona string, respuesta *[]map[string]interface{}, TipoSolicitud *map[string]interface{}, Estado *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, resultado *map[string]interface{}) interface{} {
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitante?query=TerceroId:"+id_persona+"&sortby=Id&order=asc&limit=0", Solicitudes)
	if errSolicitud == nil {
		if *Solicitudes != nil && fmt.Sprintf("%v", (*Solicitudes)[0]) != "map[]" {
			var data interface{}
			*respuesta = make([]map[string]interface{}, len(*Solicitudes))
			for i := 0; i < len(*Solicitudes); i++ {
				data = solicitudTipoGetActualizacion(*Solicitudes, i, id_persona, respuesta, TipoSolicitud, Estado, errorGetAll, alertas, alerta)
				if data != nil {
					return data
				}
			}
			(*resultado)["Response"] = *respuesta
			return data
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

// FUNCIONES QUE SE USAN EN POST SOLICITUD ACTUALIZACION DATOS

func solicitudEvolucionEstadoPostActualizacion(IdTercero interface{}, IdSolicitud interface{}, IdEstadoTipoSolicitud int, SolicitudJson interface{}, SolicitudEvolucionEstadoPost *map[string]interface{}, resultado *map[string]interface{}, SolicitantePost map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, errSolicitante error) interface{} {
	SolicitudEvolucionEstado := map[string]interface{}{
		"TerceroId": IdTercero,
		"SolicitudId": map[string]interface{}{
			"Id": IdSolicitud,
		},
		"EstadoTipoSolicitudIdAnterior": nil,
		"EstadoTipoSolicitudId": map[string]interface{}{
			"Id": IdEstadoTipoSolicitud,
		},
		"Activo":      true,
		"FechaLimite": fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["FechaSolicitud"]),
	}

	errSolicitudEvolucionEstado := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud_evolucion_estado", "POST", SolicitudEvolucionEstadoPost, SolicitudEvolucionEstado)
	if errSolicitudEvolucionEstado == nil {
		if *SolicitudEvolucionEstadoPost != nil && fmt.Sprintf("%v", *SolicitudEvolucionEstadoPost) != "map[]" {
			(*resultado)["Solicitante"] = SolicitantePost["Data"]
			return nil
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		var resultado2 map[string]interface{}
		request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+fmt.Sprintf("%v", IdSolicitud), "DELETE", &resultado2, nil)
		request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitante/"+fmt.Sprintf("%v", SolicitantePost["Id"]), "DELETE", &resultado2, nil)
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitante)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudPostPostActualizacion(IdTercero interface{}, IdSolicitud interface{}, IdEstadoTipoSolicitud int, SolicitudJson interface{}, SolicitudEvolucionEstadoPost *map[string]interface{}, resultado *map[string]interface{}, SolicitantePost *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert) interface{} {
	Solicitante := map[string]interface{}{
		"TerceroId": IdTercero,
		"SolicitudId": map[string]interface{}{
			"Id": IdSolicitud,
		},
		"Activo": true,
	}

	errSolicitante := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitante", "POST", SolicitantePost, Solicitante)
	if errSolicitante == nil && fmt.Sprintf("%v", (*SolicitantePost)["Status"]) != "400" {
		if *SolicitantePost != nil && fmt.Sprintf("%v", *SolicitantePost) != "map[]" {
			//POST a la tabla solicitud_evolucion estado
			return solicitudEvolucionEstadoPostActualizacion(IdTercero, IdSolicitud, IdEstadoTipoSolicitud, SolicitudJson, SolicitudEvolucionEstadoPost, resultado, *SolicitantePost, errorGetAll, alertas, alerta, errSolicitante)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		//Se elimina el registro de solicitud si no se puede hacer el POST a la tabla solicitante
		var resultado2 map[string]interface{}
		request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+fmt.Sprintf("%v", IdSolicitud), "DELETE", &resultado2, nil)
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitante)
		return map[string]interface{}{"Response": *alerta}
	}
}

func solicitudActualizacionPostActualizacion(IdTercero interface{}, IdEstadoTipoSolicitud int, SolicitudJson interface{}, SolicitudEvolucionEstadoPost *map[string]interface{}, resultado *map[string]interface{}, SolicitantePost *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, SolicitudPost *map[string]interface{}, SolicitudActualizacion map[string]interface{}) interface{} {
	errSolicitud := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud", "POST", SolicitudPost, SolicitudActualizacion)
	if errSolicitud == nil {
		if *SolicitudPost != nil && fmt.Sprintf("%v", *SolicitudPost) != "map[]" {
			(*resultado)["Solicitud"] = (*SolicitudPost)["Data"]
			IdSolicitud := (*SolicitudPost)["Data"].(map[string]interface{})["Id"]

			//POST tabla solicitante
			return solicitudPostPostActualizacion(IdTercero, IdSolicitud, IdEstadoTipoSolicitud, SolicitudJson, SolicitudEvolucionEstadoPost, resultado, SolicitantePost, errorGetAll, alertas, alerta)
		} else {
			ManejoError(alerta, alertas, "No data found", errorGetAll)
			return map[string]interface{}{"Response": *alerta}
		}
	} else {
		ManejoError(alerta, alertas, "", errorGetAll, errSolicitud)
		return map[string]interface{}{"Response": *alerta}
	}
}

func asignarSolicitudActualizacion(SolicitudPadre map[string]interface{}, SolicitudActualizacion *map[string]interface{}, IdEstadoTipoSolicitud int, Referencia string, SolicitudJson interface{}) {
	*SolicitudActualizacion = map[string]interface{}{
		"EstadoTipoSolicitudId": map[string]interface{}{"Id": IdEstadoTipoSolicitud},
		"Referencia":            Referencia,
		"Resultado":             "",
		"FechaRadicacion":       fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["FechaSolicitud"]),
		"Activo":                true,
		"SolicitudPadreId":      SolicitudPadre,
	}
}

func ManejoSolicitudesPostActualizacion(IdEstadoTipoSolicitud int, SolicitudEvolucionEstadoPost *map[string]interface{}, resultado *map[string]interface{}, SolicitantePost *map[string]interface{}, errorGetAll *bool, alertas *[]interface{}, alerta *models.Alert, SolicitudPost *map[string]interface{}, Solicitud map[string]interface{}, Referencia string, SolicitudPadre *map[string]interface{}) interface{} {
	IdTercero := Solicitud["Solicitante"]
	SolicitudJson := Solicitud["Solicitud"]
	TipoSolicitud := Solicitud["TipoSolicitud"]
	f, _ := strconv.ParseFloat(fmt.Sprintf("%v", TipoSolicitud), 64)
	j, _ := strconv.Atoi(fmt.Sprintf("%v", f))
	if j == 3 {
		//Tipo de solicitud de actualización de datos por ID
		Referencia = "{\n\"DocumentoId\":" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["Documento"]) + ",\n\"DatosAnteriores\": {\n\"FechaExpedicionActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["FechaExpedicionActual"]) + "\", \n\"NumeroActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NumeroActual"]) + "\",\n\"TipoDocumentoActual\": {\n\"Id\": " + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["TipoDocumentoActual"].(map[string]interface{})["Id"]) + "\n}\n}, \n\"DatosNuevos\": {\n\"FechaExpedicionNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["FechaExpedicionNuevo"]) + "\",\n\"NumeroNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NumeroNuevo"]) + "\",\n\"TipoDocumentoNuevo\": {\n\"Id\": " + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["TipoDocumentoNuevo"].(map[string]interface{})["Id"]) + "\n}\n}\n}"
		IdEstadoTipoSolicitud = 15
		if Solicitud["SolicitudPadreId"] != nil {
			IdEstadoTipoSolicitud = 33
		}
	} else if j == 4 {
		//Tipo de solicitud de actualización de datos por nombre
		Referencia = "{\n\"DocumentoId\":" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["Documento"]) + ",\n\"DatosAnteriores\":{\n\"NombreActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NombreActual"]) + "\",\n\"ApellidoActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["ApellidoActual"]) + "\"\n},\n\"DatosNuevos\":{\n\"NombreNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NombreNuevo"]) + "\",\n\"ApellidoNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["ApellidoNuevo"]) + "\"\n}\n}"
		IdEstadoTipoSolicitud = 16
		if Solicitud["SolicitudPadreId"] != nil {
			IdEstadoTipoSolicitud = 32
		}
	}
	SolicitudActualizacion := map[string]interface{}{}
	IdSolicutudPadre := string("")

	if Solicitud["SolicitudPadreId"] != nil {
		IdSolicutudPadre = Solicitud["SolicitudPadreId"].(string)
		errSolicitudPadre := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+IdSolicutudPadre, SolicitudPadre)
		if errSolicitudPadre == nil {
			//POST tabla solicitud
			asignarSolicitudActualizacion(*SolicitudPadre, &SolicitudActualizacion, IdEstadoTipoSolicitud, Referencia, SolicitudJson)
		} else {
			//POST tabla solicitud
			asignarSolicitudActualizacion(nil, &SolicitudActualizacion, IdEstadoTipoSolicitud, Referencia, SolicitudJson)
		}
	} else {
		asignarSolicitudActualizacion(nil, &SolicitudActualizacion, IdEstadoTipoSolicitud, Referencia, SolicitudJson)
	}

	return solicitudActualizacionPostActualizacion(IdTercero, IdEstadoTipoSolicitud, SolicitudJson, SolicitudEvolucionEstadoPost, resultado, SolicitantePost, errorGetAll, alertas, alerta, SolicitudPost, SolicitudActualizacion)
}

// FUNCIONES QUE SE USAN EN POST SOLICITUD ACTUALIZACION DATOS

func manejoSolicitudRejectPutSolicitud(solicitudEvaluacion map[string]interface{}, idSolicitud string, resultadoPutSolicitud map[string]interface{}, resultadoRechazo map[string]interface{}) (interface{}, interface{}) {
	if solicitudReject, errPrepared := models.PreparedRejectState(solicitudEvaluacion); errPrepared == nil {
		if resultado, errPut := models.PutSolicitudDocente(solicitudReject, idSolicitud); errPut == nil {
			resultadoPutSolicitud = resultado
			mensaje := "La invitación ha sido rechazada, por favor cierre la pestaña o ventana"
			resultadoRechazo["Resultado"] = map[string]interface{}{
				"Mensaje": mensaje,
			}
			return nil, resultadoRechazo
		} else {
			logs.Error(errPut)
			return resultadoPutSolicitud, nil
		}
	} else {
		logs.Error(errPrepared)
		return resultadoPutSolicitud, nil
	}
}

func ManejoSolicitudes(solicitudEvaluacion map[string]interface{}, solicitudEvaluacionList []interface{}, resultadoRechazo map[string]interface{}, idSolicitud string, resultadoPutSolicitud map[string]interface{}) (interface{}, interface{}) {
	solicitudEvaluacion = solicitudEvaluacionList[0].(map[string]interface{})
	if fmt.Sprintf("%v", solicitudEvaluacion["EstadoTipoSolicitudId"].(map[string]interface{})["EstadoId"].(map[string]interface{})["Id"]) == "11" {
		mensaje := "La invitación ya ha sido rechazada anteriormente, por favor cierre la pestaña o ventana"
		resultadoRechazo["Resultado"] = map[string]interface{}{
			"Mensaje": mensaje,
		}
		return nil, resultadoRechazo
	} else {
		if errorSystem, dataJson := manejoSolicitudRejectPutSolicitud(solicitudEvaluacion, idSolicitud, resultadoPutSolicitud, resultadoRechazo); errorSystem == nil {
			return nil, dataJson
		} else {
			return errorSystem, nil
		}
	}
}

// FUNCIONES QUE SE USAN EN VARIOS ENDPOINTS

func ManejoError(alerta *models.Alert, alertas *[]interface{}, mensaje string, errorGetAll *bool, err ...error) {
	var msj string
	if len(err) > 0 && err[0] != nil {
		msj = mensaje + err[0].Error()
	} else {
		msj = mensaje
	}
	*errorGetAll = true
	*alertas = append(*alertas, msj)
	(*alerta).Body = *alertas
	(*alerta).Type = "error"
	(*alerta).Code = "400"
}

func ManejoExito(alerta *models.Alert, alertas *[]interface{}, resultado map[string]interface{}) {
	*alertas = append(*alertas, resultado)
	(*alerta).Code = "200"
	(*alerta).Type = "OK"
	(*alerta).Body = *alertas
}
