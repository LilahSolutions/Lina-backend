package templateViewHandler

import (
	"middlend-template/pkg/domain/innerDomain/response"
	"middlend-template/pkg/useCases/Handlers/apiHandlers/templateApiHandler"
)

func TemplateFunction(templateString string) (interface{}, response.Status) {
	return templateApiHandler.TemplateApiGet(templateString)
}
