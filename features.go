package visionsdk

import "google.golang.org/api/vision/v1"

func (vo *VisionObject) appendFeature(l string, mr int64) {
	f := &vision.Feature{
		MaxResults: mr,
		Type:       l,
	}
	vo.features = append(vo.features, f)
}

func (vo *VisionObject) AddLabelDetection(maxResults int64) {
	vo.appendFeature("LABEL_DETECTION", maxResults)
}

func (vo *VisionObject) AddTextDetection(maxResults int64) {
	vo.appendFeature("TEXT_DETECTION", maxResults)
}

func (vo *VisionObject) AddFaceDetection(maxResults int64) {
	vo.appendFeature("FACE_DETECTION", maxResults)
}

func (vo *VisionObject) AddLogoDetection(maxResults int64) {
	vo.appendFeature("LOGO_DETECTION", maxResults)
}

func (vo *VisionObject) AddLandmarkDetection(maxResults int64) {
	vo.appendFeature("LANDMARK_DETECTION", maxResults)
}

func (vo *VisionObject) AddImageProperties(maxResults int64) {
	vo.appendFeature("IMAGE_PROPERTIES", maxResults)
}
