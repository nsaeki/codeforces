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
	mx := 100000
	ans := make([]byte, mx)
	for i := 0; i < mx; i++ {
		ans[i] = 'a'
	}

	for i := 0; i < t; i++ {
		n, k := scanInt(), scanInt()
		x := 0
		s := 0
		for s < k {
			x++
			s = x * (x + 1) / 2
		}

		l, r := x, x-(s-k)-1
		ans[n-l-1] = 'b'
		ans[n-r-1] = 'b'
		fmt.Println(string(ans[:n]))
		ans[n-l-1] = 'a'
		ans[n-r-1] = 'a'
	}
}
