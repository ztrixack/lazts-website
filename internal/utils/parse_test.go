package utils

import (
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		setup         func() error
		teardown      func()
		expectedErr   bool
		expectedFiles int
	}{
		{
			name: "Successful template parsing",
			path: "*.tmpl",
			setup: func() error {
				file1, err := os.Create("test1.tmpl")
				if err != nil {
					return err
				}
				file1.Close()

				file2, err := os.Create("test2.tmpl")
				if err != nil {
					return err
				}
				file2.Close()

				return nil
			},
			teardown: func() {
				os.RemoveAll("test1.tmpl")
				os.RemoveAll("test2.tmpl")
			},
			expectedErr:   false,
			expectedFiles: 2,
		},
		{
			name: "Error in file path globbing",
			path: "[invalid pattern",
			setup: func() error {
				return nil
			},
			teardown: func() {
			},
			expectedErr:   true,
			expectedFiles: 0,
		},
		{
			name: "Error in template parsing",
			path: "*.tmpl",
			setup: func() error {
				file, err := os.Create("test.tmpl")
				if err != nil {
					return err
				}
				file.WriteString("{{ define }}") // Malformed template
				file.Close()
				return nil
			},
			teardown: func() {
				os.RemoveAll("test.tmpl")
			},
			expectedErr:   true,
			expectedFiles: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setup()
			defer tt.teardown()
			assert.NoError(t, err, "failed to setup test")

			tmpl := template.New("test")
			files, err := ParseAnyTemplates(tmpl, tt.path)
			if tt.expectedErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during parsing")
			}

			assert.Len(t, files, tt.expectedFiles, "unexpected number of files parsed")
		})
	}
}
