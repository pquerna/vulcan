package vmod

import (
	"net/http"
)

type Request struct {
	http.Request
	// TODO: more fields to make Jump() work.
}

func (vreq *Request) Jump(stackName string) {
	// TODO: jump to another stack
}

type ResponseWriter struct {
	written bool
	http.ResponseWriter
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.written = true
	w.ResponseWriter.WriteHeader(status)
}

type HandlerFunc func(w *ResponseWriter, r *Request) error

type Handler interface {
	Name() string
	// TODO: figure out config params
	//	Configure (interface{}) (HandlerFunc, error)
	Configure() (HandlerFunc, error)
}

type Stack struct {
	Name     string
	Handlers []HandlerFunc
}

func NewStack(name string, funcs ...HandlerFunc) *Stack {
	return &Stack{Name: name, Handlers: funcs}
}

type Registry struct {
	reg map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (vr *Registry) Add(h Handler) {
	vr.reg[h.Name()] = h
}

func (vr *Registry) Get(name string) Handler {
	return vr.reg[name]
}

var global_registry *Registry

func Register(h Handler) {
	if global_registry == nil {
		global_registry = NewRegistry()
	}
	global_registry.Add(h)
}
