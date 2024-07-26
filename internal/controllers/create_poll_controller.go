package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/services"
)

type createPollRequestBody struct {
	Info      entity.Poll       `json:"info"`
	Questions []entity.Question `json:"questions"`
}

type response struct {
	Message string `json:"message"`
}

type pollid struct {
	Id string `json:"id"`
}

func CreatePollController(w http.ResponseWriter, r *http.Request) {
	conn, error := database.Connection()
	if error != nil {
		fmt.Println("error try connect in db happen: ", error)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}
	defer conn.Close()

	var input createPollRequestBody

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}

	id, err := services.CreatePoll(conn, &input.Info, input.Questions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}

	responseMessage, err := json.Marshal(&pollid{Id: id})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseMessage)
}
