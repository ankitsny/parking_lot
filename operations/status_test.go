package operations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/models"
	"github.com/anks333/parking_lot/store"
)

func TestNewStatusCMD(t *testing.T) {
	tests := []struct {
		name string
		want *Status
	}{
		{
			name: "1",
			want: &Status{
				opName: "status",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_GetName(t *testing.T) {
	type fields struct {
		opName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				opName: "status",
			},
			want: "status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Status{
				opName: tt.fields.opName,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("Status.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Parse(t *testing.T) {
	type fields struct {
		opName string
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				opName: "status",
			},
			args: args{
				argVal: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Status{
				opName: tt.fields.opName,
			}
			if err := s.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("Status.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatus_Execute(t *testing.T) {
	store.SetStorage(nil)
	store.NewStore(1)
	store.GetStorage().Park("KA-1-1212", "Red")
	type fields struct {
		opName  string
		storage *models.ParkingLot
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// {
		// 	name: "1",
		// 	fields: fields{
		// 		opName: "status",
		// 	},
		// 	args: args{
		// 		argVal: "",
		// 	},
		// 	want: "Parking lot is not created",
		// },
		{
			name: "2",
			fields: fields{
				opName: "status",
			},
			args: args{
				argVal: "",
			},
			want: fmt.Sprintf(
				"%v\t%v\t%v\n%v\t%v\t%v",
				"Slot No.",
				"Registration No",
				"Colour",
				1,
				"KA-1-1212",
				"Red",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Status{
				opName: tt.fields.opName,
			}
			if got := s.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("Status.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
