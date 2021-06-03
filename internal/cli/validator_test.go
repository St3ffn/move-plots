package cli

import (
	"errors"
	"io/fs"
	"testing"
	"time"
)

type testStat struct {
	error error
	isDir bool
}

func (t testStat) stat(_ string) (fs.FileInfo, error) {
	if t.error != nil {
		return nil, t.error
	}
	return testDirectory{
		error: t.error,
		isDir: t.isDir,
	}, nil
}

type testDirectory struct {
	error error
	isDir bool
}

func (t testDirectory) Name() string {
	panic("implement me")
}

func (t testDirectory) Size() int64 {
	panic("implement me")
}

func (t testDirectory) Mode() fs.FileMode {
	panic("implement me")
}

func (t testDirectory) ModTime() time.Time {
	panic("implement me")
}

func (t testDirectory) Sys() interface{} {
	panic("implement me")
}

func (t testDirectory) IsDir() bool {
	return t.isDir
}

func TestIsDirectory_Enforce(t *testing.T) {
	tests := []struct {
		name    string
		stat    testStat
		path    string
		wantErr bool
	}{
		{
			name: "ok",
			stat: testStat{
				isDir: true,
			},
			path:    "/hallo",
			wantErr: false,
		},
		{
			name: "no directory",
			stat: testStat{
				isDir: false,
			},
			path:    "/hallo",
			wantErr: true,
		},
		{
			name: "stat error",
			stat: testStat{
				error: errors.New("boom"),
			},
			path:    "/hallo",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := IsDirectory{
				stat: tt.stat.stat,
			}
			if err := d.Enforce(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("Enforce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
