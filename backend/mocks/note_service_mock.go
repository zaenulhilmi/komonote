package mocks

import (
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
