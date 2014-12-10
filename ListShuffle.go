// ListShuffle
package main

import (
	"fmt"
	"math/rand"
)

// -- | Append two slices
func myappend(s1, s2 []int) []int {
	for i := range s2 {
		s1 = append(s1, s2[i])
	}
	return s1
}

// -- | Removes the the nth element from list
func drop(i int, l []int) []int {
	if len(l) == 0 {
		return []int{}
	}
	return myappend(l[:i], l[i+1:])
}

// -- | Pop one element from a specific list position
func pop(i int, l []int) (int, []int) {
	var e = 0
	if len(l) > 0 {
		e = l[i]
	}
	return e, drop(i, l)
}

// -- | Generate random number within interval a-b with seed s
func myrandom(a, b, s int) int {
	if a == 0 && b == 0 {
		return 0
	}
	rand.Seed(int64(s))
	return rand.Intn(b-a) + a
}

// -- | Shuffle a list (golang slice)
func shuffle(l []int) []int {
	return shufflep(myrandom(0, len(l)-1, 0), l, []int{})
}

// -- | Shuffle' auxiliary
func shufflep(r int, l1, l2 []int) []int {
	var llen = len(l1)
	h, t := pop(r, l1)
	if llen == 0 {
		return l2
	} else {
		//return l1
		return shufflep(myrandom(0, llen-1, r), t, append(l2, h))
	}
	return l2
}

// -- | The main entry point.
func main() {
	var v = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("List Shuffle!")
	fmt.Println(shuffle(v))
}
