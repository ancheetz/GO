package main

import "fmt"

func main() {

	// N := 10
	// mySlice = create([]int,N)
	// for i := 0, i <N, i ++ {
	//	mySlice[i]
	// }
	// fmt.Println(mySlice)

	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(n)

	for index, i := range n {
		if i%2 != 0 {
			fmt.Println(index, " is Odd")
		} else {
			fmt.Println(index, " is Even")
		}
	}

}
