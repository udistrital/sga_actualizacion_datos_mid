package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_actualizacion_dato_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/requestresponse"
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
		"FechaExpedicion": ReferenciaJson["DatosNuevos"].(map[string]interface{})["FechaExpedicionNuevo"].(string) + "T00:00:00Z",
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
				case 17.0: // Rectificar -> Modificada
					EstadoTipoSolicitudId = 33
				default:
					return map[string]interface{}{"Response": "Tipo de solicitud no permitido"}
				}
			} else if TipoSolicitudId == 4 {
				//El tipo de solicitud es de cambio de nombre
				switch Estado {
				case 9.0: // Aprobado
					EstadoTipoSolicitudId = 18
				case 11.0: // Rechazado
					EstadoTipoSolicitudId = 19
				case 17.0: // Rectificar -> Modificada
					EstadoTipoSolicitudId = 32
				default:
					return map[string]interface{}{"Response": "Tipo de solicitud no permitido"}
				}
			} else {
				return map[string]interface{}{"Response": "Tipo de solicitud no permitido"}
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
			// Filtrar solicitudes activas de actualización de datos
			var solicitudesActivas []map[string]interface{}
			for _, solicitud := range *Solicitudes {
				if solicitudId, ok := solicitud["SolicitudId"].(map[string]interface{}); ok {
					if activo, ok := solicitudId["Activo"].(bool); ok && activo {
						solicitudesActivas = append(solicitudesActivas, solicitud)
					}
				}
			}

			if len(solicitudesActivas) == 0 {
				ManejoError(alerta, alertas, "No active data found", errorGetAll)
				return map[string]interface{}{"Response": *alerta}
			}

			var data interface{}
			*respuesta = make([]map[string]interface{}, len(solicitudesActivas))
			for i := 0; i < len(solicitudesActivas); i++ {
				data = solicitudTipoGetActualizacion(solicitudesActivas, i, id_persona, respuesta, TipoSolicitud, Estado, errorGetAll, alertas, alerta)
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

			// Verifica si la solicitud tiene una solicitud padre asociada
			if SolicitudPadre, ok := SolicitudActualizacion["SolicitudPadreId"].(map[string]interface{}); ok && len(SolicitudPadre) > 0 {
				IdSolicitudPadre := fmt.Sprintf("%v", SolicitudPadre["Id"])
				// Inactiva la solicitud padre
				SolicitudPadre["Activo"] = false
				// Hace el put para inactivar la solicitud padre
				var respuestaUpdate map[string]interface{}
				errSolicitudPadreUpdate := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+IdSolicitudPadre, "PUT", &respuestaUpdate, SolicitudPadre)
				if errSolicitudPadreUpdate != nil {
					ManejoError(alerta, alertas, "", errorGetAll, errSolicitudPadreUpdate)
					return map[string]interface{}{"Response": *alerta}
				}
			}

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
	} else if j == 4 {
		//Tipo de solicitud de actualización de datos por nombre
		Referencia = "{\n\"DocumentoId\":" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["Documento"]) + ",\n\"DatosAnteriores\":{\n\"NombreActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NombreActual"]) + "\",\n\"ApellidoActual\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["ApellidoActual"]) + "\"\n},\n\"DatosNuevos\":{\n\"NombreNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["NombreNuevo"]) + "\",\n\"ApellidoNuevo\": \"" + fmt.Sprintf("%v", SolicitudJson.(map[string]interface{})["ApellidoNuevo"]) + "\"\n}\n}"
		IdEstadoTipoSolicitud = 16
	}
	SolicitudActualizacion := map[string]interface{}{}
	IdSolicutudPadre := string("")

	if Solicitud["SolicitudPadreId"] != nil {
		IdSolicutudPadre = Solicitud["SolicitudPadreId"].(string)
		errSolicitudPadre := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+IdSolicutudPadre, SolicitudPadre)
		if errSolicitudPadre == nil {
			//POST tabla solicitud con solicitud padre asociada
			asignarSolicitudActualizacion(*SolicitudPadre, &SolicitudActualizacion, IdEstadoTipoSolicitud, Referencia, SolicitudJson)
		} else {
			//POST tabla solicitud sin solicitud padre asociada
			asignarSolicitudActualizacion(nil, &SolicitudActualizacion, IdEstadoTipoSolicitud, Referencia, SolicitudJson)
		}
	} else {
		//POST tabla solicitud sin solicitud padre asociada
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

func DatosSolicitud(id_solicitud string) (APIResponseDTO requestresponse.APIResponse) {
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
					ConfigurarResultadoGetSolicitudId(&resultado, &ReferenciaJson, 1)

					if respuestaTipo := SolicitudTipoDocGetSolicitudId(TipoDocumento, &TipoDocumentoGet, &resultado, &alerta, &alertas, &errorGetAll); respuestaTipo != nil {
						APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuestaTipo)
					}

					if respuestaDoc := SolicitudDocActualGetSolicitudId(ReferenciaJson, &TipoDocumentoActualGet, &resultado, &alerta, &alertas, &errorGetAll); respuestaDoc != nil {
						APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuestaDoc)
					}
				} else if TipoSolicitudId == 16 || TipoSolicitudId == 18 || TipoSolicitudId == 19 || TipoSolicitudId == 32 {
					ConfigurarResultadoGetSolicitudId(&resultado, &ReferenciaJson, 2)
				}
			}
		} else {
			ManejoError(&alerta, &alertas, "No data found", &errorGetAll)
			APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, errorGetAll)
		}
	} else {
		ManejoError(&alerta, &alertas, "", &errorGetAll, errSolicitud)
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, errSolicitud.Error())
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, resultado, nil)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func SolicitudEvolucion(data []byte) (APIResponseDTO requestresponse.APIResponse) {

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

	if err := json.Unmarshal(data, &Solicitud); err == nil {
		if respuesta := SolicitudEstadoPostSolicitud(&resultado, &SolicitudEvolucionEstadoPost, &SolicitudAux, EstadoTipoSolicitudId, &SolicitudAuxPost, Solicitud, &ObservacionPost, &errorGetAll, &alerta, &alertas, &SolicitudAprob, &Tercero, &TerceroPut, &DatosIdentificacion, &DatosIdentificacionPut, &DatosIdentificacionPost, &SolicitudEvolucionEstado); respuesta != nil {
			APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuesta)
		}
	} else {
		ManejoError(&alerta, &alertas, "", &errorGetAll, err)
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, err.Error())
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, nil, resultado)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func SolicitudActualizacionDatos(id_estado_tipo_sol string) (APIResponseDTO requestresponse.APIResponse) {
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

	if resp := ManejoSolicitudesGetAll(&Solicitudes, Observacion, &respuesta, &TipoSolicitud, &Estado, &errorGetAll, &alertas, &alerta, id_estado_tipo_sol, &resultado); resp != nil {
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, resp)
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, resultado)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func GetDatosSolicitud(id_persona string, id_estado_tipo_solicitud string) (APIResponseDTO requestresponse.APIResponse) {
	var Solicitudes []map[string]interface{}
	var TipoDocumentoGet map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	if respuesta := SolicitudGetDatos(&resultado, &TipoDocumentoGet, &errorGetAll, &alertas, &alerta, id_persona, id_estado_tipo_solicitud, &Solicitudes); respuesta != nil {
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuesta)
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, resultado)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func GetSolictudActualizacion(id_persona string) (APIResponseDTO requestresponse.APIResponse) {
	var Solicitudes []map[string]interface{}
	var TipoSolicitud map[string]interface{}
	var Estado map[string]interface{}
	var respuesta []map[string]interface{}
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})
	var alerta models.Alert
	var errorGetAll bool
	alertas := append([]interface{}{})

	if respuesta := ManejoSolicitudesGetActualizacion(&Solicitudes, id_persona, &respuesta, &TipoSolicitud, &Estado, &errorGetAll, &alertas, &alerta, &resultado); respuesta != nil {
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuesta)
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, resultado)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func ActualizacionDatosPost(data []byte) (APIResponseDTO requestresponse.APIResponse) {
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

	if err := json.Unmarshal(data, &Solicitud); err == nil {
		if respuesta := ManejoSolicitudesPostActualizacion(IdEstadoTipoSolicitud, &SolicitudEvolucionEstadoPost, &resultado, &SolicitantePost, &errorGetAll, &alertas, &alerta, &SolicitudPost, Solicitud, Referencia, &SolicitudPadre); respuesta != nil {
			APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, respuesta)
		}
	} else {
		APIResponseDTO = requestresponse.APIResponseDTO(false, 404, nil, err.Error())
	}

	if !errorGetAll {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, resultado)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func SolicitudEvaluacionPut(idSolicitud string) (APIResponseDTO requestresponse.APIResponse) {
	//resultado resultado final
	var resultadoPutSolicitud map[string]interface{}
	resultadoRechazo := make(map[string]interface{})

	var solicitudEvaluacion map[string]interface{}
	if solicitudEvaluacionList, errGet := models.GetOneSolicitudDocente(idSolicitud); errGet == nil {
		if errorSystem, dataJson := ManejoSolicitudes(solicitudEvaluacion, solicitudEvaluacionList, resultadoRechazo, idSolicitud, resultadoPutSolicitud); errorSystem == nil {
			APIResponseDTO = requestresponse.APIResponseDTO(true, 200, dataJson)
		} else {
			APIResponseDTO = requestresponse.APIResponseDTO(false, 400, nil, errorSystem)
			return APIResponseDTO
		}
	} else {
		logs.Error(errGet)
		APIResponseDTO = requestresponse.APIResponseDTO(false, 400, nil, errGet)
		return APIResponseDTO
	}
	return APIResponseDTO
}

