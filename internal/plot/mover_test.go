package plot

import (
	"errors"
	"github.com/St3ffn/plots-left/pkg/disk"
	"move-plots/internal/test"
	"reflect"
	"testing"
)

func TestMovePlot(t *testing.T) {
	type args struct {
		sourcePath     string
		sourceFilename string
		targetDisk     *disk.PlotInfo
	}
	tests := []struct {
		name       string
		args       args
		dummyMover test.DummyMover
		wantErr    error
	}{
		{
			name: "ok",
			args: args{
				sourcePath:     "/plots/ssd",
				sourceFilename: "myfancyplot.plot",
				targetDisk: &disk.PlotInfo{
					Disk: &disk.Disk{
						Path: "/plots/target",
					},
				},
			},
		},
		{
			name: "error moving plot-mv",
			args: args{
				sourcePath:     "/plots/ssd",
				sourceFilename: "myfancyplot.plot",
				targetDisk: &disk.PlotInfo{
					Disk: &disk.Disk{
						Path: "/plots/target",
					},
				},
			},
			dummyMover: test.DummyMover{
				TargetFileError: map[string]error{
					"/plots/target/myfancyplot.plot-mv": errors.New("boom"),
				},
			},
			wantErr: errors.New("can not move /plots/ssd/myfancyplot.plot to /plots/target/myfancyplot.plot-mv: boom"),
		},
		{
			name: "error moving plot",
			args: args{
				sourcePath:     "/plots/ssd",
				sourceFilename: "myfancyplot.plot",
				targetDisk: &disk.PlotInfo{
					Disk: &disk.Disk{
						Path: "/plots/target",
					},
				},
			},
			dummyMover: test.DummyMover{
				TargetFileError: map[string]error{
					"/plots/target/myfancyplot.plot": errors.New("boom"),
				},
			},
			wantErr: errors.New("can not move /plots/target/myfancyplot.plot-mv to /plots/target/myfancyplot.plot: boom"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mover = tt.dummyMover
			err := MovePlot(tt.args.sourcePath, tt.args.sourceFilename, tt.args.targetDisk)

			if (tt.wantErr != nil && err == nil) || (tt.wantErr == nil && err != nil) {
				t.Errorf("FindDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("FindDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbsoluteFilename(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		filename string
		want     string
	}{
		{
			name:     "ok",
			path:     "/my/path",
			filename: "somefilename.txt",
			want:     "/my/path/somefilename.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsoluteFilename(tt.path, tt.filename); got != tt.want {
				t.Errorf("AbsoluteFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
