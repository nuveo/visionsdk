package visionsdk

import "testing"

func TestFeatures(t *testing.T) {
	vs, _ := New()
	if len(vs.features) != 0 {
		t.Errorf("Expected 0 features. Got %d", len(vs.features))
	}

	vs.AddFaceDetection(2)
	if len(vs.features) != 1 {
		t.Errorf("Expected 1 features. Got %d", len(vs.features))
	}
	if vs.features[0].MaxResults != 2 {
		t.Errorf("Expected 2 MAX_RESULTS. Got %d", vs.features[0].MaxResults)
	}

	if vs.features[0].Type != "FACE_DETECTION" {
		t.Errorf("Expected 'FACE_DETECTION' Type. Got %s", vs.features[0].Type)

	}

	vs.AddTextDetection(10)
	if len(vs.features) != 2 {
		t.Errorf("Expected 2 features. Got %d", len(vs.features))
	}
	if vs.features[1].MaxResults != 10 {
		t.Errorf("Expected 10 MAX_RESULTS. Got %d", vs.features[1].MaxResults)
	}

	if vs.features[1].Type != "TEXT_DETECTION" {
		t.Errorf("Expected 'TEXT_DETECTION' Type. Got %s", vs.features[1].Type)

	}

	vs.AddLabelDetection(1)
	if len(vs.features) != 3 {
		t.Errorf("Expected 3 features. Got %d", len(vs.features))
	}
	if vs.features[2].MaxResults != 1 {
		t.Errorf("Expected 1 MAX_RESULTS. Got %d", vs.features[2].MaxResults)
	}

	if vs.features[2].Type != "LABEL_DETECTION" {
		t.Errorf("Expected 'LABEL_DETECTION' Type. Got %s", vs.features[2].Type)
	}
}

func TestVisionSDK(t *testing.T) {
	vs, _ := New()
	vs.AddLabelDetection(2)

	_, err := vs.Parse("./tests/funny_lazy_cat-wallpaper-1280x1024.jpg")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}
