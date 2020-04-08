package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("Despues de Im in Foo debe ir esto")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("I'm in Foo")
}
func bar(){
	fmt.Println("Aca termina el programa")
}

//Control Flow:
// (1) Sequence
// (2) Loop; Iterative
// (3) Conditional
