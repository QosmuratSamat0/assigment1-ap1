package Library

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books  map[string]*Book
	nextID int64
}

func New() *Library {
	return &Library{Books: make(map[string]*Book), nextID: 1000}
}

func (l *Library) AddBook(title, author string) {
	bookID := fmt.Sprintf("B%04d", l.nextID)
	l.nextID++
	book := &Book{
		ID:     bookID,
		Title:  title,
		Author: author,
	}
	l.Books[bookID] = book
	fmt.Println("Added new book:", book.ID)
}

func (l *Library) BorrowBook(id string) bool {
	b, ok := l.Books[id]
	if !ok || b.IsBorrowed {
		fmt.Println("Borrowed book:", id)
		return false
	}
	b.IsBorrowed = true
	l.Books[id] = b
	fmt.Printf("Book %s, %s is borrowed\n", b.ID, b.Title)
	return true
}

func (l *Library) ReturnBook(id string) error {
	b, ok := l.Books[id]
	if !ok || !b.IsBorrowed {
		return fmt.Errorf("Book not found or already borrowed")
	}
	b.IsBorrowed = false
	l.Books[id] = b
	fmt.Printf("Returned book:\nid: %s\ntitle: %s\nauthor: %s\nborrowed: %v", b.ID, b.Title, b.Author, b.IsBorrowed)
	return nil
}

func (l *Library) ListAvailableBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available")
		return
	}

	fmt.Println("\n--- Books ---")
	for _, b := range l.Books {
		status := "available"
		if b.IsBorrowed {
			status = "borrowed"
		}
		fmt.Printf("id: %s, title: %s, author: %s, status: %s\n",
			b.ID, b.Title, b.Author, status)
	}
}
