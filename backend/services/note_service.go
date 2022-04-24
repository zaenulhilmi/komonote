package services

import (
	"github.com/google/uuid"
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteService interface {
	CreateNote(title, content string) (*entities.Note, error)
	GetNote(id uuid.UUID) (*entities.Note, error)
}

func NewNoteService() NoteService {
	return &noteService{}
}

type noteService struct {
}

// CreateNote implements NoteService
func (*noteService) CreateNote(title string, content string) (*entities.Note, error) {
	panic("unimplemented")
}

// GetNote implements NoteService
func (*noteService) GetNote(id uuid.UUID) (*entities.Note, error) {
    var note *entities.Note
    return note,nil
}
