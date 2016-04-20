package visionsdk

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"
)

func newVisionService() (*vision.Service, error) {
	ctx := context.TODO()
	client, err := google.DefaultClient(ctx, vision.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	srv, err := vision.New(client)
	if err != nil {
		return nil, err
	}

	return srv, nil

}

func Parse(filePath string) ([]byte, error) {
	//Initialing VisionService
	srv, err := newVisionService()

	if err != nil {
		return nil, err
	}

	//Create request
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	req := &vision.AnnotateImageRequest{
		Image: &vision.Image{
			Content: base64.StdEncoding.EncodeToString(b),
		},
		Features: features,
	}

	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	res, err := srv.Images.Annotate(batch).Do()
	if err != nil {
		return nil, err
	}

	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		return nil, err
	}

	return body, nil
}
