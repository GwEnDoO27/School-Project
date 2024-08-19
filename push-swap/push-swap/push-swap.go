package main

import "fmt"

// inite la fonction push-swwap
func Push_Swap(a, b, ResFinal []int, lastFunction string) {
	if len(a) != 0 && !IsTrie(a) {
		min := a[0]

		for _, v := range a {
			if min > v {
				min = v
			}
		}

		if len(a) > 1 && a[0] > a[1] { // POUR SA
			a, b = sa(a, b)
			fmt.Println("sa")

			Push_Swap(a, b, ResFinal, "sa")
		} else if len(b) > 1 && b[0] > b[1] && IsTrie(b) { // POUR SB
			a, b = sb(a, b)
			fmt.Println("sb")

			Push_Swap(a, b, ResFinal, "sb")
		} else if len(a) > 1 && a[0] < a[1] && a[0] < a[len(a)-1] && lastFunction != "pa" { // POUR PB
			a, b = pb(a, b)
			fmt.Println("pb")

			Push_Swap(a, b, ResFinal, "pb")
		} else if len(a) > 1 && len(b) > 1 && a[0] > a[len(a)-1] && a[1] > a[len(a)-1] && b[0] < b[len(b)-1] && b[1] < b[len(b)-1] && lastFunction != "rr" { // POUR RRR
			a, b = rrr(a, b)
			fmt.Println("rrr")

			Push_Swap(a, b, ResFinal, "rrr")
		} else if len(a) > 1 && a[0] > a[len(a)-1] && a[1] > a[len(a)-1] && lastFunction != "ra" { // POUR RRA
			a, b = rra(a, b)
			fmt.Println("rra")

			Push_Swap(a, b, ResFinal, "rra")
		} else if len(b) > 1 && b[0] < b[len(b)-1] && b[1] < b[len(b)-1] && lastFunction != "rb" { // POUR RRB
			a, b = rrb(a, b)
			fmt.Println("rrb")

			Push_Swap(a, b, ResFinal, "rrb")
		} else if len(a) > 1 && len(b) > 1 && a[0] > a[1] && a[1] <= a[len(a)-1] && b[0] < b[1] && b[1] >= b[len(b)-1] && lastFunction != "rrr" { // POUR RR
			a, b = rr(a, b)
			fmt.Println("rr")

			Push_Swap(a, b, ResFinal, "rr")
		} else if len(a) > 1 && a[0] > a[1] && a[1] <= a[len(a)-1] && lastFunction != "rra" { // POUR RA
			a, b = ra(a, b)
			fmt.Println("ra")

			Push_Swap(a, b, ResFinal, "ra")
		} else if len(b) > 1 && b[0] < b[1] && b[1] >= b[len(b)-1] && lastFunction != "rrb" { // POUR RB
			a, b = rb(a, b)
			fmt.Println("rb")

			Push_Swap(a, b, ResFinal, "rb")
		} else if len(b) > 1 && b[0] < b[1] && lastFunction != "pb" { // POUR PA
			a, b = pa(a, b)
			fmt.Println("pa")

			Push_Swap(a, b, ResFinal, "pa")
		}
	} else if len(b) > 0 {
		a, b = pa(a, b)
		fmt.Println("pa")

		Push_Swap(a, b, ResFinal, "pa")
	}
}
