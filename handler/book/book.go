package book

import "Jinshuzhai-Bookstore/model"

//swagger struct

// AddRequest represents a request for adding a book.
type AddRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	PublishDate string  `json:"publishDate"`
	Category    string  `json:"category"`
	IsSell      bool    `json:"isSell"`
	Number      uint64  `json:"number"`
}

type AddResponse struct {
	BookID uint64 `json:"bookId"`
	Title  string `json:"title"`
}

type ListRequest struct {
	Title    string `json:"title"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	BookList   []*model.BookInfo `json:"bookList"`
}
