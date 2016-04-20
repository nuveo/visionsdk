package visionsdk

import "google.golang.org/api/vision/v1"

var features []*vision.Feature

func appendFeature(l string, mr int64) {
	f := &vision.Feature{
		MaxResults: mr,
		Type:       l,
	}
	features = append(features, f)
}

func AddLabelDetection(maxResults int64) {
	appendFeature("LABEL_DETECTION", maxResults)
}

func AddTextDetection(maxResults int64) {
	appendFeature("TEXT_DETECTION", maxResults)
}

func AddFaceDetection(maxResults int64) {
	appendFeature("FACE_DETECTION", maxResults)
}

func AddLogoDetection(maxResults int64) {
	appendFeature("LOGO_DETECTION", maxResults)
}

func AddLandmarkDetection(maxResults int64) {
	appendFeature("LANDMARK_DETECTION", maxResults)
}

func AddImageProperties(maxResults int64) {
	appendFeature("IMAGE_PROPERTIES", maxResults)
}
