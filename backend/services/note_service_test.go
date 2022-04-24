package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/komonote/services"
)

func Test_Service_Get_Service_Not_Found(t *testing.T) {
	service := services.NewNoteService()
	note, err := service.GetNote(uuid.New())
	assert.Nil(t, note)
	assert.Nil(t, err)
}
