package cli

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

type testValidation struct {
	error error
}

func (v testValidation) Enforce(_ string) error {
	if v.error != nil {
		return v.error
	}
	return nil
}

func TestRunCli(t *testing.T) {
	tests := []struct {
		name      string
		validator Validator
		args      []string
		want      Context
		wantErr   error
	}{
		{
			name: "ok",
			validator: testValidation{},
			args: []string{"move-plots", "/source", "/target1", "/target2"},
			want: Context{
				Reserved: Reserved,
				Source:   "/source",
				Targets:  []string{"/target1", "/target2"},
			},
		},
		{
			name: "help short",
			validator: testValidation{},
			args: []string{"move-plots", "-h"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "help long",
			validator: testValidation{},
			args: []string{"move-plots", "--help"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "reserve none short",
			validator: testValidation{},
			args: []string{"move-plots", "-r", "0", "/source", "/target1", "/target2"},
			want: Context{
				Reserved: 0,
				Source:   "/source",
				Targets:  []string{"/target1", "/target2"},
			},
		},
		{
			name: "reserve 11231230 long",
			validator: testValidation{},
			args: []string{"move-plots", "--reserve", "11231230", "/source", "/target1"},
			want: Context{
				Reserved: 11231230,
				Source:   "/source",
				Targets:  []string{"/target1"},
			},
		},
		{
			name:    "err no source and target",
			validator: testValidation{},
			args:    []string{"move-plots"},
			wantErr: errors.New("SOURCE_DIRECTORY and TARGET_DIRECTORY missing"),
		},
		{
			name:    "err no target",
			validator: testValidation{},
			args:    []string{"move-plots", "/tmp"},
			wantErr: errors.New("TARGET_DIRECTORY missing"),
		},
		{
			name:    "unknown parameter -x",
			validator: testValidation{},
			args:    []string{"move-plots", "-x", "asdas"},
			wantErr: errors.New("flag provided but not defined: -x"),
		},
		{
			name:    "invalid reserve paramter",
			validator: testValidation{},
			args:    []string{"move-plots", "-r", "12.12", "/home/steffen"},
			wantErr: errors.New("invalid value \"12.12\" for flag -r: parse error"),
		},
		{
			name: "show version short",
			validator: testValidation{},
			args: []string{"move-plots", "-V"},
			want: Context{
				Done: true,
			},
		},
		{
			name: "show version long",
			validator: testValidation{},
			args: []string{"move-plots", "--version"},
			want: Context{
				Done: true,
			},
		},
		{
			name: "verbose mode short",
			validator: testValidation{},
			args: []string{"move-plots", "-v", "/tmp", "/target1", "/target2"},
			want: Context{
				Source:  "/tmp",
				Targets: []string{"/target1", "/target2"},
				Verbose: true,
			},
		},
		{
			name: "verbose mode long",
			validator: testValidation{},
			args: []string{"move-plots", "--verbose", "/tmp", "/target1", "/target2"},
			want: Context{
				Source:  "/tmp",
				Targets: []string{"/target1", "/target2"},
				Verbose: true,
			},
		},
		{
			name: "validation fails mode long",
			validator: testValidation{ errors.New("path broken") },
			args: []string{"move-plots", "--verbose", "/tmp", "/target1", "/target2"},
			wantErr: errors.New("path broken"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Args = tt.args
			Validation = tt.validator
			got, err := RunCli(os.Stdout, "idk")
			if err != nil {
				if tt.wantErr == nil || !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("RunCli() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("RunCli() got = %v, want %v", got, tt.want)
			}
		})
	}
}
