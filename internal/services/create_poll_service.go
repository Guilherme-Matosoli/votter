package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreatePoll(db *sql.DB, pollTitle string, questions []entity.Question) (string, error) {
	newPoll := entity.NewPoll(pollTitle)

	_, err := db.Exec("INSERT INTO polls (id, title, created_at) VALUES ($1,$2,$3)", newPoll.Id, newPoll.Title, newPoll.Created_at)
	if err != nil {
		fmt.Println("Error in create poll service: ", err)
		return "", err
	}

	for _, question := range questions {
		props := &entity.Question{Poll_id: newPoll.Id, Option: question.Option}

		_, err := CreateQuestion(db, props)
		if err != nil {
			fmt.Println("Error in create_poll_service: ", err)
		}
	}

	return newPoll.Id, err
}
