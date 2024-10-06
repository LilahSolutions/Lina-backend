package routes

import (
	"agrobal-backend/pkg/domain/response"
	"agrobal-backend/pkg/useCases/Handlers/viewHandlers/chatViewHandler"
	"agrobal-backend/pkg/useCases/Helpers/responseHelper"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ChatRouter struct {
}

func (router *ChatRouter) GetChatResponse(w http.ResponseWriter, r *http.Request) {
	var inputStruct struct {
		Input string `json:"input"`
	}

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&inputStruct)
	if err != nil {
		responseHelper.WriteResponse(w, response.BadRequest, nil)
		return
	}

	// Get response
	resp, status := chatViewHandler.GetChatResponse(inputStruct.Input)
	if status != response.SuccessfulParse {
		responseHelper.WriteResponse(w, status, nil)
		return
	}

	// Write response
	responseHelper.WriteResponse(w, response.SuccessfulParse, resp)
}

func (router *ChatRouter) Routes() http.Handler {
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

	r.Post("/", router.GetChatResponse)

	return r
}
