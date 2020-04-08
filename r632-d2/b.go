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
		a := scanInts(n)
		b := scanInts(n)
		minus, plus := false, false
		for i := n - 1; i >= 0; i-- {
			if a[i] < 0 {
				minus = false
			} else if a[i] > 0 {
				plus = false
			}
			if a[i] > b[i] {
				minus = true
			} else if a[i] < b[i] {
				plus = true
			}
		}
		if minus || plus {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
