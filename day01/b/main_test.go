package main

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

var testData = []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
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
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ReadInput(tt.args.filename); !reflect.DeepEqual(got, tt.want) || err != tt.wantErr {
				t.Errorf("ReadInput(%s) = %v,%v ; want %v, %v", tt.args.filename, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "good",
			args: args{"12345"},
			want: "54321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.input); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
