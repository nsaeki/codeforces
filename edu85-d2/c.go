package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type intArr []int64

func (a intArr) Len() int           { return len(a) }
func (a intArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intArr) Less(i, j int) bool { return a[i] < a[j] }

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(r, &n)
		a := make(intArr, n)
		b := make(intArr, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i], &b[i])
		}
		sum := max(a[0]-b[n-1], 0)
		for i := 1; i < n; i++ {
			sum += max(a[i]-b[i-1], 0)
		}
		sort.Sort(a)
		if sum == 0 {
			fmt.Fprintln(w, a[0])
		} else {
			fmt.Fprintln(w, sum+a[0])
		}
	}
}
