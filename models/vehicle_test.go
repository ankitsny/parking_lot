package models

import (
	"reflect"
	"testing"
)

func TestCreateVehicle(t *testing.T) {
	type args struct {
		color     string
		vehicleNo string
		size      string
	}
	tests := []struct {
		name string
		args args
		want *Vehicle
	}{
		{
			name: "Create Vehicle Test",
			args: args{
				color:     "Red",
				size:      "SM",
				vehicleNo: "KA-51EZ-1234",
			},
			want: &Vehicle{
				color:     "Red",
				size:      "SM",
				vehicleNo: "KA-51EZ-1234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateVehicle(tt.args.color, tt.args.vehicleNo, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVehicle_GetColor(t *testing.T) {
	type fields struct {
		color     string
		vehicleNo string
		size      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Color",
			fields: fields{
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
				size:      "SM",
			},
			want: "Red",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vehicle{
				color:     tt.fields.color,
				vehicleNo: tt.fields.vehicleNo,
				size:      tt.fields.size,
			}
			if got := v.GetColor(); got != tt.want {
				t.Errorf("Vehicle.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVehicle_GetVehicleNo(t *testing.T) {
	type fields struct {
		color     string
		vehicleNo string
		size      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Vehicle No",
			fields: fields{
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
				size:      "SM",
			},
			want: "KA-51EZ-1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vehicle{
				color:     tt.fields.color,
				vehicleNo: tt.fields.vehicleNo,
				size:      tt.fields.size,
			}
			if got := v.GetVehicleNo(); got != tt.want {
				t.Errorf("Vehicle.GetVehicleNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVehicle_GetSize(t *testing.T) {
	type fields struct {
		color     string
		vehicleNo string
		size      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Vehicle No",
			fields: fields{
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
				size:      "SM",
			},
			want: "SM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vehicle{
				color:     tt.fields.color,
				vehicleNo: tt.fields.vehicleNo,
				size:      tt.fields.size,
			}
			if got := v.GetSize(); got != tt.want {
				t.Errorf("Vehicle.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
