package book

type BookResponse struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type BookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
}

type GetBookDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}

func ToBook(book Book) BookResponse {
	res := BookResponse{}
	res.ID = book.ID
	res.Title = book.Title
	res.Author = book.Author
	res.Desc = book.DescBook

	return res
}

func ToBooks(books []Book) []BookResponse {
	bookList := []BookResponse{}

	for _, book := range books {
		toBook := ToBook(book)
		bookList = append(bookList, toBook)
	}

	return bookList
}

func ToEntity(req BookRequest) Book {
	book := Book{}
	book.Title = req.Title
	book.Author = req.Author
	book.DescBook = req.Desc

	return book
}
