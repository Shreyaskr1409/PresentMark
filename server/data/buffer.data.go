package data

import "time"

type Buffer struct {
	FileName      string
	FileExtension string
	lastModified  time.Time
	lastAuthor    string
}

type Change struct {
	PosX      int
	PosY      int
	Text      string
	Author    string
	Timestamp time.Time
}
