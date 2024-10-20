package ruby_test

import (
	"fmt"
	"testing"

	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
	"github.com/stretchr/testify/assert"
)

func TestBool2Int(t *testing.T) {
	tests := []struct {
		arg  bool
		want int
	}{
		{
			arg:  true,
			want: 1,
		},
		{
			arg:  false,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arg), func(t *testing.T) {
			got := ruby.Bool2Int(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
