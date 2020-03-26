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

type vertex struct {
	parent, depth int
	ancesters     []int
}

func main() {
	n, m := scanInt(), scanInt()
	v := make([]*vertex, n+1)
	v[1] = &vertex{0, 0, []int{}}
	for i := 1; i < n; i++ {
		v1, v2 := scanInt(), scanInt()
		if v[v1] == nil {
			v1, v2 = v2, v1
		}
		p := []int{v2}
		x := v1
		for v[x] != nil {
			p = append(p, x)
			x = v[x].parent
		}
		v[v2] = &vertex{v1, v[v1].depth + 1, p}
	}

	for i := 0; i < m; i++ {
		res := true
		k := scanInt()
		u := scanInts(k)
		mx := 0
		md := 0
		for _, x := range u {
			if md < v[x].depth {
				mx = x
				md = v[x].depth
			}
		}

		fmt.Println(v[mx].ancesters)
		for _, x := range u {
			c := false
			for _, y := range v[mx].ancesters {
				if x == y || v[x].parent == y {
					c = true
					break
				}
			}
			if !c {
				res = false
				break
			}
		}
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
