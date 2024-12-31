package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

}

// A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

// For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

func gap(n int) int {
	bin := strconv.FormatInt(int64(n), 2)

	var gap, temp int
	for _, r := range bin {
		if r == '0' {
			temp++
		} else {
			if temp > gap {
				gap = temp
			}
			temp = 0
		}
	}
	return gap
}

// An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

// The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.
func rotate(a []int, k int) []int {
	if len(a) == 0 {
		return []int{}
	}
	len := len(a)
	for i := 0; i < k; i++ {
		mid := a[0 : len-1]
		tail := a[len-1]
		mid = append([]int{tail}, mid...)
		a = mid
	}
	return a
}

// A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

// For example, in array A such that:

//   A[0] = 9  A[1] = 3  A[2] = 9
//   A[3] = 3  A[4] = 9  A[5] = 7
//   A[6] = 9
// the elements at indexes 0 and 2 have value 9,
// the elements at indexes 1 and 3 have value 3,
// the elements at indexes 4 and 6 have value 9,
// the element at index 5 has value 7 and is unpaired.

func unpaired(a []int) int {
	m := make(map[int]int, len(a))
	for _, n := range a {
		_, ok := m[n]
		if !ok {
			m[n] = 1
		} else {
			m[n] = m[n] + 1
		}
	}
	for i, mm := range m {
		if mm%2 == 1 {
			return i
		}
	}
	return 0
}

// A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.
// Count the minimal number of jumps that the small frog must perform to reach its target.

func jump(x, y, d int) int {
	q := (y - x) / d
	m := (y - x) % d
	if m == 0 {
		return q
	}
	return q + 1
}

// An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.
// Your goal is to find that missing element.
// Write a function:
// func Solution(A []int) int
// that, given an array A, returns the value of the missing element.
// For example, given array A such that:
//
//	A[0] = 2
//	A[1] = 3
//	A[2] = 1
//	A[3] = 5
//
// the function should return 4, as it is the missing element.
func missing(a []int) int {
	n := len(a)
	if n == 0 {
		return 1
	}
	m := make(map[int]struct{}, n)
	for _, aa := range a {
		m[aa] = struct{}{}
	}
	for i := 1; i < n+2; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0
}

// A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.
// Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].
// The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|
// In other words, it is the absolute difference between the sum of the first part and the sum of the second part.
// For example, consider array A such that:
//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3
// We can split this tape in four places:
// P = 1, difference = |3 − 10| = 7
// P = 2, difference = |4 − 9| = 5
// P = 3, difference = |6 − 7| = 1
// P = 4, difference = |10 − 3| = 7

func mixAbs(a []int) int {
	sum := func(a []int) int {
		var sum int
		for _, aa := range a {
			sum += aa
		}
		return sum
	}
	totalSum := sum(a)
	min := math.MaxInt
	leftSum := 0
	for p := 0; p < len(a)-1; p++ {
		leftSum += a[p]
		rightSum := totalSum - leftSum
		candidate := leftSum - rightSum
		if candidate < 0 {
			candidate = -candidate
		}
		if min > candidate {
			min = candidate
		}
	}
	return min
}

// A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.
// You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.
// The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.
// For example, you are given integer X = 5 and array A such that:
//
//	A[0] = 1
//	A[1] = 3
//	A[2] = 1
//	A[3] = 4
//	A[4] = 2
//	A[5] = 3
//	A[6] = 5
//	A[7] = 4
//
// In second 6, a leaf falls into position 5. This is the earliest time when leaves appear in every position across the river.
// Write a function:
// func Solution(X int, A []int) int
// that, given a non-empty array A consisting of N integers and integer X, returns the earliest time when the frog can jump to the other side of the river.
// If the frog is never able to jump to the other side of the river, the function should return −1.
func jumpLeaf(x int, a []int) int {
	m := make(map[int]struct{})
	for i, aa := range a {
		if aa <= x {
			m[aa] = struct{}{}
			if len(m) == x {
				return i
			}
		}
	}
	return -1
}

// A non-empty array A consisting of N integers is given.
// A permutation is a sequence containing each element from 1 to N once, and only once.
// For example, array A such that:
//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
//     A[3] = 2
// is a permutation, but array A such that:
//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
// is not a permutation, because value 2 is missing.
// The goal is to check whether array A is a permutation.
//
// Write a function:
// func Solution(A []int) int
// that, given an array A, returns 1 if array A is a permutation and 0 if it is not.
// For example, given array A such that:
//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
//     A[3] = 2
// the function should return 1.

// Given array A such that:
//
//	A[0] = 4
//	A[1] = 1
//	A[2] = 3
//
// the function should return 0.
func isPermutation(a []int) bool {
	return jumpLeaf(len(a), a) != -1
}

func minimumProduct(a []int) int {
	sort.Ints(a)
	len := len(a)
	max := a[len-1] * a[len-2] * a[len-3]
	if a[0] < 0 {
		if a[0]*a[1]*a[len-1] > max {
			max = a[0] * a[1] * a[len-1]
		}
	}
	return max
}

func parkingLotFee(e, l string) int {
	et, err := convertTime(e)
	if err != nil {
		panic(err)
	}
	lt, err := convertTime(l)
	if err != nil {
		panic(err)
	}
	var fee = 2 + 3
	if lt.Sub(et) < time.Hour {
		return fee
	}
	et = et.Add(time.Hour)
	for lt.Sub(et) > time.Second {
		fee += 4
		et = et.Add(time.Hour)
	}
	return fee
}

func convertTime(str string) (time.Time, error) {
	s := strings.Split(str, ":")
	hh, err := strconv.Atoi(s[0])
	if err != nil {
		return time.Time{}, err
	}
	ss, err := strconv.Atoi(s[1])
	if err != nil {
		return time.Time{}, err
	}
	t := time.Date(0, 0, 0, hh, ss, 0, 0, time.Local)
	return t, nil
}

func symmetryPoint(s string) int {
	len := len(s)
	if len == 0 {
		return -1
	}
	if len == 1 {
		return 0
	}
	if len%2 == 0 {
		return -1
	}

	mid := (len - 1) / 2

	l := s[:mid]
	r := s[mid+1:]

	for i := 0; i < mid; i++ {
		if l[i] != r[mid-i-1] {
			return -1
		}
	}

	return mid
}

func isProperlyNested(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	st := make([]rune, 0, length)
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			st = append(st, r)
		} else {
			if len(st) == 0 {
				return false
			}
			p := st[len(st)-1]
			if r == ')' && p != '(' {
				return false
			}
			if r == '}' && p != '{' {
				return false
			}
			if r == ']' && p != '[' {
				return false
			}
			st = st[:len(st)-1]
		}
	}

	return len(st) == 0
}
