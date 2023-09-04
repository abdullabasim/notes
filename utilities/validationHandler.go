package utilities

import (
	"github.com/Masterminds/squirrel"
	"net/http"
	"notesTask/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ValidateNote checks if the provided note has valid fields
func ValidateTitleTextNote(c *gin.Context, title, text string) bool {
	if title == "" || text == "" {
		ErrorResponse(c, http.StatusBadRequest, "Both title and text are required")
		return false
	}

	if len(title) > 100 {
		ErrorResponse(c, http.StatusBadRequest, "Title should be less than 100 characters")
		return false
	}

	if len(text) > 250 {
		ErrorResponse(c, http.StatusBadRequest, "Text should be less than 250 characters")
		return false
	}

	return true
}

// IsInteger checks if the given string can be converted to an integer
func IsInteger(c *gin.Context, s string) (bool, error) {
	_, err := strconv.Atoi(s)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Invalid ID")
		return false, err
	}
	return true, nil
}

// Check if the provided note IDs exist in the database
func AreNotesIDsValid(ids []int) bool {
	// Perform a database query to check if the IDs exist
	query := squirrel.Select("id").From("notes").
		Where(squirrel.Eq{"id": ids})

	sql, args, _ := query.PlaceholderFormat(squirrel.Dollar).ToSql()

	rows, err := database.InitDB().Query(sql, args...)
	if err != nil {

		return false
	}
	defer rows.Close()

	// Create a map to store the existing IDs for quick lookup
	existingIDs := make(map[int]bool)

	// Iterate through the rows and populate the map
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {

			return false
		}
		existingIDs[id] = true
	}

	// Check if all provided IDs exist in the map
	for _, id := range ids {
		if !existingIDs[id] {
			return false
		}
	}

	return true
}
