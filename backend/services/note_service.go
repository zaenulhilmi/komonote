package services

import (
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteService interface {
	CreateNote(title, content string) (*entities.Note, error)
}
