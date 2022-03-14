package sha

import (
	"reflect"
	"testing"
)

func TestSha1Encrypt(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1Encrypt(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sha1Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
