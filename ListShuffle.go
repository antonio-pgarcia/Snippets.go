package main

import(
    "fmt"
    "math/rand"
)

// -- | Append two slices
func myappend(s1, s2 []int) []int {
	for i := range s2 {
		s1= append(s1, s2[i])
	}
	return s1	
}

// -- | Removes the the nth element from list
func drop(i int, l []int) [] int {
	return myappend(l[:i],l[i+1:])
}

// -- | Pop one element from a specific list position
func pop(i int, l []int) (int, []int) {
	return l[i], l
}

// -- | Generate random number within interval a-b with seed s
func myrandom(a, b, s int) int {
    rand.Seed(int64(s))
    return rand.Intn(b - a) + a
}

// -- | Shuffle' auxiliary
func shufflep() {
}

// -- | The main entry point.
func main() {
	var l= []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("List Shuffle!")
	//fmt.Println(myappend(l,l[:i]))
	fmt.Println(drop(4,l))
}
