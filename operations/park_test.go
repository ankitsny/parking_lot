package operations

import (
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/store"
)

func TestNewParkCMD(t *testing.T) {
	tests := []struct {
		name string
		want *Park
	}{
		{
			name: "NewParkCMD",
			want: &Park{opName: "park"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParkCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParkCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPark_GetName(t *testing.T) {
	type fields struct {
		opName    string
		vehicleNo string
		color     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetName",
			fields: fields{
				opName: "park",
			},
			want: "park",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Park{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
				color:     tt.fields.color,
			}
			if got := p.GetName(); got != tt.want {
				t.Errorf("Park.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPark_Parse(t *testing.T) {
	type fields struct {
		opName    string
		vehicleNo string
		color     string
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
			name: "Park Parse[valid]",
			fields: fields{
				color:     "Red",
				opName:    "park",
				vehicleNo: "KA-51EZ-1234",
			},
			args: args{
				argVal: "KA-51EZ-1234 Red",
			},
			wantErr: false,
		},
		{
			name: "Park Parse[invalid]",
			fields: fields{
				opName:    "park",
				vehicleNo: "KA-51EZ-1234",
			},
			args: args{
				argVal: "KA-51EZ-1234",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Park{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
				color:     tt.fields.color,
			}
			if err := p.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("Park.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPark_Execute(t *testing.T) {
	store.NewStore(1)
	type fields struct {
		opName    string
		vehicleNo string
		color     string
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
		{
			name: "Park execute [invalid args]",
			fields: fields{
				opName:    "park",
				vehicleNo: "KA-51EZ-1234",
			},
			args: args{
				argVal: "KA-51EZ-1234",
			},
			want: "Invalid args for park command",
		},
		{
			name: "Park execute [valid]",
			fields: fields{
				opName:    "park",
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
			},
			args: args{
				argVal: "KA-51EZ-1234 Red",
			},
			want: "Allocated slot number: 1",
		},
		{
			name: "Park execute [Parking lot full]",
			fields: fields{
				opName:    "park",
				color:     "Red",
				vehicleNo: "KA-51EZ-1234",
			},
			args: args{
				argVal: "KA-51EZ-1234 Red",
			},
			want: "Sorry, parking lot is full",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Park{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
				color:     tt.fields.color,
			}
			if got := p.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("Park.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
