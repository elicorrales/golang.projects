package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {

	return fmt.Sprintf("Title:%s, Author:%s, Date:%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID: 1, Title: "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams", YearPublished: 1979,
	},
	Book{
		ID: 2, Title: "The Hobbit",
		Author: "J.R.R. Tolkien", YearPublished: 1937,
	},
	Book{
		ID: 3, Title: "A Tale of Two Cities",
		Author: "Charles Dickens", YearPublished: 1859,
	},
	Book{
		ID: 4, Title: "Les Miserables",
		Author: "Victor Hugo", YearPublished: 1862,
	},
	Book{
		ID: 5, Title: "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling", YearPublished: 1997,
	},
	Book{
		ID: 6, Title: "I, Robot",
		Author: "Isaac Asimov", YearPublished: 1950,
	},
	Book{
		ID: 7, Title: "The Gods THemselves",
		Author: "Isaac Asimov", YearPublished: 1973,
	},
	Book{
		ID: 8, Title: "The Moon is a Harsh Mistress",
		Author: "Robert A. Heinlein", YearPublished: 1966,
	},
	Book{
		ID: 9, Title: "On Basilisk Station",
		Author: "David Weber", YearPublished: 1993,
	},
	Book{
		ID: 10, Title: "The Android's Dream",
		Author: "John Scalzi", YearPublished: 2006,
	},

	Book{
		ID: 11, Title: "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams", YearPublished: 1979,
	},
	Book{
		ID: 12, Title: "The Hobbit",
		Author: "J.R.R. Tolkien", YearPublished: 1937,
	},
	Book{
		ID: 13, Title: "A Tale of Two Cities",
		Author: "Charles Dickens", YearPublished: 1859,
	},
	Book{
		ID: 14, Title: "Les Miserables",
		Author: "Victor Hugo", YearPublished: 1862,
	},
	Book{
		ID: 15, Title: "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling", YearPublished: 1997,
	},
	Book{
		ID: 16, Title: "I, Robot",
		Author: "Isaac Asimov", YearPublished: 1950,
	},
	Book{
		ID: 17, Title: "The Gods THemselves",
		Author: "Isaac Asimov", YearPublished: 1973,
	},
	Book{
		ID: 18, Title: "The Moon is a Harsh Mistress",
		Author: "Robert A. Heinlein", YearPublished: 1966,
	},
	Book{
		ID: 19, Title: "On Basilisk Station",
		Author: "David Weber", YearPublished: 1993,
	},
	Book{
		ID: 20, Title: "The Android's Dream",
		Author: "John Scalzi", YearPublished: 2006,
	},

	Book{
		ID: 21, Title: "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams", YearPublished: 1979,
	},
	Book{
		ID: 22, Title: "The Hobbit",
		Author: "J.R.R. Tolkien", YearPublished: 1937,
	},
	Book{
		ID: 23, Title: "A Tale of Two Cities",
		Author: "Charles Dickens", YearPublished: 1859,
	},
	Book{
		ID: 24, Title: "Les Miserables",
		Author: "Victor Hugo", YearPublished: 1862,
	},
	Book{
		ID: 25, Title: "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling", YearPublished: 1997,
	},
	Book{
		ID: 26, Title: "I, Robot",
		Author: "Isaac Asimov", YearPublished: 1950,
	},
	Book{
		ID: 27, Title: "The Gods THemselves",
		Author: "Isaac Asimov", YearPublished: 1973,
	},
	Book{
		ID: 28, Title: "The Moon is a Harsh Mistress",
		Author: "Robert A. Heinlein", YearPublished: 1966,
	},
	Book{
		ID: 29, Title: "On Basilisk Station",
		Author: "David Weber", YearPublished: 1993,
	},
	Book{
		ID: 30, Title: "The Android's Dream",
		Author: "John Scalzi", YearPublished: 2006,
	},

	Book{
		ID: 31, Title: "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams", YearPublished: 1979,
	},
	Book{
		ID: 32, Title: "The Hobbit",
		Author: "J.R.R. Tolkien", YearPublished: 1937,
	},
	Book{
		ID: 33, Title: "A Tale of Two Cities",
		Author: "Charles Dickens", YearPublished: 1859,
	},
	Book{
		ID: 34, Title: "Les Miserables",
		Author: "Victor Hugo", YearPublished: 1862,
	},
	Book{
		ID: 35, Title: "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling", YearPublished: 1997,
	},
	Book{
		ID: 36, Title: "I, Robot",
		Author: "Isaac Asimov", YearPublished: 1950,
	},
	Book{
		ID: 37, Title: "The Gods THemselves",
		Author: "Isaac Asimov", YearPublished: 1973,
	},
	Book{
		ID: 38, Title: "The Moon is a Harsh Mistress",
		Author: "Robert A. Heinlein", YearPublished: 1966,
	},
	Book{
		ID: 39, Title: "On Basilisk Station",
		Author: "David Weber", YearPublished: 1993,
	},
	Book{
		ID: 40, Title: "The Android's Dream",
		Author: "John Scalzi", YearPublished: 2006,
	},
}
