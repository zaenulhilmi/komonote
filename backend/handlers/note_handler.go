package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zaenulhilmi/komonote/entities"
)

type NoteService interface {
	CreateNote(title, content string) (*entities.Note, error)
}

type NoteHandler interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
	GetNote(w http.ResponseWriter, r *http.Request)
}

type noteHandler struct {
	service NoteService
}

func NewNoteHandler(service NoteService) NoteHandler {
	return &noteHandler{
		service: service,
	}
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

	note, err := n.service.CreateNote(req.Title, req.Content)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := note.MarshalJSON()

	w.Write([]byte(response))
}

func (n *noteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
