package go_gh_pages

import "testing"

func TestGit_getRemoteUrl(t *testing.T) {
	type fields struct {
		Cwd    string
		Cmd    string
		Err    error
		Output string
	}
	type args struct {
		remote string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:"1111",
			fields:fields{
			Cwd:".",
			Cmd:"git",
			},
			args:args{
				remote:"origin",
			},
			want:"https://chaunceyjiang@github.com/chaunceyjiang/LanguagePractice.git",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				Cwd:    tt.fields.Cwd,
				Cmd:    tt.fields.Cmd,
				Err:    tt.fields.Err,
				Output: tt.fields.Output,
			}
			if got := g.getRemoteUrl(tt.args.remote); got != tt.want {
				t.Errorf("Git.getRemoteUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
