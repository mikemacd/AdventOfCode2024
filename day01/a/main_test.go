package main

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

var testData = []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)
var testFile, _ = os.CreateTemp(os.TempDir(), "testdata")

func init() {
	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)
}

func LogAndDie(err error) {
	slog.Error("Fatal error:", "error", err)
	os.Exit(1)
}

func getTestDatafile() string {
	slog.Debug("testFile: %s\n", "filename", testFile.Name())

	if _, err := testFile.Write(testData); err != nil {
		LogAndDie(err)
	}

	if err := testFile.Close(); err != nil {
		LogAndDie(err)
	}

	return testFile.Name()
}

func Test_ReadInput(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Datarows
		wantErr error
	}{
		{
			name: "Valid Datafile",
			args: args{
				filename: getTestDatafile(),
			},
			want: Datarows{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInput(tt.args.filename)
			if !reflect.DeepEqual(got, tt.want) || err != tt.wantErr {
				t.Errorf("ReadInput(%s) = %v,%v ; want %v, %v", tt.args.filename, got, err, tt.want, tt.wantErr)
			}
			rv, err := ProcessData(got)
			_ = err
			_ = rv
		})
	}
}
