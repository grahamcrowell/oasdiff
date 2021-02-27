package diff

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// ServerDiff is a diff between server objects: https://swagger.io/specification/#server-object
type ServerDiff struct {
	ExtensionProps  *ExtensionsDiff `json:"extensions,omitempty"`
	URLDiff         *ValueDiff      `json:"urlType,omitempty"`
	DescriptionDiff *ValueDiff      `json:"description,omitempty"`
	// Variables
}

func (diff ServerDiff) empty() bool {
	return diff == ServerDiff{}
}

func getServerDiff(value1, value2 *openapi3.Server) ServerDiff {
	result := ServerDiff{}

	if diff := getExtensionsDiff(value1.ExtensionProps, value2.ExtensionProps); !diff.empty() {
		result.ExtensionProps = diff
	}

	result.URLDiff = getValueDiff(value1.URL, value2.URL)
	result.DescriptionDiff = getValueDiff(value1.Description, value2.Description)

	return result
}