package templateApiHandler

import "middlend-template/pkg/domain/innerDomain/response"

func TemplateApiGet(templateString string) (interface{}, response.Status) {
	return templateString, response.SuccessfulSearch
}
