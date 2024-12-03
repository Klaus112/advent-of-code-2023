package common

import "strings"

const (
	InputFile = "files/input.txt"
	TestFile  = "files/test.txt"
)

// HashIntersect returns the intersecting elements between a nd b
//
// Taken from here: https://github.com/juliangruber/go-intersect/blob/master/intersect.go
func HashIntersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func RemoveEmpty(in string) string {
	in, _ = strings.CutPrefix(in, " ")
	in, _ = strings.CutSuffix(in, " ")

	return in
}
