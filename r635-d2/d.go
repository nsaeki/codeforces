package main

import (
	"bufio"
	"fmt"
	"math"
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
		n := scanInts(3)
		w := make([][]int, 3)
		for i := 0; i < 3; i++ {
			w[i] = scanInts(n[i])
		}

		for i := 0; i < 3; i++ {
			sort.Ints(w[i])
		}


		f := func(i, j, k int) int64 {
			r, g, b := int64(w[0][i]), int64(w[1][j]), int64(w[2][k])
			return (r-g)*(r-g) + (g-b)*(g-b) + (b-r)*(b-r)
		}

		min := func(a, b int64) int64 {
			if a < b {
				return a
			}
			return b
		}

		var ans int64 = math.MaxInt64
		for i := 0; i < n[0]; i++ {
			j := sort.Search(n[1], func(j int) bool { return w[1][j] >= w[0][i] })
			if j >= 0 && j < n[1] {
				k := sort.Search(n[2], func(k int) bool { return w[2][k] >= (w[0][i]+w[1][j])/2 })
				if k < n[2] {
					ans = min(ans, f(i, j, k))
				}
				if k > 0 {
					ans = min(ans, f(i, j, k-1))
				}
			}
			for j2 := j-1; j2 >= 0; j2-- {
				if j < n[1] && w[1][j2] == w[1][j] {
					continue
				}
				k := sort.Search(n[2], func(k int) bool { return w[2][k] >= (w[0][i]+w[1][j2])/2 })
				if k < n[2] {
					ans = min(ans, f(i, j2, k))
				}
				if k > 0 {
					ans = min(ans, f(i, j2, k-1))
				}
				break
			}
			for j2 := j+1; j2 < n[1]; j2++ {
				if j < n[1] && w[1][j2] == w[1][j] {
					continue
				}
				k := sort.Search(n[2], func(k int) bool { return w[2][k] >= (w[0][i]+w[1][j2])/2 })
				if k < n[2] {
					ans = min(ans, f(i, j2, k))
				}
				if k > 0 {
					ans = min(ans, f(i, j2, k-1))
				}
				break
			}
		}
		fmt.Println(ans)
	}
}
