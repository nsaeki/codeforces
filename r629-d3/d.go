package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		solve()
	}
}

func solve() {
	n := scanInt()
	m := make(map[int]int)
	res := make([]int, n)
	max := 1
	for i := 0; i < n; i++ {
		f := scanInt()
		if _, ok := m[f]; ok {
			res[i] = m[f]
		} else {
			res[i] = max
			m[f] = max
			max++
		}
	}
	fmt.Println(len(m))
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}
