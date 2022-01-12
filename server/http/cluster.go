package http

import "net/http"

type clusterHandler struct {
	*Server
}

func (h *clusterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
