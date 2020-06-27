package main

import (
	"fmt"
	"strings"
)

type IntSet struct {
	words []uint64
}

func (this *IntSet) Has(x int) bool {
	idx, bit := x/64, uint(x%64)
	return idx < len(this.words) && this.words[idx]&(1<<bit) != 0
}

func (this *IntSet) Add(x int) {
	idx, bit := x/64, uint(x%64)
	for idx >= len(this.words) {
		this.words = append(this.words, 0)
	}
	this.words[idx] |= (1 << bit)
}

func (this *IntSet) UnionWith(other *IntSet) {
	for idx, v := range other.words {
		if idx < len(this.words) {
			this.words[idx] |= v
		} else {
			this.words = append(this.words, v)
		}
	}
}

func (this *IntSet) String() string {
	var b strings.Builder
	b.WriteByte('{')
	for idx, v := range this.words {
		if v == 0 {
			continue
		}
		for offset := 0; offset < 64; offset++ {
			if v&(1<<offset) != 0 {
				if b.Len() > len("{") {
					b.WriteByte(',')
				}
				fmt.Fprintf(&b, "%d", 64*idx+offset)
			}
		}
	}
	b.WriteByte('}')
	return b.String()
}

// return the number of elements
func (this *IntSet) Len() int {
	 size  := 0
	for _, v := range this.words {
		for j := 0; j < 64; j++ {
			if v & (1<<j) !=0 {
				size++
			}
		}
	}
	return size
}

// remove x from the set
func (this *IntSet) Remove(x int) {
	idx, bit := x/64, uint(x%64)
	if idx < len(this.words) {
		this.words[idx] &= ^(1 << bit)
	}
}

// remove all elements from the set
func (this *IntSet) Clear() {
	this.words = make([]uint64, 0)
}

// return a copy of the set
func (this *IntSet) Copy() *IntSet {
	dst:= make([]uint64, len(this.words))
	copy(dst, this.words)
	return &IntSet{dst}
}

func main() {
	var x, y, z IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	fmt.Printf("Len of x is %v\n", x.Len())
	fmt.Printf("Len of y is %v\n", y.Len())
	fmt.Printf("Len of z is %v\n", z.Len())

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9)) // "true"
	x.Remove(9)
	fmt.Println(x.Has(9)) // "true false"
	fmt.Printf("Len of x is %v\n", x.Len())

	w := x.Copy()
	x.Clear()
	fmt.Println("After clearing ", x.String())
	fmt.Println("After clearing ", w.String())
}
