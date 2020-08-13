package fib

//斐波那契数列
func Fibonacci() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}