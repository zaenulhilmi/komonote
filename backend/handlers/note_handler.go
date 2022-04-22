package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zaenulhilmi/komonote/entities"
)

type NoteHandler interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
	FindNote(w http.ResponseWriter, r *http.Request)
}

type noteHandler struct{}

func NewNoteHandler() NoteHandler {
	return &noteHandler{}
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *noteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" {
		http.Error(w, "Title and Content are required", http.StatusBadRequest)
		return
	}

	note := &entities.Note{
		Title:   req.Title,
		Content: req.Content,
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := note.MarshalJSON()

	w.Write([]byte(response))
}

func (n *noteHandler) FindNote(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
