package web

import "fmt"

var (
	ErrNotFound       = fmt.Errorf("template not found")
	ErrCloneTemplates = fmt.Errorf("failed to clone templates")
	ErrParseTemplate  = fmt.Errorf("failed to parse template")
	ErrParseContent   = fmt.Errorf("failed to parse content")
)
