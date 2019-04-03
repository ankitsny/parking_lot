package modes

import (
	"reflect"
	"testing"
)

func TestNewInterActiveMode(t *testing.T) {
	tests := []struct {
		name string
		want *InteractiveMode
	}{
		{
			name: "1",
			want: &InteractiveMode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterActiveMode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInterActiveMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Infinte loop
// func TestInteractiveMode_Start(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		im   *InteractiveMode
// 	}{
// 		{
// 			name: "1",
// 			im:   &InteractiveMode{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			im := &InteractiveMode{}
// 			im.Start()
// 		})
// 	}
// }
