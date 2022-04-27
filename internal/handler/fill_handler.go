package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//@TODO fixme
	print(vars)
}
