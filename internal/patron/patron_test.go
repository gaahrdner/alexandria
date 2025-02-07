package patron

import (
	"reflect"
	"testing"
)

func TestPatronLessons(t *testing.T) {
	t.Run("Lesson02", func(t *testing.T) {
		t.Run("TestPatron_NewPatron", TestPatron_NewPatron)
		t.Run("TestPatron_CheckOutBook", TestPatron_CheckOutBook)
		t.Run("TestPatron_ReturnBook", TestPatron_ReturnBook)
		t.Run("TestPatron_HasCheckedOutBook", TestPatron_HasCheckedOutBook)
		t.Run("TestPatron_NewPatron_EmptyName", TestPatron_NewPatron_EmptyName) // Edge Case
	})
	t.Run("Lesson11", func(t *testing.T) {
		t.Run("TestPatron_String", TestPatron_String)
	})
}

// --- Lesson 02: Patron Management (Slices) and Basic Checkout/Return ---

func TestPatron_NewPatron(t *testing.T) {
	p := NewPatron("Alice Smith", 123)
	if p.Name != "Alice Smith" {
		t.Errorf("Expected name 'Alice Smith', got '%s'", p.Name)
	}
	if p.ID != 123 {
		t.Errorf("Expected ID 123, got %d", p.ID)
	}
	if len(p.CheckedOutBooks) != 0 {
		t.Errorf("Expected CheckedOutBooks to be empty, got length %d", len(p.CheckedOutBooks))
	}
}

func TestPatron_CheckOutBook(t *testing.T) {
	p := NewPatron("Bob Johnson", 456)
	bookISBN := "ISBN-123"
	p.CheckOutBook(bookISBN)
	if len(p.CheckedOutBooks) != 1 {
		t.Errorf("Expected 1 checked-out book, got %d", len(p.CheckedOutBooks))
	}
	if p.CheckedOutBooks[0] != bookISBN {
		t.Errorf("Expected checked-out book to be '%s', got '%s'", bookISBN, p.CheckedOutBooks[0])
	}
	p.CheckOutBook(bookISBN) // Check out same book again
	if len(p.CheckedOutBooks) != 1 {
		t.Errorf("Expected 1 checked-out book (no duplicates), got %d", len(p.CheckedOutBooks))
	}
}

func TestPatron_ReturnBook(t *testing.T) {
	p := NewPatron("Charlie Brown", 789)
	bookISBN1 := "ISBN-123"
	bookISBN2 := "ISBN-456"
	p.CheckOutBook(bookISBN1)
	p.CheckOutBook(bookISBN2)
	err := p.ReturnBook(bookISBN1)
	if err != nil {
		t.Errorf("Unexpected error returning book: %v", err)
	}
	if len(p.CheckedOutBooks) != 1 {
		t.Errorf("Expected 1 checked-out book, got %d", len(p.CheckedOutBooks))
	}
	if p.CheckedOutBooks[0] != bookISBN2 {
		t.Errorf("Expected checked-out book to be '%s', got '%s'", bookISBN2, p.CheckedOutBooks[0])
	}
	err = p.ReturnBook("ISBN-NonExistent")
	if err == nil {
		t.Errorf("Expected error when returning a non-checked-out book, got nil")
	}
	err = p.ReturnBook(bookISBN2)
	if err != nil {
		t.Errorf("Unexpected error returning second book. %v", err)
	}
	if len(p.CheckedOutBooks) != 0 {
		t.Errorf("Expected CheckedOutBooks to be empty, got len %d", len(p.CheckedOutBooks))
	}
}

func TestPatron_HasCheckedOutBook(t *testing.T) {
	p := NewPatron("Diana Davis", 101112)
	bookISBN1 := "ISBN-789"
	bookISBN2 := "ISBN-012"
	p.CheckOutBook(bookISBN1)
	if !p.HasCheckedOutBook(bookISBN1) {
		t.Errorf("Expected patron to have checked out book '%s'", bookISBN1)
	}
	if p.HasCheckedOutBook(bookISBN2) {
		t.Errorf("Expected patron not to have checked out book '%s'", bookISBN2)
	}
}
func TestPatron_NewPatron_EmptyName(t *testing.T) {
	p := NewPatron("", 0) // Test with empty name and ID 0.
	if p.Name != "" {
		t.Errorf("Expected empty name, got '%s'", p.Name)
	}
	if p.ID != 0 {
		t.Errorf("Expected ID 0, got %d", p.ID)
	}
}

// --- Lesson 11: Refactoring and UI Improvements ---
func TestPatron_String(t *testing.T) {
	p := NewPatron("Eve Evans", 131415)
	p.CheckOutBook("ISBN-345")
	p.CheckOutBook("ISBN-678")

	expectedString := "Patron: Eve Evans (ID: 131415)\nChecked Out Books:\n- ISBN-345\n- ISBN-678"
	actualString := p.String()

	if !reflect.DeepEqual(actualString, expectedString) {
		t.Errorf("Expected string representation:\n%s\nGot:\n%s", expectedString, actualString)
	}

	p2 := NewPatron("No Books", 1)
	expectedNoBooks := "Patron: No Books (ID: 1)\nChecked Out Books: None"
	actualNoBooks := p2.String()

	if !reflect.DeepEqual(actualNoBooks, expectedNoBooks) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedNoBooks, actualNoBooks)
	}
}
