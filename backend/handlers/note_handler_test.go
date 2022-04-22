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


var noteHandler handlers.NoteHandler = handlers.NewNoteHandler()

func Test_Handler_Returns_201_Created_When_Resource_Created(t *testing.T) {
	for _, tt := range getValidCreateRequestTable() {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/notes", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			recorder := runHandler(request, noteHandler.CreateNote)

			assert.Equal(t, 201, recorder.Code)
			expectedResult, _ := tt.expected.MarshalJSON()
			assert.JSONEq(t, string(expectedResult), recorder.Body.String())

		})
	}
}

func Test_Handler_Return_400_BadRequest_When_Request_Empty_Or_Invalid_JSON(t *testing.T) {
	for _, tt := range getInvalidCreateRequestTable() {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/notes", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			recorder := runHandler(request, noteHandler.CreateNote)
			assert.Equal(t, 400, recorder.Code)
		})
	}

}

func Test_Handler_Returns_200_OK_When_Resource_Found(t *testing.T) {
    request, err := http.NewRequest("GET", "/notes", nil)
    if err != nil {
        t.Fatal(err)
    }
    recorder := runHandler(request, noteHandler.FindNote)

    assert.Equal(t, 200, recorder.Code)
}

func runHandler(request *http.Request, handlerFunc http.HandlerFunc) *httptest.ResponseRecorder {

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(recorder, request)

	return recorder
}

func getInvalidCreateRequestTable() []struct {
	name string
	body string
} {
	table := []struct {
		name string
		body string
	}{
		{
			name: "empty body",
			body: ``,
		},

		{
			name: "invalid json",
			body: `title`,
		},
		{
			name: "Create_New_Note_With_Empty_Title",
			body: `{"title":"","content":"test"}`,
		},
		{
			name: "Create_New_Note_With_Empty_Content",
			body: `{"title":"test","content":""}`,
		},
		{
			name: "Create_New_Note_With_Empty_Title_And_Content",
			body: `{"title":"","content":""}`,
		},
	}

	return table
}

func getValidCreateRequestTable() []struct {
	name     string
	body     string
	expected entities.Note
} {
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

	return table
}
