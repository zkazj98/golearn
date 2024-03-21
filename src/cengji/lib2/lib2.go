package main

func init() {
	println("lib2 init")
}

const (
	Name1 = iota
)

func main() {
	q := [...]int{2, 3, 4}
	for i, v := range q {
		println(i)
		println(v)
	}
}

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
