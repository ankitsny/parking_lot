package operations

import (
	"reflect"
	"testing"
)

func TestNewOperation(t *testing.T) {
	tests := []struct {
		name string
		want IOperation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
