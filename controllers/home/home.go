package home

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func SayHello(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Write([]byte("Hello world"))
}
