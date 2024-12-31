package main

import (
	"reflect"
	"testing"
)

func Test_gap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				n: 1041,
			},
			want: 5,
		},
		{
			name: "2nd",
			args: args{
				n: 32,
			},
			want: 0,
		},
		{
			name: "3rd",
			args: args{
				n: 561892,
			},
			want: 3,
		},
		{
			name: "4th",
			args: args{
				n: 74901729,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gap(tt.args.n); got != tt.want {
				t.Errorf("gap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		a []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 3,
			},
			want: []int{9, 7, 6, 3, 8},
		},
		{
			name: "2nd",
			args: args{
				a: []int{0, 0, 0},
				k: 1,
			},
			want: []int{0, 0, 0},
		},
		{
			name: "3rd",
			args: args{
				a: []int{1, 2, 3, 4},
				k: 4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "empty array",
			args: args{
				a: []int{},
				k: 1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unpaired(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				[]int{9, 3, 9, 3, 9, 7, 9},
			},
			want: 7,
		},
		{
			name: "empty slice",
			args: args{
				[]int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unpaired(tt.args.a); got != tt.want {
				t.Errorf("unpaired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jump(t *testing.T) {
	type args struct {
		x int
		y int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				x: 10,
				y: 85,
				d: 30,
			},
			want: 3,
		},
		{
			name: "2nd",
			args: args{
				x: 10,
				y: 60,
				d: 5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.x, tt.args.y, tt.args.d); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missing(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				a: []int{2, 3, 1, 5},
			},
			want: 4,
		},
		{
			name: "corner case: N = 0",
			args: args{a: []int{}},
			want: 1,
		},
		{
			name: "corner case: N = 1",
			args: args{a: []int{1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missing(tt.args.a); got != tt.want {
				t.Errorf("missing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mixAbs(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				a: []int{3, 1, 2, 4, 3},
			},
			want: 1,
		},
		{
			name: "2nd",
			args: args{
				a: []int{2, 5, 7},
			},
			want: 0,
		},
		{
			name: "N=2",
			args: args{
				a: []int{1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mixAbs(tt.args.a); got != tt.want {
				t.Errorf("mixAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jumpLeaf(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				x: 5,
				a: []int{1, 3, 1, 4, 2, 3, 5, 4},
			},
			want: 6,
		},
		{
			name: "Corner case: N=1",
			args: args{
				x: 1,
				a: []int{1},
			},
			want: 0,
		},
		{
			name: "N=2",
			args: args{
				x: 2,
				a: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "N=3",
			args: args{
				x: 2,
				a: []int{1, 3, 2},
			},
			want: 2,
		},
		{
			name: "N=4",
			args: args{
				x: 5,
				a: []int{1, 5, 2, 3},
			},
			want: -1,
		},
		{
			name: "N=4",
			args: args{
				x: 5,
				a: []int{1, 5, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpLeaf(tt.args.x, tt.args.a); got != tt.want {
				t.Errorf("jumpLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPermutation(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example #1",
			args: args{
				a: []int{4, 3, 1, 2},
			},
			want: true,
		},
		{
			name: "Example #2",
			args: args{
				a: []int{4, 1, 3},
			},
			want: false,
		},
		{
			name: "Corner case: N=1 #1",
			args: args{
				a: []int{1},
			},
			want: true,
		},
		{
			name: "Corner case: N=1 #2",
			args: args{
				a: []int{2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPermutation(tt.args.a); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumProduct(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				a: []int{-3, 1, 2, 0, -2, 5, 6},
			},
			want: 60,
		},
		{
			name: "Edge case",
			args: args{
				a: []int{-10, -5, 1, 0, 2, 5, 6},
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumProduct(tt.args.a); got != tt.want {
				t.Errorf("minimumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parkingLotFee(t *testing.T) {
	type args struct {
		e string
		l string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				e: "10:00",
				l: "11:42",
			},
			want: 9,
		},
		{
			name: "Case #1",
			args: args{
				e: "10:00",
				l: "10:42",
			},
			want: 5,
		},
		{
			name: "Edge Case #2",
			args: args{
				e: "10:00",
				l: "10:00",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parkingLotFee(tt.args.e, tt.args.l); got != tt.want {
				t.Errorf("parkingLotFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_symmetryPoint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				s: "racecar",
			},
			want: 3,
		},
		{
			name: "Asymmetry",
			args: args{
				s: "hogeg",
			},
			want: -1,
		},
		{
			name: "Even string length",
			args: args{
				s: "hoge",
			},
			want: -1,
		},
		{
			name: "Edge case: empty string",
			args: args{
				s: "",
			},
			want: -1,
		},
		{
			name: "Edge case: only 1 charactor",
			args: args{
				s: "x",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := symmetryPoint(tt.args.s); got != tt.want {
				t.Errorf("symmetryPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isProperlyNested(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[]",
			args: args{"[]"},
			want: true,
		},
		{
			name: "()",
			args: args{"()"},
			want: true,
		},
		{
			name: "{}",
			args: args{"{}"},
			want: true,
		},
		{
			name: "{()}",
			args: args{"{}"},
			want: true,
		},
		{
			name: "{[()()]}",
			args: args{"{[()()]}"},
			want: true,
		},
		{
			name: "()[]",
			args: args{"()[]"},
			want: true,
		},
		{
			name: "([)()]",
			args: args{"([)()]"},
			want: false,
		},
		{
			name: "((",
			args: args{"(("},
			want: false,
		},
		{
			name: "Edge case: empty string",
			args: args{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isProperlyNested(tt.args.s); got != tt.want {
				t.Errorf("isProperlyNested() = %v, want %v", got, tt.want)
			}
		})
	}
}
