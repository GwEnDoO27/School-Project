package main

// This function return true if the array of int is already sorted else she return false
func IsTrie(a []int) bool {
	var b []int
	b = append(b, a...)

	TriABulle(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// this function make a bubblesort
func TriABulle(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for y := len(a)-1; y > i; y-- {
			if a[i] > a[y] {
				a[i], a[y] = a[y], a[i]
			}
		}
	}
}