package main

import (
	"fmt"
)

func BubbleSort(sl []int) {
	for i := range sl {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				Swap(sl, j)
			} else {
				continue
			}
		}
	}
	fmt.Println(sl)
}

func Swap(sl []int, i int) {
	t := sl[i]
	sl[i] = sl[i+1]
	sl[i+1] = t
}

func main() {
	sli := make([]int, 10)
	fmt.Println("Enter 10 numbers")
	var num int
	for i := range sli {
		fmt.Scan(&num)
		sli[i] = num
	}
	BubbleSort(sli)
}
