package main

import "fmt"

//每一个数都可以分解成素数的乘积，例如 84 = 2^2 * 3^1 * 5^0 * 7^1 * 11^0 * 13^0 * …
//令 x = 2^m0 * 3^m1 * 5^m2 * 7^m3 * 11^m4 * …
//令 y = 2^n0 * 3^n1 * 5^n2 * 7^n3 * 11^n4 * …
//如果 x 整除 y（y mod x == 0），则对于所有 i，mi <= ni(i = 0, 1, 2 ...)
//x 和 y 的最大公约数为：gcd(x,y) = 2^min(m0,n0) * 3^min(m1,n1) * 5^min(m2,n2) * ...
//x 和 y 的最小公倍数为：lcm(x,y) = 2^max(m0,n0) * 3^max(m1,n1) * 5^max(m2,n2) * ...

// 埃拉托斯特尼筛法在每次找到一个素数时，将能被素数整除的数排除掉
func countPrimes(n int) int {
	notPrime := make([]bool, n+1)
	cnt := 0
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		cnt++
		for j := 2 * i; j < n; j += i {
			notPrime[j] = true
		}
	}
	return cnt
}

// 求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 求最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	ans := countPrimes(10)
	fmt.Println(ans)
}
