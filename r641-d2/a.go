package main

import "fmt"

func main() {
	n := int(1e6)
	data := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if data[i] != 0 {
			continue
		}
		data[i] = i
		if i > n/i {
			continue
		}
		for j := i * i; j <= n; j += i {
			if data[j] == 0 {
				data[j] = i
			}
		}
	}

	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)
		ans := int64(n + data[n])
		k--
		if k > 0 {
			ans += int64(2 * k)
		}
		fmt.Println(ans)
	}
}
