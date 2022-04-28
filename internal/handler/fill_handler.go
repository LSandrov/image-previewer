package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//@TODO fixme
	print(vars)
}
