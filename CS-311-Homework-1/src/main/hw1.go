package main

//Joshua Dunne 20225_CS_311_01_1

import "fmt"
import "math"

//a function compPoly that takes 4 parameters: 
//a, b, c and x and returns ax2 + bx + c

func compPoly(a, b, c, x float64) float64 {
	term1 := x * x * a
	term2 := x * b
	term3 := c
	return term1 + term2 + term3
}

//A function calcRoots that takes 3 parameters: 
//a, b, c and uses the quadratic formula to return the two roots

func calcRoots(a, b, c float64) (root1, root2 float64) {
	root1 = ((-b) - math.Sqrt((b * b) - (4 * a * c))) / (2 * a)
	root2 = ((-b) + math.Sqrt((b * b) - (4 * a * c))) / (2 * a)
	return
}

//A function find_close that return list of numerical values in seq 
//that are close enough to the target parameter.

func find_close(toFind float64, toSearch []float64, delta float64) (found []float64) {
	for i := 0; i < len(toSearch); i++ {
		if math.Abs(toSearch[i] - toFind) <= delta {
			found = append(found, toSearch[i])
		}
	}
	return
}

func main() {
	fmt.Println(compPoly(3, 4, 5, 6));
	fmt.Println(calcRoots(2, -11, 5));
	fmt.Println(find_close(10, []float64 {11, 7, 8, 0, 20}, 2))
}

