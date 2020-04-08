package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)
	t := []byte(s)
	i := 0
	m := make([][]int, n*n+1)
	sum := 0
	for ; i <= n*n; i++ {
		j := 0
		for j < n-1 {
			if t[j] == 'R' && t[j+1] == 'L' {
				m[i] = append(m[i], j+1)
				t[j], t[j+1] = 'L', 'R'
				j += 2
			} else {
				j++
			}
		}
		sum += len(m[i])
		if len(m[i]) == 0 {
			break
		}
	}

	m = m[:i]
	if sum < k || i > k {
		fmt.Println(-1)
	} else {
		d := k - i
		for _, x := range m {
			if d > 0 {
				if d > len(x) {
					d -= len(x)
					for j := 0; j < len(x); j++ {
						fmt.Fprintln(w, 1, x[j])
					}
				} else {
					j := 0
					for ; j < d; j++ {
						fmt.Fprintln(w, 1, x[j])
						d--
					}
					if j < len(x) {
						fmt.Fprintln(w, len(x[j:]), strings.Trim(fmt.Sprint(x[j:]), "[]"))
					}
				}
			} else {
				fmt.Fprintln(w, len(x), strings.Trim(fmt.Sprint(x), "[]"))
			}
		}
	}
}
