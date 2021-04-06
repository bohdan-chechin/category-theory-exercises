package main

func main() {
	memo_s := memoize(square)
	println(memo_s(5))
	println(memo_s(6))
	println(memo_s(5))

}

func square(a int) int {
	return a*a + 5
}

func memoize(f func(int) int) func(int) int {
	memo := make(map[int]int)
	return func(a int) int {
		if r, ok := memo[a]; ok == true {
			return r
		}
		memo[a] = f(a)
		return memo[a]
	}
}
