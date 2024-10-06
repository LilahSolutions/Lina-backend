package geminiHandler

import (
	"agrobal-backend/pkg/constants"
	"agrobal-backend/pkg/domain/response"
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ExtractString(contentResponse *genai.GenerateContentResponse) (string, response.Status) {
	if len(contentResponse.Candidates) == 0 {
		return "", response.ParseFailed
	}

	candidate := contentResponse.Candidates[0]
	if len(candidate.Content.Parts) == 0 {
		return "", response.ParseFailed
	}

	resp := fmt.Sprintln(candidate.Content.Parts[0])

	// Remove the last character if it is a newline
	resp = strings.TrimSuffix(resp, "\n")

	return resp, response.SuccessfulParse

}

func RequestGemini(promptsString []string) (*genai.GenerateContentResponse, response.Status) {
	ctx := context.Background()

	// Create gemini client
	geminiClient, err := genai.NewClient(ctx, option.WithAPIKey(constants.GEMINI_API_KEY))
	if err != nil {
		fmt.Print("Error creating Gemini client: \n", err)
		return nil, response.InternalServerError
	}
	defer geminiClient.Close()

	// Set the model and temperature
	model := geminiClient.GenerativeModel(constants.GEMINI_MODEL)
	model.SetTemperature(float32(constants.GEMINI_TEMPERATURE))

	// Set the safety settings to not block any content
	model.SafetySettings = []*genai.SafetySetting{
		{Category: genai.HarmCategoryDangerousContent, Threshold: genai.HarmBlockNone},
		{Category: genai.HarmCategoryHarassment, Threshold: genai.HarmBlockNone},
		{Category: genai.HarmCategorySexuallyExplicit, Threshold: genai.HarmBlockNone},
		{Category: genai.HarmCategoryHateSpeech, Threshold: genai.HarmBlockNone},
	}

	// Set the prompt
	prompt := []genai.Part{}
	for _, promptString := range promptsString {
		prompt = append(prompt, genai.Text(promptString))
	}

	contentResponse, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		fmt.Print("Error generating content: \n", err)
		return nil, response.FailedRequest
	}

	return contentResponse, response.SuccessfulRequest
}
