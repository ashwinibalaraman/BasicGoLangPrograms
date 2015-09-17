package fibonacci

func Fib(fibo int) int {

	if fibo == 0 || fibo == 1 {
		return fibo

	}

	fibo = Fib(fibo-1) + Fib(fibo-2)
	return fibo

}
