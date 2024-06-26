package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

func CreateVote(db *sql.DB, props *entity.Vote) (string, error) {
	var lastVote entity.Vote
	err := db.QueryRow(`SELECT * FROM votes WHERE "ip_address" = $1`, props.Ip_address).Scan(&lastVote.Id, &lastVote.Ip_address, &lastVote.Voted_at, &lastVote.Voted_in, &lastVote.Poll_id)
	if err != nil {
		fmt.Println("Error happens: ", err)
	}

	validTimeToVote := utils.ValidateTime(lastVote.Voted_at)

	if !validTimeToVote {
		return "Ip already vote in the last 24h", err
	}

	newVote := entity.NewVote(props.Ip_address, props.Poll_id, props.Voted_in)

	_, error := db.Exec("INSERT INTO votes (id, ip_address, voted_at, voted_in,poll_id) VALUES ($1, $2, $3, $4, $5)",
		newVote.Id, newVote.Ip_address, newVote.Voted_at, newVote.Voted_in, newVote.Poll_id)

	if error != nil {
		fmt.Println("Error has happen on create_vote_service: ", err)
	}

	return "Success", error
}
