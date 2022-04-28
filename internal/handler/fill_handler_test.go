package handler

import (
	"net/http"
	"testing"
)

//@TODO fixme.
func TestHandlers_FillHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{}
			h.FillHandler(tt.args.w, tt.args.r)
		})
	}
}
