package operations

import (
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/store"
)

func TestNewSlotNoBasedOnRegNoCMD(t *testing.T) {
	tests := []struct {
		name string
		want *SlotNoBasedOnRegNo
	}{
		{
			name: "slot_number_for_registration_number CMD",
			want: &SlotNoBasedOnRegNo{
				opName: "slot_number_for_registration_number",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlotNoBasedOnRegNoCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlotNoBasedOnRegNoCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlotNoBasedOnRegNo_GetName(t *testing.T) {
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
			name: "Get Name",
			fields: fields{
				vehicleNo: "Red",
				opName:    "slot_number_for_registration_number",
			},
			want: "slot_number_for_registration_number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &SlotNoBasedOnRegNo{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if got := rbc.GetName(); got != tt.want {
				t.Errorf("SlotNoBasedOnRegNo.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlotNoBasedOnRegNo_Parse(t *testing.T) {
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
			name: "Parse",
			fields: fields{
				opName: "slot_number_for_registration_number",
			},
			args: args{
				argVal: "KA-51EZ-1234",
			},
			wantErr: false,
		},
		{
			name: "Parse[invalid]",
			fields: fields{
				opName: "slot_number_for_registration_number",
			},
			args: args{
				argVal: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &SlotNoBasedOnRegNo{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if err := rbc.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("SlotNoBasedOnRegNo.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlotNoBasedOnRegNo_Execute(t *testing.T) {
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
			name: "1",
			fields: fields{
				opName: "slot_number_for_registration_number",
			},
			args: args{argVal: "KA-51EZ-1234"},
			want: "1",
		},
		{
			name: "2",
			fields: fields{
				opName: "slot_number_for_registration_number",
			},
			args: args{argVal: "KA-AA-1232"},
			want: "Not Found",
		},
		{
			name: "3",
			fields: fields{
				opName: "slot_number_for_registration_number",
			},
			args: args{argVal: ""},
			want: "Invalid args",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &SlotNoBasedOnRegNo{
				opName:    tt.fields.opName,
				vehicleNo: tt.fields.vehicleNo,
			}
			if got := rbc.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("SlotNoBasedOnRegNo.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