func PutSolicitudReferencia(idSolicitud string, referencia map[string]interface{}) (APIResponseDTO requestresponse.APIResponse) {
	// Consultar solicitud por id
	var solicitud map[string]interface{}
	errSolicitud := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+idSolicitud, &solicitud)
	if errSolicitud != nil {
		return requestresponse.APIResponseDTO(false, 404, nil, errSolicitud)
	}
	//17 Rectificar
	//9 Acta aprobnada
	//1 Solicitud generada
	//11 Solicitud rechazada
	var estadoSolicitudGenerada int = 1
	esUnaSolicitudModificable, err := solicitudTieneEstado(solicitud, estadoSolicitudGenerada)
	if err != nil {
		return requestresponse.APIResponseDTO(false, 400, nil, err)
	}

	if esUnaSolicitudModificable {

		// Formatear la referencia a JSON
		referenciaJson, errReferenciaJson := formatReferenciaJson(referencia)
		if errReferenciaJson != nil {
			fmt.Println("Error:", errReferenciaJson)
			return requestresponse.APIResponseDTO(false, 400, nil, "error en la funcion del formatReferenciaJson")
		}

		// Reemplazar la referencia en la solicitud con el JSON string formateado
		solicitud["Referencia"] = referenciaJson

		fmt.Print("------------------------------")
		jsonData, _ := json.MarshalIndent(solicitud, "", "")
		fmt.Println(string(jsonData))
		fmt.Println("---------------------------")

		// Guardar la solicitud
		var responseSolicitud map[string]interface{}
		err := request.SendJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"solicitud/"+idSolicitud, "PUT", &responseSolicitud, solicitud)
		if err != nil {
			fmt.Println(err)
			return requestresponse.APIResponseDTO(false, 404, nil, err)
		}

		// Retornar respuesta exitosa
		return requestresponse.APIResponseDTO(true, 200, nil)
	} else {
		return requestresponse.APIResponseDTO(false, 200, nil, "La solicitud solo se puede modificar si esta en estado 'Solicitud generada' o 'Rectificar'")
	}

}

