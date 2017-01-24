package main


const (
	c = iota
	a = "abc"
	b = len(a)
	d = iota
)

const (
	k1 = 1 >>iota
	k2 = 3 << iota
	k3
	k4
)

func main() {
	println(a,b,c,d)
	println(k1,k2,k3,k4)
}
