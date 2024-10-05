package routes

import (
	"middlend-template/pkg/constants"
	"middlend-template/pkg/domain/innerDomain/response"
	"middlend-template/pkg/useCases/Handlers/apiHandlers/loginApiHandler"
	"middlend-template/pkg/useCases/Handlers/viewHandlers/templateViewHandler"
	"middlend-template/pkg/useCases/Helpers/responseHelper"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type TemplateRouter struct {
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	_, status := loginApiHandler.Authenticate(r.Header.Get("Authorization"), []int{constants.TEMPLATE_ROLE}, nil)
	if status != response.SuccessfulSearch {
		responseHelper.WriteResponse(w, response.Forbidden, nil)
		return
	}
	template := r.URL.Query().Get("template")
	body, status := templateViewHandler.TemplateFunction(template)

	responseHelper.WriteResponse(w, status, body)
}

func (rfr *TemplateRouter) Routes() http.Handler {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
		MaxAge:             300,
	}))

	r.Get("/", TemplateFunction)

	return r
}