func solicitudTieneEstado(solicitud map[string]interface{}, estadoId int) (bool, error) {
	estadoTipoSolicitudIdRequest, ok := solicitud["EstadoTipoSolicitudId"].(map[string]interface{})["Id"]
	if !ok {
		return false, errors.New("invalid solicitud format: missing EstadoTipoSolicitudId")
	}
	estadoTipoSolicitudId, _ := strconv.Atoi(fmt.Sprintf("%v", estadoTipoSolicitudIdRequest))
	// Consultar por id a la tabla estado_tipo_solicitud del modelo "solicitudes" para obtener el estado de la solicitud
	var estadoTipoSolicitudRequest map[string]interface{}
	errEstadoTipoSolicitudRequest := request.GetJson("http://"+beego.AppConfig.String("SolicitudDocenteService")+"estado_tipo_solicitud/"+strconv.Itoa(estadoTipoSolicitudId), &estadoTipoSolicitudRequest)
	if errEstadoTipoSolicitudRequest != nil {
		return false, errEstadoTipoSolicitudRequest
	}
	estadoIdRequest, ok := estadoTipoSolicitudRequest["Data"].(map[string]interface{})["EstadoId"].(map[string]interface{})["Id"]
	if !ok {
		return false, errors.New("invalid solicitud format: missing EstadoId")
	}
	id, _ := strconv.Atoi(fmt.Sprintf("%v", estadoIdRequest))
	if id != estadoId {
		return false, nil
	} else {
		return true, nil
	}
}

func formatReferenciaJson(referencia map[string]interface{}) (string, error) {
	// Serializar el mapa a JSON con indentación
	referenciaJsonBytes, err := json.MarshalIndent(referencia, "", "    ")
	if err != nil {
		return "", err
	}

	// Convertir los bytes a string
	referenciaJson := string(referenciaJsonBytes)

	//Eliminar los espacios en blanco
	referenciaJsonSinEspacios := strings.ReplaceAll(referenciaJson, " ", "")

	return referenciaJsonSinEspacios, nil
}
