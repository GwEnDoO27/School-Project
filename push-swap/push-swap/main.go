package main

import (
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

	var tabArgInt []int
	for _, i := range tabArg {
		if _, err := strconv.Atoi(i); err != nil {
			fmt.Println("Error")
			return
		} else {
			conv, _ := strconv.Atoi(i)
			tabArgInt = append(tabArgInt, conv)
		}
	}

	if IsInTab(tabArgInt) {
		fmt.Println("Error")
		return
	}

	if len(tabArgInt) == 1 {
		return
	}

	var resFinal []int
	resFinal = append(resFinal, tabArgInt...)

	TriABulle(resFinal)

	Push_Swap(tabArgInt, []int{}, resFinal, "")
}
