package chatViewHandler

import (
	"agrobal-backend/pkg/constants"
	"agrobal-backend/pkg/domain/response"
	"agrobal-backend/pkg/useCases/Handlers/apiHandlers/geminiHandler"
)

func GetChatResponse(input string) (interface{}, response.Status) {
	/**
	  * This data is obtained from the constants package but must be obtained from an external API
	  * But for the purpose of this example, it is hardcoded
	**/
	weatherData := constants.WEATHER_DATA
	cropData := constants.CROP_DATA

	prompts := []string{
		"This is a question of a farmer, consider the following scenario: ",
		"The weather info is " + weatherData,
		"The question is: " + input,
		"Consider the following data: " + cropData,
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

	// Enhance the response
	prompts = []string{
		"These is a response for a farmer who asked about thir crop or farm",
		"Regenerate the response in a more human-friendly way",
		"Give the answer first, then the explanation",
		"Return the anser CONCISELY it in plain text, only paragraphs",
		"The answer is: " + resp,
	}

	// Request to Gemini
	contentResponse, status = geminiHandler.RequestGemini(prompts)
	if status != response.SuccessfulRequest {
		return nil, status
	}

	// Extract the response
	resp, status = geminiHandler.ExtractString(contentResponse)
	if status != response.SuccessfulParse {
		return nil, status
	}

	return resp, response.SuccessfulParse
}
