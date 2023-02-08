package xmltargets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindElement(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "testSafe",
			args: args{
				path: replacerPath("//book[@category='WEB']/title"),
			},
			want: "{}",
		},
		{
			name: "testUnSafe",
			args: args{
				path: "//book[@category='WEB']/title",
			},
			want: `{"/bookstore/book/title":"Learning XML"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindElement(tt.args.path)
			require.Equal(t, tt.want, string(got))
		})

	}
}
