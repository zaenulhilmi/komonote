package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/komonote/entities"
)

type NoteServiceMock struct {
	mock.Mock
}

func (m *NoteServiceMock) CreateNote(note *entities.Note) (*entities.Note, error) {
	args := m.Called(note)
	return args.Get(0).(*entities.Note), args.Error(1)
}

func (m *NoteServiceMock) GetNote(id string) (*entities.Note, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Note), args.Error(1)
}
