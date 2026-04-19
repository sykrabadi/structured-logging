package server

import (
	"fmt"
	"net/http"
	"structured-logging/internal/service"
	"structured-logging/utils/log"
)

func NewLoggingMiddleware(next func(http.ResponseWriter, *http.Request) error, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoFromCtx(r.Context())

		err := next(w, r)
		if err != nil {
			logger.ErrorStack(err)
		}
	}
}

type Handler struct{
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler{
	return &Handler{
		service: svc,
	}
}

func RunHTTPServer(
	address string,
	svc *service.Service,
	logger *log.Logger,
) {
	logger.Info(fmt.Sprintf("%v\n", address))

	h := NewHandler(svc)
	http.Handle("/home", NewLoggingMiddleware(h.homeHandler, logger))
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Error(err)
	}
}
