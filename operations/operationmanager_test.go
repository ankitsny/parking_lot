package operations

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestNewOperationManager(t *testing.T) {
	tests := []struct {
		name string
		want *OperationManager
	}{
		{
			name: "New Operation Manager",
			want: &OperationManager{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperationManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperationManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationManager_Parse(t *testing.T) {
	type fields struct {
		cOp      string
		argVal   string
		commands map[string]ICommand
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Operation Manager Parse[Valid]",
			fields: fields{
				argVal: "6",
				cOp:    "create_parking_lot",
			},
			args: args{
				input: "create_parking_lot 6",
			},
			wantErr: false,
		},
		{
			name: "Operation Manager Parse[inValid]",
			fields: fields{
				argVal: "6",
				cOp:    "create_parking_lot",
			},
			args: args{
				input: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opm := &OperationManager{
				cOp:      tt.fields.cOp,
				argVal:   tt.fields.argVal,
				commands: tt.fields.commands,
			}
			if err := opm.Parse(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("OperationManager.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type fakecmd struct {
	CmdName  string
	capacity int
}

// Execute :
func (fcmd *fakecmd) Execute(argsVal string) (string, error) {

	if err := fcmd.Parse(argsVal); err != nil {
		return "", err
	}
	return fmt.Sprintf("Created a parking lot with %v slots", fcmd.capacity), nil
}

// GetName :
func (fcmd *fakecmd) GetName() string {
	return fcmd.CmdName
}

// Execute :
func (fcmd *fakecmd) Parse(argsVal string) error {
	if argsVal == "" {
		return errors.New("Invalid auguments for create_parking_lot")
	}
	valStr := strings.Split(argsVal, " ")
	if val, err := strconv.Atoi(valStr[0]); err == nil {
		fcmd.capacity = val
	} else {
		return errors.New("Invalid auguments for create_parking_lot")
	}
	return nil
}

func TestOperationManager_Execute(t *testing.T) {
	fCMD := &fakecmd{
		CmdName:  "create_parking_lot",
		capacity: 6,
	}
	type fields struct {
		cOp      string
		argVal   string
		commands map[string]ICommand
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Execute Invalid",
			fields: fields{
				argVal:   "6",
				cOp:      "create_parking_lot",
				commands: map[string]ICommand{"create_parking_lot": fCMD},
			},
			args: args{
				input: "create_parking_lotttt 6", // Invalid Command
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Execute Invalid",
			fields: fields{
				argVal:   "abc",
				cOp:      "create_parking_lot abc",
				commands: map[string]ICommand{"create_parking_lot": fCMD},
			},
			args: args{
				input: "create_parking_lot abc", // Invalid Command
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Execute Invalid",
			fields: fields{
				argVal:   "",
				cOp:      "",
				commands: map[string]ICommand{"create_parking_lot": fCMD},
			},
			args: args{
				input: "", // Invalid Command
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Execute",
			fields: fields{
				argVal:   "6",
				cOp:      "create_parking_lot",
				commands: map[string]ICommand{"create_parking_lot": fCMD},
			},
			args: args{
				input: "create_parking_lot 6",
			},
			want:    fmt.Sprintf("Created a parking lot with %v slots", 6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opm := &OperationManager{
				cOp:      tt.fields.cOp,
				argVal:   tt.fields.argVal,
				commands: tt.fields.commands,
			}
			got, err := opm.Execute(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("OperationManager.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OperationManager.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationManager_AddOperation(t *testing.T) {
	fCMD := &fakecmd{
		CmdName:  "create_parking_lot",
		capacity: 6,
	}
	type fields struct {
		cOp      string
		argVal   string
		commands map[string]ICommand
	}
	type args struct {
		cmd ICommand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "AddOperation",
			fields: fields{
				argVal:   "6",
				cOp:      "create_parking_lot",
				commands: map[string]ICommand{"create_parking_lot": fCMD},
			},
			args: args{
				cmd: fCMD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opm := &OperationManager{
				cOp:      tt.fields.cOp,
				argVal:   tt.fields.argVal,
				commands: tt.fields.commands,
			}
			opm.AddOperation(tt.args.cmd)
		})
	}
}
