package entities

import "time"

type category struct {
	Id         uint
	Name       string
	CreatedAt  time.Time
	UpdateAt   time.Time
}