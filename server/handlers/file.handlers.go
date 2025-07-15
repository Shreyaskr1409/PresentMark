package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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

type GetFileRequest struct {
	Filename string `json:"filename"`
}

func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	var req GetFileRequest
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

	file, err := os.Open(fp)
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error opening the requested file: ", err),
			http.StatusBadRequest,
		)
		h.l.Println(err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error obtaining file info: ", err),
			http.StatusInternalServerError,
		)
		h.l.Println(err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error sending file", http.StatusInternalServerError)
		h.l.Println(err)
		return
	}
}

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
		FileExtension: filepath.Ext(fp),
		LastModified:  time.Now(),
		LastAuthor:    "Author (HARDCODED)",
	}

	_, err = os.Create(fp)
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error marshalling json for response: ", err),
			http.StatusInternalServerError,
		)
		h.l.Println(err)
		return
	}

	h.l.Println("Received request")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(
			w,
			fmt.Sprint("error marshalling json for response: ", err),
			http.StatusInternalServerError,
		)
		h.l.Println(err)
		return
	}
}

type UpdateFileRequest struct {
	Filename string         `json:"filename"`
	Changes  []*data.Change `json:"changes"`
}

func (h *FileHandler) UpdateFile(w http.ResponseWriter, r *http.Request) {
	var req UpdateFileRequest
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

	byte_arr, _ := json.MarshalIndent(req, "", "	")
	h.l.Println(string(byte_arr))

	fp := req.Filename
	fp = filepath.Join("public/storage/", fp)
	fp, _ = filepath.Abs(fp)

	file, err := os.ReadFile(fp)
	if err != nil {
		h.l.Println("error opening file: ", err)
		return
	}
	lines := strings.Split(string(file), "\n")

	sort.Slice(req.Changes, func(i, j int) bool {
		return req.Changes[i].Timestamp.Before(req.Changes[j].Timestamp)
	})

	for _, change := range req.Changes {
		if len(change.Text) < 2 {
			continue
		}

		operation := change.Text[0]
		text := change.Text[1:]

		lineNum := change.PosY

		switch operation {
		case '+':
			if lineNum == len(lines) {
				newLine := text
				lines = append(lines, newLine)
			} else {
				currentLine := lines[lineNum]
				if change.PosX <= len(currentLine) {
					newLine := currentLine[:change.PosX] + text + currentLine[change.PosX:]
					lines[change.PosY] = newLine
				} else {
					// If position is beyond line length, append with spaces
					padding := strings.Repeat(" ", change.PosX-len(currentLine))
					lines[change.PosY] = currentLine + padding + text
				}
			}
		default:
			h.l.Println("unknown operation: ", operation, "for change ", change)
		}
	}

	newContent := strings.Join(lines, "\n")
	err = os.WriteFile(fp, []byte(newContent), 0o644)
	if err != nil {
		h.l.Println("error writing file: ", err)
		http.Error(w, "error writing file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Changes applied successfully"))
}

func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {}
