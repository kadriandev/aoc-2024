package utils

import (
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Str_to_int(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}
