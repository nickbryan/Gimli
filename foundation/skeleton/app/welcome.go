package app

import "net/http"

type WelcomeHandler struct {
	Message string
}

func (handler *WelcomeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(handler.Message))
}
