package main

import "testing"

func Test_write2file(t *testing.T) {
	type args struct {
		out []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := write2file(tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("write2file() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
