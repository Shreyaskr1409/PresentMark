package handlers

import (
	"log"
	"net/http"
)

type FileHandler struct {
	l *log.Logger
}

func InitFileHandler(l *log.Logger) *FileHandler {
	return &FileHandler{
		l: l,
	}
}

func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request)
func (h *FileHandler) CreateFile(w http.ResponseWriter, r *http.Request)
func (h *FileHandler) UpdateFile(w http.ResponseWriter, r *http.Request)
func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request)
