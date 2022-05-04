package handler

import (
	"github.com/rs/zerolog"
	"image-previewer/pkg/previewer"
	"net/http"
	"reflect"
	"testing"
)

func TestHandlers_FillHandler(t *testing.T) {
	type fields struct {
		l   zerolog.Logger
		svc previewer.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				l:   tt.fields.l,
				svc: tt.fields.svc,
			}
			h.FillHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlers_parseFillHandlerVars(t *testing.T) {
	type fields struct {
		l   zerolog.Logger
		svc previewer.Service
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   *FillRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				l:   tt.fields.l,
				svc: tt.fields.svc,
			}
			gotR, err := h.parseFillHandlerVars(tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFillHandlerVars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("parseFillHandlerVars() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
