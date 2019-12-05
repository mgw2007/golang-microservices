package utils

import "sort"

//BubbleSort for sort using Bubble algorithm
// input {6,3,2,4,5,1}
// input {1,2,3,4,5,6}
func BubbleSort(elements []int) {
	keepRuninning := true
	for keepRuninning {
		keepRuninning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRuninning = true
			}
		}
	}

}

//Sort elements if < 1000 use bubble else use sort.Ints
func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
	} else {
		sort.Ints(elements)
	}

}
