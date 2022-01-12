package http

import "net/http"

type rebalancedHandler struct {
	*Server
}

func (h *rebalancedHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
