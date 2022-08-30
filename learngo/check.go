package main

type shape interface {
	getNumSides() int
	getArea() int
}

type square struct {
	len int
}

func (s square) getNumSides() int {
	return 4
}

func (s square) getArea() int {
	return s.len * 2
}

func main() {
	//Verify that *square implement shape

}
