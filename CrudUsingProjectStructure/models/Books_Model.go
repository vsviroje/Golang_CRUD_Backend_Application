package models

type Book struct {
	BookID      int    `db:"BOOK_ID" json:"bookid,omitempty"`
	BookName    string `db:"BOOK_NAME" json:"bookname,omitempty"`
	BookDetails string `db:"BOOK_DETAILS" json:"bookdetails,omitempty"`
	BookAuthor  string `db:"BOOK_AUTHOR" json:"bookauthor,omitempty"`
	BookPrice   int    `db:"BOOK_PRICE" json:"bookprice"`
	BookRequest *int   `db:"BOOK_REQUEST" json:"bookrequest"`
}
