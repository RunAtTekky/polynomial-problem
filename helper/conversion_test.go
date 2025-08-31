package helper

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	cases := []struct {
		base  string
		value string
		want  int
	}{
		{base: "2", value: "111", want: 7},
		{base: "10", value: "12", want: 12},
		{base: "4", value: "213", want: 39},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("base %q, value %q", tt.base, tt.value), func(t *testing.T) {
			got := BaseConversion(tt.base, tt.value)
			want := tt.want

			if got != want {
				t.Fatalf("got %d but want %d", got, want)
			}

		})
	}

}
