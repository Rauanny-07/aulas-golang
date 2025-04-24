package main

import "fmt"

func main(){
  votos := make(map[string]int)


  votos["Ana"] += 1
  votos["Carlos"] += 1
  votos["Ana"] += 1
  votos["Bruno"] += 1
  votos["Ana"] += 1
  votos["Ana"] += 1
  votos["Carlos"] += 1

  for candidato, total := range votos {
    fmt.Println("%s resebeu %d votos\n", candidato, total)
  }
}