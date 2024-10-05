package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
	FORBIDDEN             = []byte("403: Forbidden")
	BAD_REQUEST           = []byte("400: Bad Request")
)

func New() http.Handler {
	r := chi.NewRouter()

	return r
}
