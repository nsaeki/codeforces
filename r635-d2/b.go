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

func main() {
	for t:=scanInt(); t > 0; t-- {
		x, n, m := scanInt(), scanInt(), scanInt()
		for i := 0; i < n; i++ {
			if x <= 20 {
				break
			}
			x = x / 2 + 10
		}

		if x <= m * 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}