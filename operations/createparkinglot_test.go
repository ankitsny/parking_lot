package operations

import (
	"reflect"
	"testing"
)

func TestNewCreateParkingLot(t *testing.T) {
	tests := []struct {
		name string
		want *CreateParkingLot
	}{
		{
			name: "new parking log",
			want: &CreateParkingLot{
				opName: "create_parking_lot",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateParkingLot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateParkingLot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateParkingLot_GetName(t *testing.T) {
	type fields struct {
		opName   string
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Name",
			fields: fields{
				opName: "create_parking_lot",
			},
			want: "create_parking_lot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpl := &CreateParkingLot{
				opName:   tt.fields.opName,
				capacity: tt.fields.capacity,
			}
			if got := cpl.GetName(); got != tt.want {
				t.Errorf("CreateParkingLot.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
