package book

type Service interface {
	GetAll() (*[]Book, error)
	GetById(id int) (*Book, error)
	Create(book BookRequest) (*Book, error)
	Update(id int, book BookRequest) (*Book, error)
	DeleteById(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAll() (*[]Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetById(id int) (*Book, error) {
	book, err := s.repo.FindById(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) Create(bookReq BookRequest) (*Book, error) {
	var book = Book{}
	book.Title = bookReq.Title
	book.Author = bookReq.Author

	bookCreated, err := s.repo.Insert(&book)
	if err != nil {
		return bookCreated, err
	}

	return bookCreated, nil
}

func (s *service) Update(id int, bookReq BookRequest) (*Book, error) {
	var book = Book{}
	book.Title = bookReq.Title
	book.Author = bookReq.Author

	updatedBook, err := s.repo.Update(id, &book)
	if err != nil {
		return updatedBook, err
	}

	return updatedBook, nil
}

func (s *service) DeleteById(id int) error {
	err := s.repo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func ToBook(book Book) BookResponse {
	res := BookResponse{}
	res.ID = book.ID
	res.Title = book.Title
	res.Author = book.Author
	res.CreatedAt = book.CreatedAt
	res.UpdatedAt = book.UpdatedAt

	return res
}

func ToBooks(books []Book) []BookResponse {
	listOfBook := []BookResponse{}

	for _, book := range books {
		toBook := ToBook(book)
		listOfBook = append(listOfBook, toBook)
	}

	return listOfBook
}
