package plot

import (
	"errors"
	"github.com/St3ffn/plots-left/pkg/disk"
	"move-plots/internal/test"
	"reflect"
	"testing"
)

func TestFindPlots(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		filesystem test.DummyFilesystem
		wantPlots  []string
		wantErr    error
	}{
		{
			name: "no plots",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				Files: []string{
					"plot-k32-2021-05-22-06-26-f30dcf41079d08a3a0dc39721a01631ba20723cf921ee418af88be76.plot.2.tmp",
					"asdas",
					"a",
					"23",
				},
			},
		},
		{
			name: "one plot",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				Files: []string{
					"plot-k32-2021-05-22-06-26-f41079d08a3a0dc33c5d9a939721a01631ba20723cf921ee418af88be76.plot.2.tmp",
					"asdas",
					"a",
					"plot-k32-2021-06-03-03-10-6fb10d66b811b13d39b2d45a6c3a1fb870b3ad19c7b3.plot",
					"23",
				},
			},
			wantPlots: []string{"plot-k32-2021-06-03-03-10-6fb10d66b811b13d39b2d45a6c3a1fb870b3ad19c7b3.plot"},
		},
		{
			name: "two plots",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				Files: []string{
					"plot-k32-2021-05-22-06-26-f41079d08a3a0dc33c5d9a939721a01631ba20723cf921ee418af88be76.plot.2.tmp",
					"asdas",
					"a",
					"plot-k32-2021-06-03-03-10-6fb10d66b811b13d39b2d45a6c3a1fb870b3ad19c7b3.plot",
					"23",
					"plot-k32-2021-06-03-04-53-b53be61ee4733bc36a9305d80dc7ac1f11b4f5c5ebd85.plot",
				},
			},
			wantPlots: []string{
				"plot-k32-2021-06-03-03-10-6fb10d66b811b13d39b2d45a6c3a1fb870b3ad19c7b3.plot",
				"plot-k32-2021-06-03-04-53-b53be61ee4733bc36a9305d80dc7ac1f11b4f5c5ebd85.plot",
			},
		},
		{
			name: "open error",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				OpenErr: errors.New("open error"),
			},
			wantErr: errors.New("error opening /some/path: open error"),
		},
		{
			name: "close error",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				CloseErr: errors.New("close error"),
			},
			wantErr: errors.New("error closing /some/path: close error"),
		},
		{
			name: "readdir error",
			path: "/some/path",
			filesystem: test.DummyFilesystem{
				ReadDirErr: errors.New("readdir error"),
			},
			wantErr: errors.New("error reading /some/path: readdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fs = tt.filesystem
			gotPlots, err := FindPlots(tt.path)
			if err != nil {
				if tt.wantErr == nil || !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("FindPlots() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotPlots, tt.wantPlots) {
				t.Errorf("FindPlots() gotPlots = %v, want %v", gotPlots, tt.wantPlots)
			}
		})
	}
}

func TestFindDisk(t *testing.T) {
	tests := []struct {
		name        string
		reserved    uint64
		paths       []string
		dummyStatfs test.DummyStatfs
		want        string
		wantErr     error
	}{
		{
			name:        "first disk",
			reserved:    1,
			paths:       []string{"/one", "/two", "/three"},
			dummyStatfs: test.DummyStatfs{PathPlotsLeft: map[string]uint64{"/one": 3, "/two": 2, "/three": 1}},
			want:        "/one",
		},
		{
			name:        "second disk",
			reserved:    1,
			paths:       []string{"/one", "/two", "/three"},
			dummyStatfs: test.DummyStatfs{PathPlotsLeft: map[string]uint64{"/one": 1, "/two": 2, "/three": 1}},
			want:        "/two",
		},
		{
			name:        "none",
			reserved:    1,
			paths:       []string{"/one", "/two", "/three"},
			dummyStatfs: test.DummyStatfs{PathPlotsLeft: map[string]uint64{"/one": 1, "/two": 0, "/three": 1}},
		},
		{
			name:        "none",
			reserved:    1,
			paths:       []string{"/one", "/two", "/three"},
			dummyStatfs: test.DummyStatfs{PathPlotsLeft: map[string]uint64{"/one": 1}, Err: errors.New("some error")},
			wantErr:     errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk.Statfs = tt.dummyStatfs.Statfs
			gotInfo, err := FindDisk(tt.reserved, tt.paths)

			if err != nil {
				if tt.wantErr == nil || !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("FindDisk() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if gotInfo == nil && tt.want != "" {
				t.Errorf("FindDisk() gotInfo = %v, want %v", gotInfo, tt.want)
				return
			}

			if gotInfo != nil && gotInfo.Path != tt.want {
				t.Errorf("FindDisk() gotInfo = %v, want %v", gotInfo.Path, tt.want)
			}
		})
	}
}
