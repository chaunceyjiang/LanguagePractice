package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLoadCsv(t *testing.T) {
	type args struct {
		filename string
	}
	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	fmt.Println(rst)
	tests := []struct {
		name string
		args args
	}{
		{name: "trans_type",
		args: args{"/home/chauncey/Goexample/bms/protocol/trans_type.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadCsv(tt.args.filename)
		})
	}
}
