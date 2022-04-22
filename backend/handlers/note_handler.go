package handlers

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/zaenulhilmi/komonote/services"
)

type NoteHandler interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
	GetNote(w http.ResponseWriter, r *http.Request)
}

type noteHandler struct {
	service services.NoteService
}

func NewNoteHandler(service services.NoteService) NoteHandler {
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
	id := path.Base(r.URL.Path)

	resId, _ := uuid.Parse(id)

	n.service.GetNote(resId)

	w.WriteHeader(http.StatusOK)
}
