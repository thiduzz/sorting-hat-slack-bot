package helpers

import "testing"

func TestFormatListBlockResponse(t *testing.T) {
	type args struct {
		listToFormat []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Send empty", args{listToFormat: []string{}}, ""},
		{"Send with strings", args{listToFormat: []string{"test1", "test2"}}, `{"blocks":[{"type":"section","text":{"text":"• test1\n• test2","type":"mrkdwn"}}]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatListBlockResponse(tt.args.listToFormat); got != tt.want {
				t.Errorf("FormatListBlockResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
