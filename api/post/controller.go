package post

type PostController struct {
	Service *PostService
}

func NewPostController(service *PostService) *PostController {
	return &PostController{Service: service}
}
