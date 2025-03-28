package main

import "fmt"

func main(){
    var username string 

	fmt.Print("Digite o nome de usuário: ") 

	fmt.Scanln(&username) 

  

	// Solicitar a senha 

	var password string 

	fmt.Print("Digite a senha: ") 

	fmt.Scanln(&password) 

  

	// Verificar se o usuário e senha estão corretos 

	if username == "admin" && password == "1234" { 

		fmt.Println("Login bem-sucedido!") 

	} else { 

		fmt.Println("Usuário ou senha incorretos.") 

}


