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

func main() {
	t := scanInt()
	for i := 0; i < t; i++ {
		n, k := scanInt(), scanInt()
		if n&1 == k&1 && n/k >= k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
