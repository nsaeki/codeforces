package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n, x := scanInt(), scanInt()
		a := scanInts(n)
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		sum := int64(0)
		i := 0
		for ; i < n; i++ {
			sum += int64(a[i])
			if sum/int64(i+1) < int64(x) {
				break
			}
		}
		fmt.Println(i)
	}
}
