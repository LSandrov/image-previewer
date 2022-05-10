package handler

import (
	"reflect"
	"testing"

	"github.com/LSandrov/image-previewer/pkg/previewer"

	"github.com/rs/zerolog"
)

func TestNewHandlers(t *testing.T) {
	type args struct {
		l   zerolog.Logger
		svc previewer.Service
	}
	tests := []struct {
		name string
		args args
		want *Handlers
	}{
		{
			name: "good",
			want: &Handlers{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlers(tt.args.l, tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
