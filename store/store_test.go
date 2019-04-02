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
		{
			name: "new store[invalid]",
			args: args{
				capacity: 0,
			},
			wantErr: true,
		},
		{
			name: "new store[valid]",
			args: args{
				capacity: 1,
			},
			wantErr: false,
		},
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
	NewStore(1)
	pLot, _ := models.CreateParkingLot("LOT_1", "Bangalore", 1)
	tests := []struct {
		name string
		want *models.ParkingLot
	}{
		{
			name: "get storage",
			want: pLot,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
