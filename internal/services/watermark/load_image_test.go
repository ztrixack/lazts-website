package watermark

import (
	"errors"
	"image"
	"lazts/internal/modules/imaging"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoadImage(t *testing.T) {
	const WEB_DIR = "test_load_image"

	watermark := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	original := image.NewNRGBA(image.Rect(0, 0, 800, 600))
	processed := image.NewNRGBA(image.Rect(0, 0, 800, 600))

	tests := []struct {
		name           string
		filename       string
		expectedError  bool
		expectedResult image.Image
		setup          func(*service, *imaging.Mock)
	}{
		{
			name:     "Successful processing",
			filename: "original.png",
			setup: func(s *service, m *imaging.Mock) {
				m.On("Open", "test_load_image/watermark.png").Return(watermark, nil)
				m.On("Resize", watermark, 32, 0).Return(watermark)
				m.On("Open", "test_load_image/original.png").Return(original, nil)
				m.On("Overlay", original, watermark, mock.AnythingOfType("image.Point"), 1.0).Return(processed)
			},
			expectedResult: processed,
			expectedError:  false,
		},
		{
			name:     "Watermark file not found error",
			filename: "original.png",
			setup: func(s *service, m *imaging.Mock) {
				m.On("Open", "test_load_image/watermark.png").Return(nil, errors.New("file not found"))
			},
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:     "Image file not found error",
			filename: "original.png",
			setup: func(s *service, m *imaging.Mock) {
				m.On("Open", "test_load_image/watermark.png").Return(watermark, nil)
				m.On("Resize", watermark, 32, 0).Return(watermark)
				m.On("Open", "test_load_image/original.png").Return(nil, errors.New("file not found"))
			},
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, m := setup(t, WEB_DIR)
			defer teardown(t, WEB_DIR)
			tt.setup(s, m)

			result, err := s.LoadImage(tt.filename)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error for test case: "+tt.name)
			} else {
				assert.NoError(t, err, "Did not expect error for test case: "+tt.name)
				assert.Equal(t, tt.expectedResult, result, "Expected image did not match for test case: "+tt.name)
			}

			m.AssertExpectations(t)
		})
	}
}
