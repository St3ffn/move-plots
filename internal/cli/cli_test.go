package cli

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestRunCli(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Context
		wantErr error
	}{
		{
			name: "ok",
			args: []string{"move-plots", "/source", "/target1", "/target2"},
			want: Context{
				Reserved: Reserved,
				Source:   "/source",
				Targets:  []string{"/target1", "/target2"},
			},
		},
		{
			name: "help short",
			args: []string{"move-plots", "-h"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "help long",
			args: []string{"move-plots", "--help"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "reserve none short",
			args: []string{"move-plots", "-r", "0", "/source", "/target1", "/target2"},
			want: Context{
				Reserved: 0,
				Source:   "/source",
				Targets:  []string{"/target1", "/target2"},
			},
		},
		{
			name: "reserve 11231230 long",
			args: []string{"move-plots", "--reserve", "11231230", "/source", "/target1"},
			want: Context{
				Reserved: 11231230,
				Source:   "/source",
				Targets:  []string{"/target1"},
			},
		},
		{
			name:    "err no source and target",
			args:    []string{"move-plots"},
			wantErr: errors.New("SOURCE_DIRECTORY and TARGET_DIRECTORY missing"),
		},
		{
			name:    "err no target",
			args:    []string{"move-plots", "/tmp"},
			wantErr: errors.New("TARGET_DIRECTORY missing"),
		},
		{
			name:    "unknown parameter -x",
			args:    []string{"move-plots", "-x", "asdas"},
			wantErr: errors.New("flag provided but not defined: -x"),
		},
		{
			name:    "invalid reserve paramter",
			args:    []string{"move-plots", "-r", "12.12", "/home/steffen"},
			wantErr: errors.New("invalid value \"12.12\" for flag -r: parse error"),
		},
		{
			name: "show version short",
			args: []string{"move-plots", "-V"},
			want: Context{
				Done: true,
			},
		},
		{
			name: "show version long",
			args: []string{"move-plots", "--version"},
			want: Context{
				Done: true,
			},
		},
		{
			name: "verbose mode short",
			args: []string{"move-plots", "-v", "/tmp", "/target1", "/target2"},
			want: Context{
				Source:  "/tmp",
				Targets: []string{"/target1", "/target2"},
				Verbose: true,
			},
		},
		{
			name: "verbose mode long",
			args: []string{"move-plots", "--verbose", "/tmp", "/target1", "/target2"},
			want: Context{
				Source:  "/tmp",
				Targets: []string{"/target1", "/target2"},
				Verbose: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Args = tt.args
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
