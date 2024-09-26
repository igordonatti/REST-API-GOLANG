package book

type BookController struct {
	Service *BookService
}

func NewBookController(service *BookService) *BookController {
	return &BookController{Service: service}
}
