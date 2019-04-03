package operations

import (
	"reflect"
	"testing"

	"github.com/anks333/parking_lot/store"
)

func TestNewRegistrationNoBasedOnColorCMD(t *testing.T) {
	tests := []struct {
		name string
		want *RegistrationNoBasedOnColor
	}{
		{
			name: "registration_numbers_for_cars_with_colour CMD",
			want: &RegistrationNoBasedOnColor{
				opName: "registration_numbers_for_cars_with_colour",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistrationNoBasedOnColorCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistrationNoBasedOnColorCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_GetName(t *testing.T) {
	type fields struct {
		opName string
		color  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Name",
			fields: fields{
				color:  "Red",
				opName: "registration_numbers_for_cars_with_colour",
			},
			want: "registration_numbers_for_cars_with_colour",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if got := rbc.GetName(); got != tt.want {
				t.Errorf("RegistrationNoBasedOnColor.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_Parse(t *testing.T) {
	type fields struct {
		opName string
		color  string
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
				opName: "registration_numbers_for_cars_with_colour",
			},
			args: args{
				argVal: "Red",
			},
			wantErr: false,
		},
		{
			name: "Parse[invalid]",
			fields: fields{
				opName: "registration_numbers_for_cars_with_colour",
			},
			args: args{
				argVal: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if err := rbc.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("RegistrationNoBasedOnColor.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_Execute(t *testing.T) {
	store.NewStore(2)
	store.GetStorage().Park("KA-51EZ-1234", "Red")
	store.GetStorage().Park("KA-51EZ-1235", "Red")
	type fields struct {
		opName string
		color  string
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
				opName: "registration_numbers_for_cars_with_colour",
			},
			args: args{argVal: "Red"},
			want: "KA-51EZ-1234, KA-51EZ-1235",
		},
		{
			name: "2",
			fields: fields{
				opName: "registration_numbers_for_cars_with_colour",
			},
			args: args{argVal: "White"},
			want: "Not Found",
		},
		{
			name: "3",
			fields: fields{
				opName: "registration_numbers_for_cars_with_colour",
			},
			args: args{argVal: ""},
			want: "Invalid args for command",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if got := rbc.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("RegistrationNoBasedOnColor.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
