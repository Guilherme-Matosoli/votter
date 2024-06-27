package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/google/uuid"
)

type Vote struct {
	Id         string
	Ip_address string
	Voted_at   time.Time
	Poll_id    string
}

func NewVote(ip_address string, poll_id string) *Vote {
	return &Vote{
		Id:         uuid.New().String(),
		Ip_address: ip_address,
		Voted_at:   utils.GetTime(),
		Poll_id:    poll_id,
	}
}
