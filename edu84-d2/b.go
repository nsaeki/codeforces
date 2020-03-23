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

var sc = newScanner()

func solve() {
	n := scanInt()
	d := make([]bool, n+1)
	k := make([]bool, n+1)
	cnt := 0
	for i := 1; i <= n; i++ {
		m := scanInt()
		for j := 0; j < m; j++ {
			x := scanInt()
			if d[i] || k[x] {
				continue
			} else {
				d[i] = true
				k[x] = true
				cnt++
			}
		}
	}
	if cnt == n {
		fmt.Println("OPTIMAL")
	} else {
		di, ki := 0, 0
		for i := 1; i <= n; i++ {
			if !d[i] {
				di = i
			}
		}
		for i := 1; i <= n; i++ {
			if !k[i] {
				ki = i
			}
		}
		fmt.Println("IMPROVE")
		fmt.Println(di, ki)
	}
}

func main() {
	t := scanInt()
	for i := 0; i < t; i++ {
		solve()
	}
}
