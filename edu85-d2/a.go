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
	for t := scanInt(); t > 0; t-- {
		n := scanInt()
		p0, c0 := scanInt(), scanInt()
		ok := p0 >= c0
		for i := 1; i < n; i++ {
			p, c := scanInt(), scanInt()
			if p < c || p < p0 || c < c0 || p-p0 < c-c0 {
				ok = false
			}
			p0, c0 = p, c
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
