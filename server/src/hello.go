//Packages
package main

import(
	"fmt"
	"math/cmplx"
	"math"
) 

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	p complex128 = cmplx.Sqrt(-5+12i)
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println("Happy", math.Pi, "Day")

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(x, y, z, c, python, java)

	const Truth = true
	fmt.Println("Go rules?", Truth)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, p, p)

}
