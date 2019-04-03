package models

import (
	"reflect"
	"testing"
)

func generateParingSpot(vehicleSize string, capacity int, level string) []*ParkingSpot {
	var parkingSpots []*ParkingSpot

	for i := 0; i < capacity; i++ {
		parkingSpots = append(parkingSpots, NewEmptyParkingSpot("SM", i+1, "0"))
	}
	return parkingSpots
}

func TestCreateParkingLot(t *testing.T) {

	parkingSpots := generateParingSpot("SM", 10, "0")
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
				capacity: 10,
				lotID:    "LOT_1",
			},
			want: &ParkingLot{
				address:       "Bangalore HSR",
				capacity:      10,
				lotID:         "LOT_1",
				hasEmptySpace: true,
				nextSpot:      1,
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

func TestParkingLot_getNearestParkingSpot(t *testing.T) {
	reservedSpot := generateParingSpot("SM", 1, "0")
	(reservedSpot[0]).vehicle = CreateVehicle("RED", "KA-51EZ-1234", "SM")
	type fields struct {
		lotID         string
		address       string
		parkingSpots  []*ParkingSpot
		capacity      int
		nextSpot      int
		hasEmptySpace bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "With vacant space",
			fields: fields{
				address:       "Bangalore HSR",
				capacity:      10,
				hasEmptySpace: true,
				lotID:         "LOT_1",
				nextSpot:      1,
				parkingSpots:  generateParingSpot("SM", 10, "0"),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "With vacant space",
			fields: fields{
				address:       "Bangalore HSR",
				capacity:      1,
				hasEmptySpace: true,
				lotID:         "LOT_1",
				nextSpot:      1,
				parkingSpots:  reservedSpot,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &ParkingLot{
				lotID:         tt.fields.lotID,
				address:       tt.fields.address,
				parkingSpots:  tt.fields.parkingSpots,
				capacity:      tt.fields.capacity,
				nextSpot:      tt.fields.nextSpot,
				hasEmptySpace: tt.fields.hasEmptySpace,
			}
			got, err := pl.GetNearestParkingSpot()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.GetNearestParkingSpot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingLot.GetNearestParkingSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingLot_GetNearestParkingSpot(t *testing.T) {
	type fields struct {
		lotID         string
		address       string
		parkingSpots  []*ParkingSpot
		capacity      int
		nextSpot      int
		hasEmptySpace bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &ParkingLot{
				lotID:         tt.fields.lotID,
				address:       tt.fields.address,
				parkingSpots:  tt.fields.parkingSpots,
				capacity:      tt.fields.capacity,
				nextSpot:      tt.fields.nextSpot,
				hasEmptySpace: tt.fields.hasEmptySpace,
			}
			got, err := pl.GetNearestParkingSpot()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.GetNearestParkingSpot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingLot.GetNearestParkingSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingLot_GetParkingSpotByVehicleNo(t *testing.T) {
	parkingSpot := generateParingSpot("SM", 1, "0")
	parkingSpotWithoutVehicle := generateParingSpot("SM", 1, "0")
	vehicle := CreateVehicle("Red", "KA-51EZ-1234", "SM")
	(parkingSpot[0]).vehicle = vehicle
	type fields struct {
		lotID         string
		address       string
		parkingSpots  []*ParkingSpot
		capacity      int
		nextSpot      int
		hasEmptySpace bool
	}
	type args struct {
		vNo string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ParkingSpot
		wantErr bool
	}{
		{
			name: "A vehicle is parked",
			fields: fields{
				address:       "Bangalore",
				capacity:      1,
				hasEmptySpace: false,
				parkingSpots:  parkingSpot,
			},
			args: args{
				vNo: "KA-51EZ-1234",
			},
			want:    parkingSpot[0],
			wantErr: false,
		},
		{
			name: "A vehicle is not parked",
			fields: fields{
				address:       "Bangalore",
				capacity:      1,
				hasEmptySpace: false,
				parkingSpots:  parkingSpot,
			},
			args: args{
				vNo: "KA-51EZ-0001",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Vehicle is not initialized",
			fields: fields{
				address:       "Bangalore",
				capacity:      1,
				hasEmptySpace: false,
				parkingSpots:  parkingSpotWithoutVehicle,
			},
			args: args{
				vNo: "KA-51EZ-0001",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &ParkingLot{
				lotID:         tt.fields.lotID,
				address:       tt.fields.address,
				parkingSpots:  tt.fields.parkingSpots,
				capacity:      tt.fields.capacity,
				nextSpot:      tt.fields.nextSpot,
				hasEmptySpace: tt.fields.hasEmptySpace,
			}
			got, err := pl.GetParkingSpotByVehicleNo(tt.args.vNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.GetParkingSpotByVehicleNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkingLot.GetParkingSpotByVehicleNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingLot_GetParkingSpotByColor(t *testing.T) {
	parkingSpot := generateParingSpot("SM", 1, "0")
	vehicle := CreateVehicle("Red", "KA-51EZ-1234", "SM")
	(parkingSpot[0]).vehicle = vehicle
	type fields struct {
		lotID         string
		address       string
		parkingSpots  []*ParkingSpot
		capacity      int
		nextSpot      int
		hasEmptySpace bool
	}
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ParkingSpot
		wantErr bool
	}{
		{
			name: "A Vehicle is parked and the color is Red",
			fields: fields{
				parkingSpots: parkingSpot,
			},
			args: args{
				color: "Red",
			},
			want:    parkingSpot,
			wantErr: false,
		},
		{
			name: "A Vehicle is not parked with the given color",
			fields: fields{
				parkingSpots: parkingSpot,
			},
			args: args{
				color: "Black",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &ParkingLot{
				lotID:         tt.fields.lotID,
				address:       tt.fields.address,
				parkingSpots:  tt.fields.parkingSpots,
				capacity:      tt.fields.capacity,
				nextSpot:      tt.fields.nextSpot,
				hasEmptySpace: tt.fields.hasEmptySpace,
			}
			got, err := pl.GetParkingSpotByColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.GetParkingSpotByColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkingLot.GetParkingSpotByColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingLot_Park(t *testing.T) {
	parkingSpot := generateParingSpot("SM", 1, "0")
	// vehicle := CreateVehicle("Red", "KA-51EZ-1234", "SM")
	// (parkingSpot[0]).vehicle = vehicle
	parkingSpotAfterParked := generateParingSpot("SM", 1, "0")
	vehicle := CreateVehicle("Red", "KA-51EZ-1234", "SM")
	(parkingSpotAfterParked[0]).vehicle = vehicle
	type fields struct {
		lotID         string
		address       string
		parkingSpots  []*ParkingSpot
		capacity      int
		nextSpot      int
		hasEmptySpace bool
	}
	type args struct {
		vehicleNo string
		color     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ParkingSpot
		wantErr bool
	}{
		{
			name: "Park[valid]",
			fields: fields{
				address:      "Bangalore",
				capacity:     1,
				lotID:        "LOT_1",
				parkingSpots: parkingSpot,
			},
			args: args{
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
			},
			want:    parkingSpotAfterParked[0],
			wantErr: false,
		},
		{
			name: "Park[invalid]",
			fields: fields{
				address:      "Bangalore",
				capacity:     1,
				lotID:        "LOT_1",
				parkingSpots: parkingSpotAfterParked,
			},
			args: args{
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &ParkingLot{
				lotID:         tt.fields.lotID,
				address:       tt.fields.address,
				parkingSpots:  tt.fields.parkingSpots,
				capacity:      tt.fields.capacity,
				nextSpot:      tt.fields.nextSpot,
				hasEmptySpace: tt.fields.hasEmptySpace,
			}
			got, err := pl.Park(tt.args.vehicleNo, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.Park() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkingLot.Park() = %v, want %v", got, tt.want)
			}
		})
	}
}
