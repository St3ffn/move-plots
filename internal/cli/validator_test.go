package cli

import (
	"errors"
	"move-plots/internal/test"
	"testing"
)

func TestIsDirectory_Enforce(t *testing.T) {
	tests := []struct {
		name       string
		filesystem test.DummyFilesystem
		path       string
		wantErr    bool
	}{
		{
			name: "ok",
			filesystem: test.DummyFilesystem{
				Directory: true,
			},
			path:    "/hallo",
			wantErr: false,
		},
		{
			name: "no directory",
			filesystem: test.DummyFilesystem{
				Directory: false,
			},
			path:    "/hallo",
			wantErr: true,
		},
		{
			name: "stat error",
			filesystem: test.DummyFilesystem{
				StatErr: errors.New("boom"),
			},
			path:    "/hallo",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := IsDirectory{tt.filesystem}
			if err := d.Enforce(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("Enforce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
