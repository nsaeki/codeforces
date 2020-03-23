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
	mx := 10000005
	maxk := 0
	sum := make([]int, mx/2)
	sum[0], sum[1] = 0, 1
	for i := 2; i < mx/2; i++ {
		sum[i] = sum[i-1] + 2*i - 1
		if sum[i] >= mx {
			maxk = 2*i - 1
			break
		}
	}

	for i := 0; i < t; i++ {
		n, k := scanInt(), scanInt()
		res := false
		if k < maxk {
			if n&1 == 0 && k&1 == 0 && n >= sum[k] {
				res = true
			}
			if n&1 == 1 && k&1 == 1 && n >= sum[k] {
				res = true
			}
		}

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
