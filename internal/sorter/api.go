package sorter

import (
	"backend/internal/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ImageContent struct {
	Content string
}

type ImageType struct {
	Image ImageContent
}

type GoogleRequest struct {
	Requests []ImageType
}

func ClassifyHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.Classify
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imagePayload := GoogleRequest{
		Requests: []ImageType{
			ImageType{Image: ImageContent{Content: requestBody.Contents}},
		},
	}

	payload, err := json.Marshal(imagePayload)

	if err != nil {
		fmt.Println("json error")
	}

	resp, err := http.Post(fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate/%s", os.Getenv("API_KEY")), "application/json", bytes.NewBuffer(payload))

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error parsing body")
	}

	sb := string(body)

	fmt.Println(sb)
}
