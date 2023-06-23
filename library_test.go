package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksByGenre(t *testing.T) {
	router := gin.New()
	libraryService := &LibraryService{
		Books:     make([]Book, 0),
		Schedules: make([]Schedule, 0),
	}
	router.GET("/books", libraryService.GetBooksByGenre)

	req, err := http.NewRequest("GET", "/books?genre=love", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSubmitSchedule(t *testing.T) {
	router := gin.New()
	libraryService := &LibraryService{
		Books:     make([]Book, 0),
		Schedules: make([]Schedule, 0),
	}
	router.POST("/schedule", libraryService.SubmitSchedule)

	scheduleJSON := `{
		"book": {
			"title": "Book Title",
			"author": "Author Name",
			"edition_number": 1
		},
		"pickup_time": "2023-06-22T12:00:00Z"
	}`

	req, err := http.NewRequest("POST", "/schedule", strings.NewReader(scheduleJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Contains(t, res.Body.String(), "Schedule Book Pickup Submitted Successfully!")
}

func TestGetBookInformationAndSchedule(t *testing.T) {
	router := gin.New()
	libraryService := &LibraryService{
		Books:     make([]Book, 0),
		Schedules: make([]Schedule, 0),
	}
	router.GET("/book-info", libraryService.GetBookInformationAndSchedule)

	schedule := Schedule{
		Book: Book{
			Title:         "Book Title",
			Author:        "Author Name",
			EditionNumber: 1,
		},
		PickupTime: time.Now(),
	}

	libraryService.Schedules = append(libraryService.Schedules, schedule)

	req, err := http.NewRequest("GET", "/book-info?title=Book Title&author=Author Name&edition_number=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetScheduledBooks(t *testing.T) {
	router := gin.New()
	libraryService := &LibraryService{
		Books:     make([]Book, 0),
		Schedules: make([]Schedule, 0),
	}
	router.GET("/scheduled-books", libraryService.GetScheduledBooks)

	schedule := Schedule{
		Book: Book{
			Title:         "Book Title",
			Author:        "Author Name",
			EditionNumber: 1,
		},
		PickupTime: time.Now(),
	}

	libraryService.Schedules = append(libraryService.Schedules, schedule)

	req, err := http.NewRequest("GET", "/scheduled-books", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestMain(m *testing.M) {
	m.Run()
}
