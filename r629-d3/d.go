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

var sc = newScanner()

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for t := scanInt(); t > 0; t-- {
		n := scanInt()
		f := make([]int, n)
		m := make(map[int]bool)
		r := -1
		for i := 0; i < n; i++ {
			f[i] = scanInt()
			m[f[i]] = true
			if r == -1 && i > 0 && f[i] == f[i-1] {
				r = i
			}
		}

		cnt := 0
		res := make([]int, n)
		if len(m) == 1 {
			cnt = 1
			for i := 0; i < n; i++ {
				res[i] = 1
			}

		} else if n&1 == 0 || f[0] == f[n-1] {
			cnt = 2
			res[0] = 1
			for i := 1; i < n; i++ {
				res[i] = res[i-1] ^ 3
			}
		} else if r > 0 {
			cnt = 2
			res[0] = 1
			for i := 1; i < n; i++ {
				res[i] = res[i-1] ^ 3
				if i == r {
					res[i] = res[i-1]
				}
			}
		} else {
			cnt = 3
			res[0], res[n-1] = 1, 3
			for i := 1; i < n-1; i++ {
				res[i] = res[i-1] ^ 3
			}
		}

		fmt.Println(cnt)
		fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
	}
}
