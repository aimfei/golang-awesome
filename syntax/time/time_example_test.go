package time_example_test

import (
	time_example "golang-awesome/syntax/time"
	"reflect"
	"testing"
	"time"
)

func TestTimeAdd(t *testing.T) {
	type args struct {
		now      time.Time
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := time_example.TimeAdd(tt.args.now, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
