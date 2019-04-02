package models

import (
	"reflect"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	var parkingSpots []*ParkingSpot
	parkingSpots = append(parkingSpots, NewEmptyParkingSpot("SM", 1, "0"))
	type args struct {
		lotID    string
		address  string
		capacity int
	}
	tests := []struct {
		name    string
		args    args
		want    *ParkingLot
		wantErr bool
	}{
		{
			name: "Check for invalid capacity",
			args: args{
				address:  "",
				capacity: 0,
				lotID:    "LOT_1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "With valid inputs",
			args: args{
				address:  "Bangalore HSR",
				capacity: 1,
				lotID:    "LOT_1",
			},
			want: &ParkingLot{
				address:       "Bangalore HSR",
				capacity:      1,
				lotID:         "LOT_1",
				hasEmptySpace: true,
				nextSpot:      0,
				parkingSpots:  parkingSpots,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateParkingLot(tt.args.lotID, tt.args.address, tt.args.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateParkingLot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParkingLot() = %v, want %v", got, tt.want)
			}
		})
	}
}
