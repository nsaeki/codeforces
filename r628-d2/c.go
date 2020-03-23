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
	// buf := make([]byte, 10000)
	// scanner.Buffer(buf, 1000000)
	return scanner
}

func scanInt() int {
	if sc.Scan() {
		t := sc.Text()
		if v, err := strconv.Atoi(t); err == nil {
			return v
		}
		panic(fmt.Sprintln("Could't scan int from input", t))
	}
	panic(sc.Err())
}

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

var sc *bufio.Scanner = newScanner()

type edge struct {
	to, label int
}

func solve() {
	n := scanInt()
	g := make([][]*edge, n)
	e := make([]*edge, n)

	for i := 0; i < n; i++ {
		f, t := scanInt(), scanInt()
		a := edge{t, nil}
		e[i] := &a
		g[f] = append(g[f], &a)
		g[t] = append(g[t], &edge{f, nil})
	}

}

func main() {
	t := scanInt()
	for i := 0; i < t; i++ {
		solve()
	}
}
