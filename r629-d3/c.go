package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func newScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func scanString() string {
	if sc.Scan() {
		return sc.Text()
	}
	panic(sc.Err())
}

var sc = newScanner()

func main() {
	t := scanInt()
	for i := 0; i < t; i++ {
		n, x := scanInt(), scanString()
		solve(n, x)
	}
}

func solve(n int, x string) {
	a := make([]int, n)
	b := make([]int, n)
	found := false
	for i := 0; i < n; i++ {
		c := int(x[i] - '0')
		if found {
			a[i] = 0
			b[i] = c
			continue
		}
		switch c {
		case 2:
			a[i] = 1
			b[i] = 1
		case 1:
			found = true
			a[i] = 1
			b[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i])
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(b[i])
	}
	fmt.Println()
}
