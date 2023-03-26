package book

type Service interface {
	GetAll() ([]Book, error)
	GetById(id int64) (Book, error)
	Create(book Book) (Book, error)
	Update(id int64, book Book) (Book, error)
	DeleteById(id int64) error

}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetById(id int64) (Book, error) {
	book, err := s.repo.FindById(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) Create(req Book) (Book, error) {
	book := Book{}

	id, err := s.repo.Insert(req)
	if err != nil {
		return book, err
	}

	book, err = s.repo.FindById(id)
	if err != nil {
		return book, err
	}

	return book, nil
}


func (s *service) Update(id int64, reqBook Book) (Book, error) {
	book, err := s.repo.FindById(id)
	if err != nil {
		return book, err
	}
	
	book.Title = reqBook.Title
	book.Author = reqBook.Author
	book.DescBook = reqBook.DescBook

	idx, err := s.repo.Update(book)
	if err != nil {
		return book, err
	}

	newBook, err := s.repo.FindById(idx)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) DeleteById(id int64) error {
	book, err := s.repo.DeleteById(id)
	if err != nil {
		return err
	}

	if book > 1{
		return nil
	}
	return nil
}