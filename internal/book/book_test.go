package book

import (
	"reflect"
	"testing"
)

func TestLessons(t *testing.T) {
	t.Run("Lesson01", func(t *testing.T) {
		t.Run("TestBook_NewBook", TestBook_NewBook)
		t.Run("TestBook_NewBook_EmptyFields", TestBook_NewBook_EmptyFields)
	})

	t.Run("Lesson02", func(t *testing.T) {
		t.Run("TestBook_CheckOut", TestBook_CheckOut)
		t.Run("TestBook_Return", TestBook_Return)
		t.Run("TestBook_IsAvailable", TestBook_IsAvailable)
	})

	t.Run("Lesson03", func(t *testing.T) {
		t.Run("TestBook_AddToWaitlist", TestBook_AddToWaitlist)
		t.Run("TestBook_RemoveFromWaitlist", TestBook_RemoveFromWaitlist)
		t.Run("TestBook_Waitlist_Integration", TestBook_Waitlist_Integration)
	})

	t.Run("Lesson11", func(t *testing.T) {
		t.Run("TestBook_String", TestBook_String)
	})
}

// --- Lesson 01: Project Setup and Basic Book Management (Slices) ---

func TestBook_NewBook(t *testing.T) {
	b := NewBook("The Lord of the Rings", "J.R.R. Tolkien", "978-0618002258", 1954)
	if b.Title != "The Lord of the Rings" {
		t.Errorf("Expected title 'The Lord of the Rings', got '%s'", b.Title)
	}
	if b.Author != "J.R.R. Tolkien" {
		t.Errorf("Expected author 'J.R.R. Tolkien', got '%s'", b.Author)
	}
	if b.ISBN != "978-0618002258" {
		t.Errorf("Expected ISBN '978-0618002258', got '%s'", b.ISBN)
	}
	if b.PublicationYear != 1954 {
		t.Errorf("Expected publication year 1954, got %d", b.PublicationYear)
	}
	if b.Available != true {
		t.Errorf("Expected Available to be true, got %v", b.Available)
	}
	if b.Waitlist == nil {
		t.Errorf("Expected Waitlist to be initialized, got nil")
	}
	if !b.Waitlist.IsEmpty() {
		t.Errorf("Expected Waitlist to be empty")
	}
}

func TestBook_NewBook_EmptyFields(t *testing.T) {
	b := NewBook("", "", "", 0)
	if b.Title != "" || b.Author != "" || b.ISBN != "" || b.PublicationYear != 0 {
		t.Errorf("Expected empty fields, got: %+v", b)
	}
}

// --- Lesson 02: Patron Management (Slices) and Basic Checkout/Return ---

func TestBook_CheckOut(t *testing.T) {
	b := NewBook("1984", "George Orwell", "978-0451524935", 1949)
	b.CheckOut()
	if b.Available != false {
		t.Errorf("Expected Available to be false after checkout, got %v", b.Available)
	}
	b.CheckOut()
	if b.Available != false {
		t.Errorf("Expected Available to remain false after second checkout attempt, got %v", b.Available)
	}
}

func TestBook_Return(t *testing.T) {
	b := NewBook("Brave New World", "Aldous Huxley", "978-0060850524", 1932)
	b.CheckOut()
	b.Return()
	if b.Available != true {
		t.Errorf("Expected Available to be true after return, got %v", b.Available)
	}
}

func TestBook_IsAvailable(t *testing.T) {
	b := NewBook("The Hobbit", "J.R.R. Tolkien", "978-0547928227", 1937)
	if !b.IsAvailable() {
		t.Errorf("Expected IsAvailable to be true initially, got %v", b.IsAvailable())
	}
	b.CheckOut()
	if b.IsAvailable() {
		t.Errorf("Expected IsAvailable to be false after checkout, got %v", b.IsAvailable())
	}
	b.Return()
	if !b.IsAvailable() {
		t.Errorf("Expected IsAvailable to be true after return, got %v", b.IsAvailable())
	}
}

// --- Lesson 03: Implementing Linked Lists and Waitlists ---

func TestBook_AddToWaitlist(t *testing.T) {
	b := NewBook("Dune", "Frank Herbert", "978-0441172719", 1965)
	patronID1 := 1
	patronID2 := 2
	b.AddToWaitlist(patronID1)
	if b.Waitlist.Size() != 1 {
		t.Errorf("Expected waitlist size 1, got %d", b.Waitlist.Size())
	}
	val, _ := b.Waitlist.Peek()
	if val != patronID1 {
		t.Errorf("Expected patron ID %d at the front of the waitlist, got %v", patronID1, val)
	}
	b.AddToWaitlist(patronID2)
	if b.Waitlist.Size() != 2 {
		t.Errorf("Expected waitlist size 2, got %d", b.Waitlist.Size())
	}
}

