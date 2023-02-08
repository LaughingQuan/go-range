package nosql

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"xmirror.cn/iast/goat/config"
)

func TestGetConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.GetConfMog()
			require.NoError(t, err)
			log.Printf("%v", got)
		})
	}
}
