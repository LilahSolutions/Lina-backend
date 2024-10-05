package responseHelper

import (
	"encoding/json"
	"middlend-template/pkg/domain/innerDomain/response"
	"net/http"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil
}

func ResponseStatusChecker(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		return
	}
}

// Parse response data object to any struct (use struct as a pointer to parse content into it)
func ParseResponseStruct(v any, res *http.Response) response.Status {
	var auxResponse response.Response
	err := json.NewDecoder(res.Body).Decode(&auxResponse)
	if err != nil {
		return response.InternalServerError
	}
	parsedObj, err := json.Marshal(auxResponse.Data)
	if err != nil {
		return response.InternalServerError
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return response.Status{Text: auxResponse.Message, Code: res.StatusCode}
	}

	err = json.Unmarshal(parsedObj, &v)
	if err != nil {
		return response.InternalServerError
	}
	return response.SuccessfulParse
}

func WriteResponse(w http.ResponseWriter, status response.Status, data interface{}) {
	response := response.Response{
		Message: status.Text,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ResponseStatusChecker(w, []byte("500: Internal Server Error"))
		return
	}

	w.WriteHeader(status.Code)
	ResponseStatusChecker(w, marshalResponse)
}
