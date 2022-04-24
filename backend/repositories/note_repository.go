package repositories

import (
	"github.com/google/uuid"
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteRepository interface {
	FindNote(id uuid.UUID) (*entities.Note, error)
}
