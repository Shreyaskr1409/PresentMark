package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Shreyaskr1409/PresentMark/data"
	"github.com/Shreyaskr1409/PresentMark/utils"
)

type FileHandler struct {
	l *log.Logger
}

func InitFileHandler(l *log.Logger) *FileHandler {
	return &FileHandler{
		l: l,
	}
}

func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request) {}

type CreateFileRequest struct {
	Filename string `json:"filename"`
}

func (h *FileHandler) CreateFile(w http.ResponseWriter, r *http.Request) {
	var req CreateFileRequest
	err := utils.ParseRequest(r, &req)
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error parsing request body: ", err),
			http.StatusBadRequest,
		)
		h.l.Println(err)
		return
	}

	fp := req.Filename
	fp = filepath.Join("public/storage/", fp)
	fp, _ = filepath.Abs(fp)
	res := data.Buffer{
		Filename:      fp,
		FileExtension: ".md",
		LastModified:  time.Now(),
		LastAuthor:    "Author (HARDCODED)",
	}

	h.l.Println("Received request")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error marshalling json for response: ", err),
			http.StatusInternalServerError,
		)
	}
}

func (h *FileHandler) UpdateFile(w http.ResponseWriter, r *http.Request) {}
func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {}
