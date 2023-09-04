package notes_test

import (
	"bytes"
	"encoding/json"
	"github.com/Masterminds/squirrel"
	"net/http"
	"net/http/httptest"
	"notesTask/database"
	"notesTask/models"
	routes "notesTask/routing"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewNote(test *testing.T) {
	// Initialize the router
	router := routes.SetupRouter()

	// Create a mock Note object
	note := models.Note{
		Title: "noteTitle",
		Text:  "noteText",
	}

	// Convert the Note object to JSON
	noteJSON, err := json.Marshal(note)
	assert.NoError(test, err)

	// Create a request with the JSON data
	req, err := http.NewRequest("POST", "/api/v1/note", bytes.NewBuffer(noteJSON))
	assert.NoError(test, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, req)

	// Check the response status code
	assert.Equal(test, http.StatusOK, responseRecorder.Code)

	// Decode the response JSON
	var response map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.NoError(test, err)

	// Check the response fields
	assert.Equal(test, "Note created successfully", response["message"])

	data, ok := response["data"].(map[string]interface{})
	assert.True(test, ok)

	assert.NotNil(test, data["ID"])
	assert.Equal(test, note.Title, data["Title"])
	assert.Equal(test, note.Text, data["Text"])
	assert.NotNil(test, data["CreatedAt"])
}
func TestCreateNoteWithMissingInput(test *testing.T) {
	// Initialize the router
	router := routes.SetupRouter()

	// Create a mock Note object with missing fields
	note := models.Note{
		Title: "", // Missing title
		Text:  "Text",
	}

	// Convert the Note object to JSON
	noteJSON, err := json.Marshal(note)
	assert.NoError(test, err)

	// Create a request with the JSON data
	req, err := http.NewRequest("POST", "/api/v1/note", bytes.NewBuffer(noteJSON))
	assert.NoError(test, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, req)

	// Check the response status code
	assert.Equal(test, http.StatusBadRequest, responseRecorder.Code)

	// Decode the response JSON
	var response map[string]string
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.NoError(test, err)

	// Check the response fields
	assert.Equal(test, "Both title and text are required", response["error"])
}

func TestCreateNoteDb(t *testing.T) {
	// Initialize and connect to the test database
	db := database.InitDB()
	defer db.Close()

	// Start a transaction for this test
	tx, err := db.Begin()
	assert.NoError(t, err)

	// Create a new Squirrel query builder
	qb := squirrel.Insert("notes").
		Columns("title", "text").
		Values("testTitle", "testText").
		Suffix("RETURNING id").
		RunWith(tx).
		PlaceholderFormat(squirrel.Dollar)

	// Execute the query and scan the result
	var noteID int
	err = qb.QueryRow().Scan(&noteID)
	assert.NoError(t, err)
	assert.NotZero(t, noteID)

	// Rollback the transaction to clean up
	tx.Rollback()
}

func TestMain(m *testing.M) {

	if err := os.Chdir(".."); err != nil {
		panic(err)
	}
	db := database.InitDB()
	defer db.Close()

	// Run tests
	m.Run()
}
