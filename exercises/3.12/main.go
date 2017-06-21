package main

import "fmt"
import "strings"
import "sort"

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}

	s1 = sortStrings(s1)
	s2 = sortStrings(s2)

	if s1 == s2 {
		return true
	}
	return false
}

func sortStrings(s string) string {
	subs := strings.Split(s, "")
	sort.Strings(subs)
	return strings.Join(subs, "")
}

func main() {
	a := "strings"
	b := "grinsts"
	fmt.Println(isAnagram(a, b)) // true

	a = "baz"
	b = "zab"
	fmt.Println(isAnagram(a, b)) // true

	a = "bar"
	b = "baz"
	fmt.Println(isAnagram(a, b)) // false
}
