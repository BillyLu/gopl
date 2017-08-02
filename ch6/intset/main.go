package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const bitLen int = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitLen, uint(x%bitLen)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitLen, uint(x%bitLen)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i <= len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i <= len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

// SymmetricDifference sets s to the SymmetricDifference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i := range s.words {
		if i <= len(t.words) {
			s.words[i] ^= t.words[i]
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLen; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitLen*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Elems returns a slice of all set elems
func (s *IntSet) Elems() []int {
	elems := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLen; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, bitLen*i+j)
			}
		}
	}
	return elems
}

// Len return the number of elements
func (s *IntSet) Len() (t int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < bitLen; i++ {
			if word&(1<<uint(i)) != 0 {
				t++
			}
		}
	}
	return
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bitLen, uint(x%bitLen)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= (1 << bit)
	return
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}
	new.words = make([]uint, len(s.words))
	copy(new.words, s.words)
	return new
}

// AddAll adds all x to set
func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false
	fmt.Println(x.Len(), y.Len())
	x.Remove(1)
	fmt.Println(x.String(), x.Len())
	z := x.Copy()
	fmt.Println(z.String(), z.Len())
	x.Clear()
	fmt.Println(x.String(), x.Len())
	fmt.Println(z.String(), z.Len())
	fmt.Println(z.Elems())
}
