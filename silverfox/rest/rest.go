package rest

import "golang.org/x/net/context"

const (
	Options = "OPTIONS"
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
)

type RestHandler interface {
	Options() AccessControlHeaders
	Get(param *UrlParam) interface{}
	Post(param *UrlParam) interface{}
	Put(param *UrlParam, json JsonParam) interface{}
	Delete(param *UrlParam) interface{}
	SetContext(context context.Context)
}

const (
	allowOrigin        = "Access-Control-Allow-Origin"
	allowHeaders       = "Access-Control-Allow-Headers"
	allowMethods       = "Access-Control-Allow-Methods"
	allowExposeHeaders = "Access-Control-Expose-Headers"
)

type AccessControlHeaders struct {
	headers map[string]string
}

func NewAccessControlHeaders() AccessControlHeaders {
	headers := AccessControlHeaders{headers: map[string]string{}}
	return headers
}

func (s *AccessControlHeaders) AllowOrigin(value string) {
	s.headers[allowOrigin] = value
}

func (s *AccessControlHeaders) AllowHeaders(value string) {
	s.headers[allowHeaders] = value
}

func (s *AccessControlHeaders) AllowMethods(value string) {
	s.headers[allowMethods] = value
}

func (s *AccessControlHeaders) AllowMethodsAll() {
	s.headers[allowMethods] = Get + "," + Post + "," + Put + "," + Delete
}

func (s *AccessControlHeaders) AllowExposeHeaders(value string) {
	s.headers[allowExposeHeaders] = value
}
func (s *AccessControlHeaders) Get() map[string]string {
	return s.headers
}
