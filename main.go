package main

import "fmt"

func main(){
	var a, b int 
	var v string 

	fmt.Println("calculator")

	fmt.Printf("=================================\n")

	fmt.Println("")

	fmt.Printf("nomar #1:")
	fmt.Scanln(&a)

	fmt.Println("")

	fmt.Printf("nomer #2:")
	fmt.Scanln(&b)

	fmt.Println("")

	fmt.Printf("+, -, *, /:")
	fmt.Scanln(&v)

	fmt.Println("")

	fmt.Scanln(a, b, v)	

	fmt.Printf("=================================\n")

	loc(a, b, v)
}
func loc(a int, b int, v string){
	fmt.Println("")
	if v == "+"{
		fmt.Println(a, "+", b, "=", a + b)
	}else if v == "-"{
		fmt.Println(a, "-", b, "=", a - b)
	}else if v == "*"{
		fmt.Println(a, "*", b, "=", a * b)
	}else if v == "/"{
		fmt.Println(a, "/", b, "=", a / b)
	}
	fmt.Println("")
	fmt.Printf("=================================\n")
	main()
}
