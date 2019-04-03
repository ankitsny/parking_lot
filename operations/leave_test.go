package operations

import (
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/store"
)

func TestNewLeaveCMD(t *testing.T) {
	tests := []struct {
		name string
		want *Leave
	}{
		{
			name: "New Leave CMD",
			want: &Leave{
				opName: "leave",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLeaveCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLeaveCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeave_GetName(t *testing.T) {
	type fields struct {
		opName    string
		vehicleNo string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Leave CMD Get name",
			fields: fields{
				opName: "leave",
			},
			want: "leave",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Leave{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if got := l.GetName(); got != tt.want {
				t.Errorf("Leave.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeave_Parse(t *testing.T) {
	type fields struct {
		opName    string
		vehicleNo string
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
			name: "LeaveParse [invalid]",
			fields: fields{
				opName: "leave",
			},
			args: args{
				argVal: "",
			},
			wantErr: true,
		},
		{
			name: "LeaveParse [valid]",
			fields: fields{
				opName: "leave",
			},
			args: args{
				argVal: "KA-51EZ-1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Leave{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if err := l.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("Leave.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeave_Execute(t *testing.T) {
	store.NewStore(1)
	store.GetStorage().Park("KA-51EZ-1234", "Red")
	type fields struct {
		opName    string
		vehicleNo string
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
			name: "Leave CMD Execute[invalid]",
			fields: fields{
				opName: "leave",
			},
			args: args{
				argVal: "KA-51EZ-1000",
			},
			want: "Not Found",
		},
		{
			name: "Leave CMD Execute without args[invalid]",
			fields: fields{
				opName: "leave",
			},
			args: args{
				argVal: "",
			},
			want: "Invalid args for leave command",
		},
		{
			name: "Leave CMD Execute[valid]",
			fields: fields{
				opName: "leave",
			},
			args: args{
				argVal: "KA-51EZ-1234",
			},
			want: "Slot number 1 is free",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Leave{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if got := l.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("Leave.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
