package stream

import (
	"reflect"
	"testing"
)

func TestStream_Filter(t *testing.T) {
	type args struct {
		predicate func(T) bool
	}
	tests := []struct {
		name   string
		stream Stream
		args   args
		want   Stream
	}{
		// ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stream.Filter(tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
