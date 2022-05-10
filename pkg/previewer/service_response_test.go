package previewer

import (
	"reflect"
	"testing"
)

func TestNewFillResponse(t *testing.T) {
	type args struct {
		img     []byte
		headers map[string][]string
	}
	tests := []struct {
		name string
		args args
		want *FillResponse
	}{
		{
			name: "good",
			want: &FillResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFillResponse(tt.args.img, tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFillResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
