package main

// basic algorithms

import (
	"fmt"
	"math/rand"
)

// SEARCH ALGORITHMS

func linearSearch(nb int, slice []int) bool {
	for i := range slice {
		if slice[i] == nb {
			return true
		}
	}
	return false
}

// SORTING ALGORITHMS

func bubbleSort(tab []int) {
	for i := 0; i < len(tab)-1; i++ {
		for j := 0; j < len(tab)-i-1; j++ {
			if tab[j] > tab[j+1] {
				tab[j], tab[j+1] = tab[j+1], tab[j]
			}
		}
	}
}

func selectionSortGetMin(tab []int) (int, int) {
	min := tab[0]
	idx := 0

	for i := range tab {
		if tab[i] < min {
			min = tab[i]
			idx = i
		}
	}

	return min, idx
}

func selectionSort(tab []int) {
	for i := range tab {
		min, idxMin := selectionSortGetMin(tab[i:])
		idxMin = idxMin + i //correction car on ne passe pas l'integralite du tableau a la fonction selectionSort_getMin()

		for j := idxMin; j > i; j-- {
			tab[j-1], tab[j] = tab[j], tab[j-1]
		}

		tab[i] = min
	}
}

func insertionSort(tab []int) {
	for i := 1; i < len(tab); i++ {
		nb := tab[i]
		j := i - 1
		for j >= 0 && tab[j] > nb {
			tab[j+1] = tab[j]
			j--
		}
		tab[j+1] = nb
	}
}

func mergeSortMerger() {

}

func mergeSort(tab []int) {
	if len(tab) >= 1 {

		middle := len(tab) / 2

		fmt.Println(tab)

		mergeSort(tab[middle:])
		mergeSort(tab[:middle])
	}
}

func test(tab []int) {
	tab[0], tab[1] = 666, 666
}

func main() {
	tab := make([]int, 20)

	for i := range tab {
		tab[i] = rand.Intn(30)
	}

	fmt.Println(tab)
	mergeSort(tab)
	fmt.Println(tab)

}
