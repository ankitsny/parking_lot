package models

import (
	"reflect"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	type args struct {
		lotID   string
		address string
		size    int
	}
	tests := []struct {
		name    string
		args    args
		want    *ParkingLot
		wantErr bool
	}{
		{
			name: "Check for invalid size",
			args: args{
				address: "",
				size:    0,
				lotID:   "LOT_1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateParkingLot(tt.args.lotID, tt.args.address, tt.args.size)
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
