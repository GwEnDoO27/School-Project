package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	tabArg := strings.Split(arg, " ")

	// verifie si il y a bine que des nombres qui sont ecrit
	var a, b []int
	for _, i := range tabArg {
		if _, err := strconv.Atoi(i); err != nil {
			fmt.Println("Error")
			return
		} else {
			conv, _ := strconv.Atoi(i)
			a = append(a, conv)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	var input string
	for scanner.Scan() {
		if scanner.Text() != "" {
			input += scanner.Text() + " "
		}
	}

	inputTab := strings.Split(input[:len(input)-1], " ")

	for _, v := range inputTab {
		if v == "pa" {
			a, b = pa(a, b)
		} else if v == "pb" {
			a, b = pb(a, b)
		} else if v == "sa" {
			a, b = sa(a, b)
		} else if v == "sb" {
			a, b = sb(a, b)
		} else if v == "ss" {
			a, b = ss(a, b)
		} else if v == "ra" {
			a, b = ra(a, b)
		} else if v == "rb" {
			a, b = rb(a, b)
		} else if v == "rr" {
			a, b = rr(a, b)
		} else if v == "rra" {
			a, b = rra(a, b)
		} else if v == "rrb" {
			a, b = rrb(a, b)
		} else if v == "rrr" {
			a, b = rrr(a, b)
		} else {
			fmt.Println("Error")
			return
		}
	}

	if len(b) != 0 {
		fmt.Println("KO")
		return
	} else if !IsTrie(a) {
		fmt.Println("KO")
		return
	} else {
		fmt.Println("OK")
		return
	}
}
