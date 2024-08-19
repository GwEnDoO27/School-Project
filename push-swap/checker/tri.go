package main

func pa(a, b []int) ([]int, []int) {
	var c []int
	c = append(c, b[0])
	c = append(c, a...)
	b = b[1:]

	return c, b
}

func pb(a, b []int) ([]int, []int) {
	var c []int
	c = append(c, a[0])
	c = append(c, b...)
	a = a[1:]

	return a, c
}

func sa(a, b []int) ([]int, []int) {
	a[0], a[1] = a[1], a[0]

	return a, b
}

func sb(a, b []int) ([]int, []int) {
	b[0], b[1] = b[1], b[0]

	return a, b
}

func ss(a, b []int) ([]int, []int) {
	a, b = sa(a, b)
	a, b = sb(a, b)

	return a, b
}

func ra(a, b []int) ([]int, []int) {
	a = append(a, a[0])
	a = a[1:]

	return a, b
}

func rb(a, b []int) ([]int, []int) {
	b = append(b, b[0])
	b = b[1:]

	return a, b
}

func rr(a, b []int) ([]int, []int) {
	a, b = ra(a, b)
	a, b = rb(a, b)

	return a, b
}

func rra(a, b []int) ([]int, []int) {
	var c []int
	c = append(c, a[len(a)-1])
	c = append(c, a[:len(a)-1]...)

	return c, b
}

func rrb(a, b []int) ([]int, []int) {
	var c []int
	c = append(c, b[len(b)-1])
	c = append(c, b[:len(b)-1]...)

	return a, c
}

func rrr(a, b []int) ([]int, []int) {
	a, b = rra(a, b)
	a, b = rrb(a, b)

	return a, b
}
