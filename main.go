package main

import "fmt"

func dividir(dividendo int, divisor int) (int, string) {
  if divisor == 0 {
    return 0, "Erro na divição"
  }
   return dividendo / divisor, "Sem erro"

}


  func operaçaoBasica(a int, b int) (int, int, int){
      soma := a + b
      multiplicaçao := a * b
      subitraçao := a - b
      return soma, multiplicaçao, subitraçao

}

func main (){
  resultado, erro := dividir(10,2)

    if erro != "Sem erro"{
        fmt.Println(erro)
}else {    
      fmt.Println("O resultado da divisao é:", resultado , erro)
}

soma, mult, sub:= operaçaoBasica(10,2)
  fmt.Println(soma)
  fmt.Println(mult)
  fmt.Println(sub)
}