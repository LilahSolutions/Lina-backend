package chatViewHandler

import (
	"agrobal-backend/pkg/constants"
	"agrobal-backend/pkg/domain/response"
	"agrobal-backend/pkg/useCases/Handlers/apiHandlers/geminiHandler"
)

func GetChatResponse(input string) (interface{}, response.Status) {
	prompts := []string{
		"This is a question of a farmer, consider the following scenario: ",
		"The weather info is " + constants.WEATHER_DATA,
		"The user data is " + constants.USER_DATA,
		"The question is: " + input,
		"Consider the following data: " + constants.CROP_DATA,
		"Return the anser CONCISELY it in plain text",
		"Id input is oustide the scope of the data, return 'I am sorry, I cannot answer that question'",
	}

	// Request to Gemini
	contentResponse, status := geminiHandler.RequestGemini(prompts)
	if status != response.SuccessfulRequest {
		return nil, status
	}

	// Extract the response
	resp, status := geminiHandler.ExtractString(contentResponse)
	if status != response.SuccessfulParse {
		return nil, status
	}

	return resp, response.SuccessfulParse
}
