package main

import (
	"fmt"

)

func dadosPessoais(idade int, nome string) (int, string){

  if idade >= 18 {
    return idade, "maior de idade"
} else {
     return idade,"menor de idade"
}
   
  }

  func main(){

    var idade int
    var nome string

  fmt.Println("Seu nome é:")
   fmt.Scan(&nome)
   fmt.Println("Sua idade é")
   fmt.Scan(&idade)
    
  var nomep, idadep = dadosPessoais(idade, nome)
fmt.Println(nomep)
fmt.Println(idadep)
}