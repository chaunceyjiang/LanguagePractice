package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_AssicHandler(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "assicHandler",
			args: args{index: 1},
			want: "A",
		},
		//{name: "assicHandler",
		//	args: args{index: 2},
		//	want: "b",
		//},
		//{name: "assicHandler",
		//	args: args{index: 26},
		//	want: "z",
		//},
		//{name: "assicHandler",
		//	args: args{index: 27},
		//	want: "aa",
		//},
		{name: "assicHandler",
			args: args{index: 52},
			want: "AZ",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsiicHandler(tt.args.index); got != tt.want {
				t.Errorf("AssicHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
