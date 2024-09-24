package common

import (
	"math"
	"reflect"
	"testing"
)

func TestBackupUUID(t *testing.T) {
	type args struct {
		val uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "max uint64",
			args: args{
				val: math.MaxUint64,
			},
			want: "ffffffff-ffff-ffff...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackupUUID(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BackupUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
