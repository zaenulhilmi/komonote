package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/komonote/entities"
	"github.com/zaenulhilmi/komonote/handlers"
)

func Test_Handler_Returns_201_Created_When_Resource_Created(t *testing.T) {

	table := []struct {
		name     string
		body     string
		expected entities.Note
	}{
		{
			name:     "Create_New_Note",
			body:     `{"title":"test","content":"test"}`,
			expected: entities.Note{Title: "test", Content: "test"},
		},
		{
			name:     "Create_New_Note_2",
			body:     `{"title":"Test 2","content":"Test 2"}`,
			expected: entities.Note{Title: "Test 2", Content: "Test 2"},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/notes", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			noteHandler := handlers.NewNoteHandler()
			handler := http.HandlerFunc(noteHandler.CreateNote)
			handler.ServeHTTP(recorder, request)

			assert.Equal(t, 201, recorder.Code)
			expectedResult, _ := tt.expected.MarshalJSON()
			assert.JSONEq(t, string(expectedResult), recorder.Body.String())

		})
	}
}
