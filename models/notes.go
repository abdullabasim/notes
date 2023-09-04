package models

import "time"

type Note struct {
	ID        int        `db:"id" json:"id"`
	Title     string     `db:"title" json:"title"`
	Text      string     `db:"text" json:"text"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type NoteCreated struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type NoteUpdateFields struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type NoteDeleteIds struct {
	IDs []int `json:"ids"`
}

//DTO
type NoteSerializer struct {
	ID        int       `json:"ID"`
	Title     string    `json:"Title"`
	Text      string    `json:"Text"`
	CreatedAt time.Time `json:"CreatedAt"`
}
