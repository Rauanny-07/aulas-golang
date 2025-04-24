package main

import ("fmt")

func PegarNome() (string, string){
      return "Barry", "Allen"
}



func main(){
    nome, sobrenome := PegarNome()
    fmt.Println(nome)
    fmt.Println(sobrenome)
}