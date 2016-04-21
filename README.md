# Vision SDK
Google Vision Go(lang) SDK

## Quickstart Tutorial about Google Vision API

https://cloud.google.com/vision/docs/getting-started

## How to use nuveo Vision SDK

### Installing visionsdk package.

```
$ go get github.com/nuveo/visionsdk
```

Create a account json file.

https://cloud.google.com/vision/docs/auth-template/cloud-api-auth

Export the account json file.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=<path_to_service_account_file>
```

### Testing if everything is ok.

```
$ cd $GOPATH/src/github.com/nuveo/visionsdk
$ go test -v
```

### Example

```go
package main

import (
    "os"
	"log"

	"github.com/nuveo/visionsdk"
)

func main() {
    vs, _ := visionsdk.New()
	//Set vision feature type and max_results
	vs.AddLabelDetection(10)
	vs.AddTextDetection(2)

	path := os.Getenv("GOPATH") + "/src/github.com/nuveo/visionsdk/tests/funny_lazy_cat-wallpaper-1280x1024.jpg"
	res, err := vs.Parse(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(res))
}
```

Returns:

```json
[
  {
    "labelAnnotations": [
      {
        "description": "pet",
        "mid": "/m/068hy",
        "score": 0.93436629
      },
      {
        "description": "cat",
        "mid": "/m/01yrx",
        "score": 0.91760761
      },
      {
        "description": "animal",
        "mid": "/m/0jbk",
        "score": 0.8943547
      },
      {
        "description": "mammal",
        "mid": "/m/04rky",
        "score": 0.87561196
      },
      {
        "description": "kitten",
        "mid": "/m/0hjzp",
        "score": 0.82297045
      },
      {
        "description": "play",
        "mid": "/m/06wtgq",
        "score": 0.81237704
      },
      {
        "description": "close up",
        "mid": "/m/02cqfm",
        "score": 0.7672019
      },
      {
        "description": "small to medium sized cats",
        "mid": "/m/07k6w8",
        "score": 0.72727275
      },
      {
        "description": "cat like mammal",
        "mid": "/m/0307l",
        "score": 0.66666669
      },
      {
        "description": "carnivoran",
        "mid": "/m/01lrl",
        "score": 0.60204762
      }
    ]
  }
]
```
