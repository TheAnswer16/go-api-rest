package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/TheAnswer16/go-api-rest/models"
	"github.com/gorilla/mux"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}

}
