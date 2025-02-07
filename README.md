# Alexandria: A Library Management System

[![Go Report Card](https://goreportcard.com/badge/github.com/gaahrdner/alexandria)](https://goreportcard.com/report/github.com/gaahrdner/alexandria)
[![GoDoc](https://godoc.org/github.com/gaahrdner/alexandria?status.svg)](https://godoc.org/github.com/gaahrdner/alexandria)

**The year is 2347.**  The world is recovering from "The Great Forgetting," a catastrophic digital collapse that wiped out most of humanity's digital knowledge.  The **Phoenix Initiative** is a global effort to rebuild, and at its heart is the **Alexandria** project – a digital library designed to be resilient, secure, and accessible to all.

**You** are a skilled software engineer joining the Phoenix Initiative. You'll be working alongside **Anya Sharma**, the Lead Architect, to build the core of the Alexandria.  You'll collaborate with a team of specialists, each with their own unique skills and motivations, to overcome technical challenges and ensure the library's success.  Your contributions will directly impact the future of knowledge preservation.

## Motivation

This project was born out of a desire for a focused, hands-on learning experience. Faced with a couple of long plane rides and a yearning to solidify my understanding of fundamental data structures and algorithms, I decided to create a project that would be both challenging and engaging. The idea of building a digital library, a repository of knowledge, resonated with me, and the narrative framework of the Phoenix Initiative added an extra layer of motivation. This project is a personal learning journey, and I hope it can be a useful resource for others as well.

Also AI wrote pretty much all of this, I merely guided.  Specifically Gemini 2.0 Pro Experimental 02-05. I sure hope it's correct!

## The Team

* **Dr. Elias Vance:** The *Visionary*.  A brilliant but occasionally eccentric historian in his late 60s, Elias is the driving force. He lost his own life's work in the Great Forgetting – a digital archive of ancient languages – and is driven by a fierce determination to prevent such loss from ever happening again. He sees the Alexandria as a continuation of his life's work, a way to rebuild what was lost. He is Lillian Vance's grandfather, a fact that brings an added dimension of mentorship and legacy to the project.

* **Anya Sharma:** The *Lead Architect*. A young (28) and fiercely talented software engineer, Anya is responsible for the system's design.  She lost her family in the chaos following the Great Forgetting and sees the Alexandria as a way to honor their memory and build a more resilient future.  The pressure of leading such a critical project weighs heavily on her, but she faces it with unwavering determination.

* **Kaito "Kai" Tanaka:** The *Performance Guru*.  A 35-year-old systems engineer obsessed with optimization.  Kai survived the Great Forgetting by living off-grid, a self-reliant "digital hermit." He's skeptical of large-scale systems, having seen their vulnerabilities firsthand, but he believes in the Alexandria's mission. He's also a distant relative of the Mr. Ito from the earlier versions, carrying on a family legacy of technical expertise. He is often blunt, but his insights are invaluable.

* **Isabelle Dubois:** The *Data Recovery Specialist*. A 40-year-old former intelligence analyst, Isabelle specializes in retrieving data from damaged and corrupted storage media.  She's cool, calm, and incredibly resourceful, able to piece together fragments of information like a digital detective.  She carries the guilt of her past work, which involved surveillance and data manipulation, and sees the Alexandria as a way to use her skills for good.

* **Marcus Cole:** The *Security Expert*. A hardened, 50-year-old cybersecurity specialist who worked for a major corporation before the collapse. He witnessed the cyberattack that contributed to the Great Forgetting and is deeply cynical about technology's unchecked power.  He's obsessed with preventing another attack, sometimes to the point of paranoia, but his vigilance is essential.

The story begins in `docs/01_the_genesis_scroll.md`

## Features (Implemented Incrementally)

* **Core Functionality:**
  * **Book Management:**
    * Add new books (title, author, ISBN, publication year).
    * List all books.
    * Mark books as available or checked out.
    * Manage waitlists for books (using a linked list).
  * **Patron Management:**
    * Add new patrons (name and ID).
    * List all patrons.
    * Track books checked out by each patron.
  * **Checkout/Return:**
    * Check out books to patrons.
    * Return books from patrons.
    * Automatic waitlist handling: notify the next patron in line.
* **Advanced Features:**
  * **Undo/Redo:**
    * Basic undo/redo functionality (add book, add patron, checkout, return) using a stack.
  * **Efficient Searching:**
    * Search for patrons by ID (using a hash table).
    * Search for books by ISBN (using a binary search tree).
    * Search for books by Title (using linear search for now, and other optimal algorithms in the future).
  * **Sorting**:
    * List all books, sorted by title, author, ISBN, or publication year (configurable sorting criteria using various algorithms).
  * **Similar Books**
    * Add relationships between books.
    * List similar books using graph traversal (using a graph).

## Data Structures and Algorithms Demonstrated

* Slices
* Linked List
* Stack
* Queue
* Hash Table
* Binary Search Tree (BST)
* Graph
* Sorting Algorithms (Bubble, Selection, Insertion, Merge, Quick, Heap)

## Prerequisites

* Go (version 1.23 or later recommended, but it might work with earlier versions, I don't know.)

## Project Structure

```shell
alexandria/
├── cmd/
│   └── alexandria/
│       └── main.go  # Main application entry point
├── internal/
│   ├── book/
│   │   ├── book.go      # Book data model and related functions
│   │   └── book_test.go # Tests for the Book struct and methods
│   ├── patron/
│   │   ├── patron.go      # Patron data model and related functions
│   │   └── patron_test.go # Tests for the Patron struct and methods
│   ├── sorting/         # Sorting algorithms
│   │   └── sorting_test.go # Tests for all sorting algorithms
│   ├── linkedlist/      # Linked list implementation
│   │   └── linkedlist_test.go
│   ├── stackqueue/      # Stack and Queue implementations
│   │   ├── queue_test.go
│   │   └── stack_test.go
│   ├── bst/             # Binary Search Tree implementation
│   │   └── bst_test.go
│   ├── hashtable/       # Hash Table implementation
│   │   └── hashtable_test.go
│   └── graph/
│       ├── graph.go # Graph implementation
│       └── graph_test.go
├── go.mod  # Go module definition
└── go.sum  # Go module checksums
```

## Building and Running

### Clone the repository

```bash
git clone https://github.com/gaahrdner/alexandria.git
```

### Build the application

```bash
go build ./cmd/alexandria
```

### Run the application

```bash
./alexandria
```

## Testing

### Run all unit tests

```bash
go test ./...
```

### Run tests for a specific package (e.g., book)

```bash
go test ./internal/book
```

### Run benchmarks (for performance analysis)

```bash
go test -bench=. -benchmem ./...
```

Add -run=^# to avoid running non-benchmark tests. Example:
bash go test -bench=. -benchmem -run=^# ./internal/sorting

## Contributing

Contributions are welcome! Please follow these steps:

Fork the repository.

Create a new branch: `git checkout -b my-feature`

Make your changes, and write unit tests.

Run the tests: go test ./...

Commit your changes: git commit -m "Add my feature"

Push your branch: git push origin my-feature

Open a pull request.

License

This project is licensed under the MIT License - see the LICENSE file for details.
