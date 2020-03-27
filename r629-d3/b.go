package main

import (
	"bufio"
	"bytes"
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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for t := scanInt(); t > 0; t-- {
		n, k := scanInt(), scanInt()
		ans := bytes.Repeat([]byte{'a'}, n)
		x, s := 0, 0
		for s < k {
			x++
			s += x
		}
		l, r := x, x-(s-k)-1
		ans[n-l-1] = 'b'
		ans[n-r-1] = 'b'
		fmt.Fprintf(w, "%s\n", ans)
	}
}
