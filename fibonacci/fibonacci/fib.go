package fibonacci

var fibMem Memoized

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}

func FibMemoized(n int) int {
	fibMem = Memoize(fib)
	return fibMem(n)
}

func fibStream(ch chan int, counter int) {
	defer close(ch)
	n1, n2 := 0, 1
	for i := 0; i < counter; i++ {
		ch <- n2
		n1, n2 = n2, n1+n2
	}
}

func FibChanneled(n int) int {
	ch := make(chan int)
	result := 0
	go fibStream(ch, n)

	i := 0
	for num := range ch {
		result = num
		i++
	}

	return result
}
