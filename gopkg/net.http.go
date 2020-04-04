// +build !appengine

package gopkg

import (
	"net/http"
	"reflect"
)

func init() {
	Packages["net/http"] = map[string]reflect.Value{
		"DefaultClient":     reflect.ValueOf(http.DefaultClient),
		"DefaultServeMux":   reflect.ValueOf(http.DefaultServeMux),
		"DefaultTransport":  reflect.ValueOf(http.DefaultTransport),
		"Handle":            reflect.ValueOf(http.Handle),
		"HandleFunc":        reflect.ValueOf(http.HandleFunc),
		"ListenAndServe":    reflect.ValueOf(http.ListenAndServe),
		"ListenAndServeTLS": reflect.ValueOf(http.ListenAndServeTLS),
		"NewRequest":        reflect.ValueOf(http.NewRequest),
		"NewServeMux":       reflect.ValueOf(http.NewServeMux),
		"Serve":             reflect.ValueOf(http.Serve),
		"SetCookie":         reflect.ValueOf(http.SetCookie),
	}
	PackageTypes["net/http"] = map[string]reflect.Type{
		"Client":   reflect.TypeOf(http.Client{}),
		"Cookie":   reflect.TypeOf(http.Cookie{}),
		"Request":  reflect.TypeOf(http.Request{}),
		"Response": reflect.TypeOf(http.Response{}),
	}
}
