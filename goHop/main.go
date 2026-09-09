package main

func main() {
	p:=3
	q:=10
	for ni := 1; ni <= q; ni++ {
		if ni%p != 0 {
			println(ni)
		} else {
			for i := 1; i <= ni/p; i++ {
				print("Hop ")
			}
			println()
		}
	}
}

