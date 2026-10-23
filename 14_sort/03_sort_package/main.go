package main

import (
	"fmt"
	"sort"
)

func main() {

	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

	//1、sort升序排序
	sort.Ints(intList)
	sort.Float64s(floatList)
	sort.Strings(stringList)

	fmt.Println(intList)    // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(floatList)  // [3.14 4.2 5.9 10.2 12.4 27.81828 31.4 50.7 99.9]
	fmt.Println(stringList) // [a b c d f i w x y z]

	//2、sort降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Println(intList)    // [9 8 7 6 5 4 3 2 1 0]
	fmt.Println(floatList)  // [99.9 50.7 31.4 27.81828 12.4 10.2 5.9 4.2 3.14]
	fmt.Println(stringList) // [z y x w i f d c b a]
}
