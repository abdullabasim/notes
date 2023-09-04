package notesRepository

import "notesTask/models"

type Repository interface {
	CreateNote(note *models.Note) (*models.Note, error)
	GetNotes(page, limit int) ([]models.NoteSerializer, error)
	UpdateNote(id string, updatedFields *models.NoteUpdateFields) error
	DeleteNote(ids []int) error
}
