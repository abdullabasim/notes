package notesRepository

import (
	"github.com/Masterminds/squirrel"
	"notesTask/database"
	"notesTask/models"
	"time"
)

type NotesRepository struct{}

func NewNotesRepository() *NotesRepository {
	return &NotesRepository{}
}

func (r *NotesRepository) CreateNote(note *models.Note) (*models.Note, error) {
	query := squirrel.Insert("notes").
		Columns("title", "text", "created_at").
		Values(note.Title, note.Text, time.Now()).
		Suffix("RETURNING id").
		RunWith(database.InitDB()).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRow().Scan(&note.ID)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *NotesRepository) GetNotes(page, limit int) ([]models.NoteSerializer, error) {
	offset := (page - 1) * limit

	var notes []models.NoteSerializer
	query := squirrel.Select("id", "title", "text", "created_at").
		From("notes").
		OrderBy("id DESC").
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		RunWith(database.InitDB())

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.NoteSerializer
		if err := rows.Scan(&note.ID, &note.Title, &note.Text, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *NotesRepository) UpdateNote(id string, updatedFields *models.NoteUpdateFields) error {
	now := time.Now()
	updateQuery := squirrel.Update("notes").
		Set("title", updatedFields.Title).
		Set("text", updatedFields.Text).
		Set("updated_at", now).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	_, err := updateQuery.RunWith(database.InitDB()).Exec()
	return err
}

func (r *NotesRepository) DeleteNote(ids []int) error {
	query := squirrel.Delete("notes").
		Where(squirrel.Eq{"id": ids}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(database.InitDB())

	_, err := query.Exec()
	return err
}
