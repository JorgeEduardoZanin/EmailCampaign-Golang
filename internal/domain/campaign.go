package domain

import (
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	Id        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

func Constructor(name string, content string, contacts []string) *Campaign {

	emails := make([]Contact, len(contacts))

	for i, email := range contacts {
		emails[i] = Contact{Email: email}
	}

	return &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  emails,
	}

}
