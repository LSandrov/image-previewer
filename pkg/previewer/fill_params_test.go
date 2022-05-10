package previewer

import (
	"context"
	"reflect"
	"testing"
)

func TestNewFillParams(t *testing.T) {
	type args struct {
		ctx     context.Context
		width   int
		height  int
		url     string
		headers map[string][]string
	}
	tests := []struct {
		name string
		args args
		want *FillParams
	}{
		{
			name: "good",
			want: &FillParams{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFillParams(tt.args.ctx, tt.args.width, tt.args.height, tt.args.url, tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFillParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
