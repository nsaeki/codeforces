package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Scan(&n, &m)
		if n*m%2 == 0 {
			for i := 0; i < n; i++ {
				if i == 0 {
					for j := 0; j < m; j++ {
						if j == 0 {
							fmt.Fprint(w, "B")
						} else if j%2 == 0 {
							fmt.Fprint(w, "W")
						} else {
							fmt.Fprint(w, "B")
						}
					}
				} else if i%2 == 0 {
					for j := 0; j < m; j++ {
						if j%2 == 0 {
							fmt.Fprint(w, "W")
						} else {
							fmt.Fprint(w, "B")
						}
					}
				} else {
					for j := 0; j < m; j++ {
						if j%2 == 0 {
							fmt.Fprint(w, "B")
						} else {
							fmt.Fprint(w, "W")
						}
					}
				}
				fmt.Fprintln(w)
			}
		} else {
			for i := 0; i < n; i++ {
				if i%2 == 0 {
					for j := 0; j < m; j++ {
						if j%2 == 0 {
							fmt.Fprint(w, "B")
						} else {
							fmt.Fprint(w, "W")
						}
					}
				} else {
					for j := 0; j < m; j++ {
						if j%2 == 0 {
							fmt.Fprint(w, "W")
						} else {
							fmt.Fprint(w, "B")
						}
					}
				}
				fmt.Fprintln(w)
			}
		}
	}
}
