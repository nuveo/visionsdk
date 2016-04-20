package visionsdk

import (
	"testing"

	"google.golang.org/api/vision/v1"
)

func TestFeatures(t *testing.T) {
	if len(features) != 0 {
		t.Errorf("Expected 0 features. Got %d", len(features))
	}

	AddFaceDetection(2)
	if len(features) != 1 {
		t.Errorf("Expected 1 features. Got %d", len(features))
	}
	if features[0].MaxResults != 2 {
		t.Errorf("Expected 2 MAX_RESULTS. Got %d", features[0].MaxResults)
	}

	if features[0].Type != "FACE_DETECTION" {
		t.Errorf("Expected 'FACE_DETECTION' Type. Got %s", features[0].Type)

	}

	AddTextDetection(10)
	if len(features) != 2 {
		t.Errorf("Expected 2 features. Got %d", len(features))
	}
	if features[1].MaxResults != 10 {
		t.Errorf("Expected 10 MAX_RESULTS. Got %d", features[1].MaxResults)
	}

	if features[1].Type != "TEXT_DETECTION" {
		t.Errorf("Expected 'TEXT_DETECTION' Type. Got %s", features[1].Type)

	}

	AddLabelDetection(1)
	if len(features) != 3 {
		t.Errorf("Expected 3 features. Got %d", len(features))
	}
	if features[2].MaxResults != 1 {
		t.Errorf("Expected 1 MAX_RESULTS. Got %d", features[2].MaxResults)
	}

	if features[2].Type != "LABEL_DETECTION" {
		t.Errorf("Expected 'LABEL_DETECTION' Type. Got %s", features[2].Type)
	}
}

func TestVisionSDK(t *testing.T) {
	features = []*vision.Feature{}
	AddLabelDetection(2)
	_, err := Parse("./tests/funny_lazy_cat-wallpaper-1280x1024.jpg")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}
