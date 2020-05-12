package main

import "fmt"

func pow(n, p int) int64 {
	if p == 0 {
		return 1
	}
	ret := pow(n, p/2)
	ret = ret * ret
	if p%2 == 1 {
		ret *= int64(n)
	}
	return ret
}

func main() {
	n := int(1e6)*2 + 10
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

	factors := func(n int) map[int]int {
		ret := make(map[int]int)
		for n > 1 {
			ret[data[n]]++
			n /= data[n]
		}
		return ret
	}

	fmt.Scan(&n)
	min1 := make(map[int]int)
	min2 := make(map[int]int)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		for k, v := range factors(a) {
			if min1[k] == 0 || min1[k] > v {
				min2[k] = min1[k]
				min1[k] = v
			} else if min2[k] == 0 || min2[k] > v {
				min2[k] = v
			}
			cnt[k]++
		}
	}

	ans := int64(1)
	for k, v := range cnt {
		if v == n-1 {
			ans *= pow(k, min1[k])
		}
		if v == n {
			ans *= pow(k, min2[k])
		}
	}
	fmt.Println(ans)
}
