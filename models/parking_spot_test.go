package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewEmptyParkingSpot(t *testing.T) {
	type args struct {
		size   string
		spotNo int
		level  string
	}
	tests := []struct {
		name string
		args args
		want *ParkingSpot
	}{
		{
			name: "New Empty parking spot",
			args: args{
				level:  "0",
				size:   "SM",
				spotNo: 1,
			},
			want: &ParkingSpot{
				level:  "0",
				size:   "SM",
				spotNo: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyParkingSpot(tt.args.size, tt.args.spotNo, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyParkingSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSpot_GetVehicle(t *testing.T) {
	type fields struct {
		size        string
		spotNo      int
		vehicle     *Vehicle
		parkingTime *time.Time
		level       string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vehicle
	}{
		{
			name: "ParkingSpot_GetVehicle",
			fields: fields{
				level:       "0",
				spotNo:      1,
				vehicle:     nil,
				parkingTime: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSpot{
				size:        tt.fields.size,
				spotNo:      tt.fields.spotNo,
				vehicle:     tt.fields.vehicle,
				parkingTime: tt.fields.parkingTime,
				level:       tt.fields.level,
			}
			if got := ps.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkingSpot.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSpot_ParkVehicle(t *testing.T) {
	parkingTime := time.Now()
	type fields struct {
		size        string
		spotNo      int
		vehicle     *Vehicle
		parkingTime *time.Time
		level       string
	}
	type args struct {
		v *Vehicle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Park Vehicle",
			fields: fields{
				level:       "0",
				parkingTime: &parkingTime,
				size:        "SM",
				spotNo:      1,
				vehicle: &Vehicle{
					color:     "Red",
					size:      "SM",
					vehicleNo: "KA-51EZ-1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSpot{
				size:        tt.fields.size,
				spotNo:      tt.fields.spotNo,
				vehicle:     tt.fields.vehicle,
				parkingTime: tt.fields.parkingTime,
				level:       tt.fields.level,
			}
			ps.ParkVehicle(tt.args.v)
		})
	}
}

func TestParkingSpot_VackateVehicle(t *testing.T) {
	type fields struct {
		size        string
		spotNo      int
		vehicle     *Vehicle
		parkingTime *time.Time
		level       string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSpot{
				size:        tt.fields.size,
				spotNo:      tt.fields.spotNo,
				vehicle:     tt.fields.vehicle,
				parkingTime: tt.fields.parkingTime,
				level:       tt.fields.level,
			}
			ps.VackateVehicle()
		})
	}
}
