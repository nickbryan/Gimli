package controllers

import (
	"net/http"

	"github.com/nickbryan/gimli/foundation/skeleton/app"
)

type WelcomeController struct {
	Printer app.PrinterService
}

func (controller *WelcomeController) Welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(controller.Printer.Render()))
}
