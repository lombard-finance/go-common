package common

import (
	"fmt"
	"testing"
	"time"
)

func TestCache_SetCache(t *testing.T) {

	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set cache",
			args: args{
				key:   "1111",
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache(nil)

			for i := 0; i < 5; i++ {
				go func() {
					c.SetCache(tt.args.key, fmt.Sprintf("%v %d", tt.args.value, i))
					cached := c.GetCache(tt.args.key)
					t.Logf("%d %v", i, cached)
				}()
			}

			time.Sleep(time.Second * 5)

			cached := c.GetCache(tt.args.key)
			t.Logf("final out %v", cached)
		})
	}
}
