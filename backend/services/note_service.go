package services

import (
	"github.com/google/uuid"
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteService interface {
	CreateNote(title, content string) (*entities.Note, error)
	GetNote(id uuid.UUID) (*entities.Note, error)
}
