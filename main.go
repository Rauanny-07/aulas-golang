package main

import ("fmt")

func AnalizarNotas(nota1, nota2 float64) (float64, string){
    media := (nota1 + nota2) / 2
    var resultado string

    if media >= 6 {
        resultado = "aprovado"
    }else{
        resultado = "reprovado"
    }
      return media, resultado
}

func main(){
  media, resultado := AnalizarNotas (7.5, 5.5)
    fmt.Println("media", media)
    fmt.Println("resultado", resultado)
}