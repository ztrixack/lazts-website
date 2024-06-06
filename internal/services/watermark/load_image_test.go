package watermark

import (
	"errors"
	"image"
	"lazts/internal/modules/imaging"
	"lazts/internal/utils"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoadImageTableDriven(t *testing.T) {
	watermark := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	original := image.NewNRGBA(image.Rect(0, 0, 800, 600))
	processed := image.NewNRGBA(image.Rect(0, 0, 800, 600))

	utils.CreateTestFile(t, "path/to", "original", "some content")
	utils.CreateTestFile(t, "path/to", "watermark", "some content")
	defer utils.RemoveTestDir(t, "path")

	testcases := []struct {
		name        string
		filePath    string
		setup       func(*service, *imaging.Mock)
		expectedImg image.Image
		expectError bool
	}{
		{
			name:     "Successful processing",
			filePath: "path/to/original",
			setup: func(s *service, m *imaging.Mock) {

				m.On("Open", "path/to/watermark").Return(watermark, nil)
				m.On("Resize", watermark, 100, 0).Return(watermark)
				m.On("Open", "path/to/original").Return(original, nil)
				m.On("Overlay", original, watermark, mock.AnythingOfType("image.Point"), 1.0).Return(processed)
			},
			expectedImg: processed,
			expectError: false,
		},
		{
			name:     "Successful processing with cache",
			filePath: "path/to/original",
			setup: func(s *service, m *imaging.Mock) {
				s.cache["path/to/original"] = processed
			},
			expectedImg: processed,
			expectError: false,
		},
		{
			name:     "Watermark file not found error",
			filePath: "path/to/original",
			setup: func(s *service, m *imaging.Mock) {
				m.On("Open", "path/to/watermark").Return(nil, errors.New("file not found"))
			},
			expectedImg: nil,
			expectError: true,
		},
		{
			name:     "Image file not found error",
			filePath: "path/to/original",
			setup: func(s *service, m *imaging.Mock) {
				m.On("Open", "path/to/watermark").Return(watermark, nil)
				m.On("Resize", watermark, 100, 0).Return(watermark)
				m.On("Open", "path/to/original").Return(nil, errors.New("file not found"))
			},
			expectedImg: nil,
			expectError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockImaging := new(imaging.Mock)
			service := &service{
				config: &config{
					Path: "path/to/watermark",
					Size: 100,
				},
				imager: mockImaging,
				cache:  make(map[string]image.Image),
				mutex:  &sync.Mutex{},
			}

			tc.setup(service, mockImaging)

			result, err := service.LoadImage(tc.filePath)

			if tc.expectError {
				assert.Error(t, err, "Expected an error for test case: "+tc.name)
			} else {
				assert.NoError(t, err, "Did not expect error for test case: "+tc.name)
				assert.Equal(t, tc.expectedImg, result, "Expected image did not match for test case: "+tc.name)
			}

			mockImaging.AssertExpectations(t)
		})
	}
}
