package gokit

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// 进行路径匹配的多路复用器
	m := http.NewServeMux()

	m.Handle("/hash", httptransport.NewServer(
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))

	return m
}
