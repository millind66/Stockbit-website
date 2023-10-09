package api

import "net/http"

type dummyWriter struct{}

func (*dummyWriter) Header() http.Header {
	return map[string][]string{}
}

func (*dummyWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (*dummyWriter) WriteHeader(statusCode int) {}
