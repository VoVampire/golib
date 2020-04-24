package main

func main() {
	start := make(chan int, 1)
	end := start

	for i := 0; i < 100; i++ {
		next := make(chan int, 1)
		go func(in, out chan int) {
			for {
				out <- <-in
			}
		}(end, next)
		end = next
	}

	start <- 1
	for {
		start <- <-end
	}
}
