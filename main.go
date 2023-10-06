package main

import (
	"fmt"
	"github.com/TheAnswer16/go-api-rest/routes"
	"github.com/TheAnswer16/go-api-rest/models"
	"github.com/TheAnswer16/go-api-rest/database"
)

func main () {

	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Nome 1", Historia: "Exemplo"},
		{Id: 2, Nome: "Nome 2", Historia: "Exemplo"},
	}

	database.ConectarComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

