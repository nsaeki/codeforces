package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(r, &n, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}

		nk := 0
		ok := false
		// kが存在してk以上の数を増やすことができれば良いので、
		// k以上の値が連続、もしくは一つ飛びで存在することが示せれば、すべてkで埋め尽くせる
		for i := 0; i < n; i++ {
			if a[i] == k {
				nk++
			}
			if i > 0 && a[i] >= k && a[i-1] >= k {
				ok = true
			} else if i > 1 && a[i] >= k && a[i-2] >= k {
				ok = true
			}
		}
		if nk <= 0 {
			ok = false
		}
		if n == 1 && a[0] == k {
			ok = true
		}

		if ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