func TestBook_RemoveFromWaitlist(t *testing.T) {
	b := NewBook("Foundation", "Isaac Asimov", "978-0553293357", 1951)
	patronID1 := 1
	patronID2 := 2
	patronID3 := 3
	b.AddToWaitlist(patronID1)
	b.AddToWaitlist(patronID2)
	b.AddToWaitlist(patronID3)
	removedPatronID, err := b.RemoveFromWaitlist()
	if err != nil {
		t.Errorf("Unexpected error removing from waitlist: %v", err)
	}
	if removedPatronID != patronID1 {
		t.Errorf("Expected removed patron ID to be %d, got %d", patronID1, removedPatronID)
	}
	if b.Waitlist.Size() != 2 {
		t.Errorf("Expected waitlist size 2, got %d", b.Waitlist.Size())
	}
	b.RemoveFromWaitlist()
	b.RemoveFromWaitlist()
	_, err = b.RemoveFromWaitlist()
	if err == nil {
		t.Errorf("Expected error when removing from an empty waitlist, got nil")
	}
}

func TestBook_Waitlist_Integration(t *testing.T) {
	b := NewBook("Test Book", "Test Author", "123-456-7890", 2023)
	p1, p2, p3 := 1, 2, 3
	b.AddToWaitlist(p1)
	b.AddToWaitlist(p2)
	b.AddToWaitlist(p3)
	b.CheckOut()
	if b.Available {
		t.Fatal("Book should be unavailable after checkout")
	}
	returnedPatronID, err := b.ReturnWithWaitlist()
	if err != nil {
		t.Fatalf("ReturnWithWaitlist failed: %v", err)
	}
	if b.Available {
		t.Error("Book should not be available after return if there is a waitlist")
	}
	if returnedPatronID != p1 {
		t.Errorf("Expected patron ID %d to be notified, got %d", p1, returnedPatronID)
	}
	returnedPatronID, err = b.ReturnWithWaitlist()
	if err != nil {
		t.Fatalf("ReturnWithWaitlist failed: %v", err)
	}
	if returnedPatronID != p2 {
		t.Errorf("Expected patron ID %d to be notified, got %d", p2, returnedPatronID)
	}
	returnedPatronID, err = b.ReturnWithWaitlist()
	if err != nil {
		t.Fatalf("ReturnWithWaitlist failed: %v", err)
	}
	if returnedPatronID != p3 {
		t.Errorf("Expected patron ID %d to be notified, got %d", p3, returnedPatronID)
	}
	b.Return()
	if !b.Available {
		t.Error("Book should be available")
	}
	returnedPatronID, err = b.ReturnWithWaitlist()
	if err == nil {
		t.Error("Return should have failed with an error")
	}
	b2 := NewBook("Empty Waitlist Book", "Test Author", "111-222-3333", 2022)
	b2.CheckOut()
	returnedPatronID, err = b2.ReturnWithWaitlist()
	if returnedPatronID != 0 {
		t.Errorf("Expected patron ID 0 to be notified, got %d", returnedPatronID)
	}
	if !b2.IsAvailable() {
		t.Error("b2 Should be available now")
	}
}

// --- Lesson 11: Refactoring and UI Improvements ---

func TestBook_String(t *testing.T) {
	b := NewBook("Pride and Prejudice", "Jane Austen", "978-0141439518", 1813)
	expectedString := "Title: Pride and Prejudice\nAuthor: Jane Austen\nISBN: 978-0141439518\nPublication Year: 1813\nAvailable: true\nWaitlist: []"
	actualString := b.String()
	if !reflect.DeepEqual(actualString, expectedString) {
		t.Errorf("Expected string representation:\n%s\nGot:\n%s", expectedString, actualString)
	}
	b.CheckOut()
	expectedCheckedOut := "Title: Pride and Prejudice\nAuthor: Jane Austen\nISBN: 978-0141439518\nPublication Year: 1813\nAvailable: false\nWaitlist: []"
	actualCheckedOut := b.String()
	if !reflect.DeepEqual(actualCheckedOut, expectedCheckedOut) {
		t.Errorf("Expected string representation:\n%s\nGot:\n%s", expectedCheckedOut, actualCheckedOut)
	}
	b.AddToWaitlist(1)
	expectedWaitList := "Title: Pride and Prejudice\nAuthor: Jane Austen\nISBN: 978-0141439518\nPublication Year: 1813\nAvailable: false\nWaitlist: [1]"
	actualWaitList := b.String()
	if !reflect.DeepEqual(actualWaitList, expectedWaitList) {
		t.Errorf("Expected string representation:\n%s\nGot:\n%s", expectedWaitList, actualWaitList)
	}
}
