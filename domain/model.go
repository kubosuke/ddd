package domain

import (
	"time"
	"errors"
)

const (
	// Status
	TODO Status = iota
	WAIT Status = iota
	DOING Status = iota
	DONE Status = iota
	CANCEL Status = iota
	// Role
	Developer Role = "Developer"
	Administrator Role = "Administrator"
)

type Status int

func (s Status) String() string {
	switch s {
	case TODO:
		return "TODO"
	case WAIT:
		return "WAIT"
	case DOING:
		return "DOING"
	case DONE:
		return "DONE"
	case CANCEL:
		return "CANCEL"
	default:
		return "Unknown"
	}
}

type Role string

type Task struct {
	Id string
	Status Status
	DueDate time.Time
	Title string
	Description string
}

type Guild struct {
	Id string
	Members []Member
	Task Task
}

type Member struct {
	Id string
	belongGuildId string
	name Name
	Tasks []Task
	Role Role
}

type Name struct {
	givenName string
	familyName string
}

func NewName(givenName,familyName string) (*Name,error) {
	if len(givenName)<2 || len(familyName)<2 {
		return nil,errors.New("[ERROR]NewName(): name needs more than 2 characters.")
	}
	return &Name{givenName: givenName, familyName: familyName},nil
}
