package main

import "fmt"

func partition(A []int, low, high int) int {
	pivot := A[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if A[j] <= pivot {
			i += 1
			A[i], A[j] = A[j], A[i]
		}
	}

	i += 1
	A[i], A[high] = A[high], A[i]

	return i
}

func quicksort(A []int, low, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partition(A, low, high)

	quicksort(A, low, p-1)
	quicksort(A, p+1, high)
}

func main() {
	A := []int{48, 2, 3, 555, 2, 1, 56, 72, 34, 88, 3, 1, 4, 7, 8, 56, 6, 7, 4, 22}
	fmt.Println(A)

	quicksort(A, 0, len(A)-1)

	fmt.Println(A)
}
