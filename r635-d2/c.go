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

type city struct {
	id int
	e  []int
	c  int
	t  bool
	d  int
}

type cities []city

func (a cities) Len() int      { return len(a) }
func (a cities) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a cities) Less(i, j int) bool {
	return a[i].c < a[j].c || (a[i].c == a[j].c && a[i].d > a[j].d)
}

var (
	seen []bool
	c    cities
)

func dfs(i, d int) int {
	if seen[i] {
		return 0
	}
	seen[i] = true
	c[i].d = d
	for _, j := range c[i].e {
		c[i].c += dfs(j, d+1)
	}
	return c[i].c + 1
}

func dfs2(i, x int) int {
	if seen[i] {
		return 0
	}
	seen[i] = true
	ret := 0
	if c[i].t {
		x += 1
	} else {
		ret += x
	}
	for _, j := range c[i].e {
		ret += dfs2(j, x)
	}
	return ret
}

func main() {
	n, k := scanInt(), scanInt()
	c = make(cities, n)
	for i := 0; i < n; i++ {
		c[i].id = i
		c[i].e = []int{}
	}
	for i := 0; i < n-1; i++ {
		u, v := scanInt(), scanInt()
		u--
		v--
		c[u].e = append(c[u].e, v)
		c[v].e = append(c[v].e, u)
	}

	seen = make([]bool, n)
	dfs(0, 0)
	d := make(cities, n)
	copy(d, c)
	sort.Sort(d)
	for i := len(d) - 1; i >= k; i-- {
		c[d[i].id].t = true
	}

	seen = make([]bool, n)
	ans := dfs2(0, 0)
	fmt.Println(ans)
}
