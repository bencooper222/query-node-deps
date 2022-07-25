package util

import "strings"

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// https://stackoverflow.com/questions/51598250/split-a-string-at-the-last-occurrence-of-the-separator-in-golang
func SplitOnLastAppearanceOfDelimiter(str string, delimiter string) (string, string) {
	lastIndex := strings.LastIndex(str, delimiter)

	return str[:lastIndex], str[lastIndex+1:]
}

func MapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
