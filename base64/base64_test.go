package common

import "testing"

func TestHexToBase64(t *testing.T) {
	type args struct {
		hexStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "successful with 0x",
			args:    args{hexStr: "0x62F10cE5b727edf787ea45776bD050308A611508"},
			want:    "YvEM5bcn7feH6kV3a9BQMIphFQg=",
			wantErr: false,
		},
		{
			name:    "successful without 0x",
			args:    args{hexStr: "62F10cE5b727edf787ea45776bD050308A611508"},
			want:    "YvEM5bcn7feH6kV3a9BQMIphFQg=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBase64(tt.args.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexToBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
