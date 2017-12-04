package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	fib := make([]int, 2)
	fib[0] = 0
	fib[1] = 1
	return func(index int) int {
		switch {
		case cap(fib) <= index:
			tmplen := cap(fib)
			for tmplen <= index {
				tmplen *= 2
			}
			tmp := make([]int, len(fib), tmplen)
			copy(tmp, fib)
			fib = tmp
			fallthrough
		case len(fib) <= index:
			for i := len(fib); i <= index; i++ {
				fib = append(fib, fib[i-1] + fib[i-2])
			}
			fallthrough
		default:
			return fib[index]
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(f(9-i))
	}
}

