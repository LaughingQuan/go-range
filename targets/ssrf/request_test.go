package ssrf

import "testing"

func TestSSRFHostCheck(t *testing.T) {

	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: "http://192.168.199.209/",
			want: false,
		},
		{
			name: "test",
			args: "http://example.com",
			want: false,
		},
		{
			name: "test",
			args: "https://www.baidu.com/",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SSRFHostCheck(tt.args); got != tt.want {
				t.Errorf("SSRFHostCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
