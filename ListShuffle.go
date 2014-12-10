package main

import(
    "fmt"
    "math/rand"
)

// -- | Pop one element from a specific list position
func pop(i int, l []int) (int, []int) {
	return i, l
}

// -- | Generate random number within interval a-b with seed s
func myrandom(a, b, s int) int {
    rand.Seed(int64(s))
    return rand.Intn(b - a) + a
}

// -- | The main entry point.
func main() {
	var list = []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Hello, playground", list, myrandom(1,100,0))
}
