package data

import "time"

type Buffer struct {
	Filename      string    `json:"filename"`
	FileExtension string    `json:"file_extension"`
	LastModified  time.Time `json:"last_modified"`
	LastAuthor    string    `json:"last_author"`
}

type Change struct {
	PosX      int
	PosY      int
	Text      string
	Author    string
	Timestamp time.Time
}
