package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	var t int
	for fmt.Fscan(rd, &t); t > 0; t-- {
		var n, l, r int64
		fmt.Fscan(rd, &n, &l, &r)
		var k int64 = 1
		for k*(k-1)+1 < l {
			k++
		}

		ans := make([]int, 0, r-l+1)
		for i := l; i <= r; i++ {
			if i == k*(k-1)+1 {
				ans = append(ans, 1)
				k++
			} else if i%2 == 0 {
				ans = append(ans, int(k))
			} else {
				x := (i-(k-1)*(k-2))/2 + 1
				ans = append(ans, int(x))
			}
		}
		fmt.Println(strings.Trim(fmt.Sprint(ans), "[]"))
	}
}
