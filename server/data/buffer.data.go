package data

import "time"

type Buffer struct {
	Filename      string    `json:"filename"`
	FileExtension string    `json:"file_extension"`
	LastModified  time.Time `json:"last_modified"`
	LastAuthor    string    `json:"last_author"`
}

type Change struct {
	PosX      int       `json:"pos_x"`
	PosY      int       `json:"pos_y"`
	Text      string    `json:"text"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}
