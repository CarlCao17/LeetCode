package slices

import "testing"

func TestSum(t *testing.T) {
	type args[T Number] struct {
		s []T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want T
	}

	tt := testCase[int32]{
		name: "int32 Sum",
		args: args[int32]{[]int32{1, 2, 3, -1, -2, -3}},
		want: 0,
	}
	if got := Sum(tt.args.s); got != tt.want {
		t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
	}

	tt1 := testCase[uint16]{
		name: "uint16 Sum",
		args: args[uint16]{[]uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		want: 45,
	}
	if got := Sum(tt1.args.s); got != tt1.want {
		t.Errorf("%s() = %v, want %v", tt1.name, got, tt1.want)
	}

	tt2 := testCase[float64]{
		name: "float64 Sum",
		args: args[float64]{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		want: 45,
	}
	if got := Sum(tt2.args.s); got != tt2.want {
		t.Errorf("%s() = %v, want %v", tt2.name, got, tt2.want)
	}

	tt3 := testCase[complex128]{
		name: "complex128 Sum",
		args: args[complex128]{[]complex128{0 + 1i, 1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i, 7 + 8i, 8 + 9i, 9 + 10i}},
		want: 45 + 55i,
	}
	if got := Sum(tt3.args.s); got != tt3.want {
		t.Errorf("%s() = %v, want %v", tt3.name, got, tt3.want)
	}
}
