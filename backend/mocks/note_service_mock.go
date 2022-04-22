package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteServiceMock struct {
	mock.Mock
}

func (m *NoteServiceMock) CreateNote(title, content string) (*entities.Note, error) {
	args := m.Called(title, content)
	return args.Get(0).(*entities.Note), args.Error(1)
}

func (m *NoteServiceMock) GetNote(id uuid.UUID) (*entities.Note, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Note), args.Error(1)
}
