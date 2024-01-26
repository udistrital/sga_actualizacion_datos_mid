package controllers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_mid_actualizacion_datos/utils"
)

type ErrorHandlerController struct {
	beego.Controller
}

func (c *ErrorHandlerController) Error404() {
    metodo := c.Ctx.Request.Method
    ruta := c.Ctx.Request.URL.Path
    statusCode := http.StatusNotFound

    message := fmt.Sprintf("nomatch|%s|%s", metodo, ruta)

    c.Ctx.Output.SetStatus(statusCode)
    c.Data["json"] = utils.APIResponseDTO(false, statusCode, nil, message)
    c.ServeJSON()
}

func HandlePanic(c *beego.Controller) {
	if r := recover(); r != nil {
		logs.Error(r)
		debug.PrintStack()
		statusCode := http.StatusInternalServerError
		message := "Error service " + beego.AppConfig.String("appname") + ": An internal server error occurred"
		c.Ctx.Output.SetStatus(statusCode)
		c.Data["json"] = utils.APIResponseDTO(false, statusCode, fmt.Sprintf("%v", r), message)
		c.ServeJSON()
	}
}