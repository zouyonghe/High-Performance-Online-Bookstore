package book

import "Jinshuzhai-Bookstore/handler"

type BaseResponse handler.BaseResponse

// swagger struct

// 创建书籍

// CreateBookRequest includes the book information
type CreateBookRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	PublishDate string  `json:"publish_date"`
	Category    string  `json:"category"`
	Number      uint64  `json:"number"`
	//ImagePath   string `json:"imagePath"`
}

// CreateBookResponse includes the bookId and the book title.
type CreateBookResponse struct {
	BookId string `json:"bookId"`
	Title  string `json:"title"`
}

// 删除书籍
// DeleteBookRequest includes the book title.
type DeleteBookRequest struct {
	Title string `json:"title"`
}

// 设置数量

// 设置上架

// 设置下架

// 设置作者
