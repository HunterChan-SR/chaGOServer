package util

import (
	"slices"
	"strings"
)

const (
	key = "kjsdbaiufguiassdhu3ihf7843ghowienfu934gbuewhc9293rh9bgu942bgu9"
)

func Encrypt(code string) string {
	res := []byte(code)
	for i := 0; i < len(res); i++ {
		res[i]++
	}
	s1 := string(res)
	slices.Reverse(res)
	s2 := string(res)
	return s2 + key + s1
}
func Decrypt(code string) string {
	strs := strings.Split(code, key)
	if len(strs) != 2 {
		return ""
	}
	res := []byte(strs[1])
	for i := 0; i < len(res); i++ {
		res[i]--
	}
	return string(res)
}
