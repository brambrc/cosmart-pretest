package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	EditionNumber int    `json:"edition_number"`
}

type Schedule struct {
	Book       Book      `json:"book"`
	PickupTime time.Time `json:"pickup_time"`
}

type LibraryService struct {
	Books     []Book
	Schedules []Schedule
}

func (ls *LibraryService) FetchBooksByGenre(genre string) ([]Book, error) {
	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json", genre)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result struct {
		Works []struct {
			Title   string `json:"title"`
			Authors []struct {
				Name string `json:"name"`
			} `json:"authors"`
			EditionCount int `json:"edition_count"`
		} `json:"works"`
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	books := make([]Book, 0)
	for _, work := range result.Works {
		authors := ""
		for _, author := range work.Authors {
			authors += author.Name + ", "
		}
		authors = authors[:len(authors)-2]

		book := Book{
			Title:         work.Title,
			Author:        authors,
			EditionNumber: work.EditionCount,
		}
		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryService) GetBooksByGenre(c *gin.Context) {
	genre := c.Query("genre")

	books, err := ls.FetchBooksByGenre(genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ls.Books = books

	c.JSON(http.StatusOK, books)
}

func (ls *LibraryService) SubmitSchedule(c *gin.Context) {
	var schedule Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ls.Schedules = append(ls.Schedules, schedule)

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule Book Pickup Submitted Successfully!"})
}

func (ls *LibraryService) GetBookInformationAndSchedule(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	editionNumberStr := c.Query("edition_number")

	editionNumber, err := strconv.Atoi(editionNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid edition number"})
		return
	}

	var bookSchedule Schedule
	for _, s := range ls.Schedules {
		if s.Book.Title == title && s.Book.Author == author && s.Book.EditionNumber == editionNumber {
			bookSchedule = s
			break
		}
	}

	if bookSchedule.Book != (Book{}) && !bookSchedule.PickupTime.IsZero() {
		c.JSON(http.StatusOK, bookSchedule)
		return
	}

	c.Status(http.StatusNotFound)
}

func (ls *LibraryService) GetScheduledBooks(c *gin.Context) {
	c.JSON(http.StatusOK, ls.Schedules)
}

func main() {
	libraryService := &LibraryService{
		Books:     make([]Book, 0),
		Schedules: make([]Schedule, 0),
	}

	router := gin.Default()

	router.GET("/books", libraryService.GetBooksByGenre)
	router.POST("/schedule", libraryService.SubmitSchedule)
	router.GET("/book-info", libraryService.GetBookInformationAndSchedule)
	router.GET("/scheduled-books", libraryService.GetScheduledBooks)

	log.Fatal(router.Run(":8080"))
}
