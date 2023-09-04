package notesController

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	notesRepository "notesTask/database/repository"
	"notesTask/models"
	responseHandler "notesTask/utilities"
	validationHandler "notesTask/utilities"
	"strconv"
)

// Home handles the main endpoint.
// @Summary Home endpoint
// @Description Returns a simple message for the main endpoint
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/ [get]
func Home(c *gin.Context) {
	responseHandler.SuccessResponse(c, "Welcome to Notes API : Index endpoint", nil)
}

// @Summary Create a new note
// @Description Creates a new note with a title and text
// @Accept json
// @Produce json
// @Param note body models.NoteCreated true "Note object"
// @Success 201 {object} models.NoteSerializer
// @Router /api/v1/note [post]
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid inputs")
		return
	}

	if !validationHandler.ValidateTitleTextNote(c, note.Title, note.Text) {
		return
	}

	// Initialize the repository
	repo := notesRepository.NewNotesRepository()

	// Call the CreateNote method from the repository
	createdNote, err := repo.CreateNote(&note)
	if err != nil {
		responseHandler.ErrorResponse(c, http.StatusInternalServerError, "Failed to create note")
		return
	}

	responseNote := models.NoteSerializer{
		ID:        createdNote.ID,
		Title:     createdNote.Title,
		Text:      createdNote.Text,
		CreatedAt: createdNote.CreatedAt,
	}

	responseHandler.SuccessResponse(c, "Note created successfully", responseNote)
}

// @Summary Get All notes with pagination
// @Description Retrieves a list of notes with pagination
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of notes per page"
// @Success 200 {array} models.Note
// @Router /api/v1/notes [get]
func GetNotes(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid limit")
		return
	}

	// Initialize the repository
	repo := notesRepository.NewNotesRepository()

	// Call the GetNotes method from the repository
	notes, err := repo.GetNotes(page, limit)
	if err != nil {
		responseHandler.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch notes")
		return
	}

	if len(notes) == 0 {
		responseHandler.SuccessResponse(c, "No data found", nil)
		return
	}

	responseHandler.SuccessResponse(c, "Notes retrieved successfully", notes)
}

// @Summary Update a note
// @Description Updates an existing note by ID
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param models.NoteUpdateFields body models.NoteUpdateFields true "Updated fields"
// @Router /api/v1/note/{id} [put]
func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	isValid, err := validationHandler.IsInteger(c, id)

	if !isValid {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var updatedFields models.NoteUpdateFields

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid inputs")
		return
	}

	if !validationHandler.ValidateTitleTextNote(c, updatedFields.Title, updatedFields.Text) {
		return
	}

	// Initialize the repository
	repo := notesRepository.NewNotesRepository()

	// Call the UpdateNote method from the repository
	err = repo.UpdateNote(id, &updatedFields) // Assign to the existing err variable

	if err != nil {
		if err == sql.ErrNoRows {
			responseHandler.ErrorResponse(c, http.StatusNotFound, "Note not found")
		} else {
			responseHandler.ErrorResponse(c, http.StatusInternalServerError, "Failed to update note")
		}
		return
	}

	responseHandler.SuccessResponse(c, "Note updated successfully", nil)
}

// @Summary Delete notes
// @Description Deletes one or more notes by IDs
// @Accept json
// @Produce json
// @Param noteIds body models.NoteDeleteIds true "Delete request"
// @Router /api/v1/notes [delete]
func DeleteNotes(c *gin.Context) {
	var noteIds models.NoteDeleteIds
	if err := c.ShouldBindJSON(&noteIds); err != nil {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid IDs")
		return
	}

	if len(noteIds.IDs) == 0 {
		responseHandler.ErrorResponse(c, http.StatusBadRequest, "Please enter at least 1 ID for note")
		return
	}

	// Validate if the IDs are valid integers
	validIDs := make([]int, 0)
	for _, idStr := range noteIds.IDs {
		isValid, _ := validationHandler.IsInteger(c, strconv.Itoa(idStr))

		if !isValid {
			responseHandler.ErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
			return
		}
		validIDs = append(validIDs, idStr)
	}

	if !validationHandler.AreNotesIDsValid(validIDs) {
		responseHandler.ErrorResponse(c, http.StatusNotFound, "One or more notes not found")
		return
	}

	// Initialize the repository
	repo := notesRepository.NewNotesRepository()

	// Call the DeleteNote method from the repository
	err := repo.DeleteNote(validIDs)
	if err != nil {
		if err == sql.ErrNoRows {
			responseHandler.ErrorResponse(c, http.StatusNotFound, "One or more notes not found")
		} else {
			responseHandler.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete notes")
		}
		return
	}

	responseHandler.SuccessResponse(c, "Notes deleted successfully", nil)
}
