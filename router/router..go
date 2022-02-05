package router

import (
	"github.com/gorilla/mux"
)

rootPath := "https://www.bungie.net/Platform"
func Router(addr string, err error) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/")
}
