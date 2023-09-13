package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{10, 10, 20},
		{4, 2, 6},
		{1, 1, 2},
		{0, 12, 12},
		{-3, 2, -1},
		{-2, -6, -8},
	}

	for _, test := range tests {

		testName := fmt.Sprintf("adding %d + %d", test.a, test.b)
		fmt.Println(testName)
		t.Run(testName, func(t *testing.T) {
			got := Add(test.a, test.b)

			assertCorrectInt(t, got, test.want)
		})
	}
}

func assertCorrectInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d got %d", want, got)
	}
}
