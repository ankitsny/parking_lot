package store

import (
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/models"
)

func TestNewStore(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewStore(tt.args.capacity); (err != nil) != tt.wantErr {
				t.Errorf("NewStore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetStorage(t *testing.T) {
	tests := []struct {
		name string
		want *models.ParkingLot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
