package main

import "fmt"

// Set

type Set map[int]bool

func NewSet(vars ...int) Set {
	var s = make(Set)
	for _, v := range vars {
		s[v] = true
	}
	return s
}

func SetKeys(x Set) []int {
	var keys []int

	for k, _ := range x {
		keys = append(keys, k)
	}

	return keys
}

func Add(x Set, v int) { x[v] = true }

func Exist(x Set, v int) bool { return x[v] }

func Empty(x Set) bool { return len(x) == 0 }

func Union(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		res[k] = true
	}

	for k, _ := range y {
		res[k] = true
	}

	return res
}

func Difference(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		if !y[k] {
			res[k] = true
		}
	}

	return res
}

func Intersection(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		if y[k] {
			res[k] = true
		}
	}

	return res
}

func IsEqual(x, y Set) bool {
	var (
		x_keys            = SetKeys(x)
		y_keys            = SetKeys(y)
		intersection_keys = SetKeys(Intersection(x, y))
	)

	return len(x_keys) == len(intersection_keys) && len(y_keys) == len(intersection_keys)
}

func IsSubset(x, y Set) bool { return len(Difference(y, x)) == 0 }

func main() {
	var (
		s1 = NewSet(0, 1, 2)
		s2 = NewSet(1, 2, 3)
		s3 = NewSet(2, 3, 4)
		s4 = NewSet(1, 2, 3, 4)
	)

	fmt.Println(Union(s1, s2))
	fmt.Println(Intersection(Intersection(s1, s2), s3))
	fmt.Println(Difference(Difference(s1, s2), s3))
	fmt.Println(IsSubset(s2, s4))

}
