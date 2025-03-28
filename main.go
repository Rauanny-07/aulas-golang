package main

import "fmt"

func main(){
	a, b :=10, 3
	fmt.Println("a soma é:", a + b)
	fmt.Println("a divisão é:", a / b)
	fmt.println("a mutiplicação é:", a * b)
	fmt.println("o resto da divisão: ", a % b)


    a+=1
	fmt.println("incrementar a", a)

	if a > 0 && b > 0 {
		fmt.println("numeros podetivos")
	}

	}


