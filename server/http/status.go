package http

import "net/http"

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
