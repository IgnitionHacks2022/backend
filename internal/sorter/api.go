package sorter

import (
	"backend/internal/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func ClassifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["userId"])
	var requestBody entity.Classify
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Done decoding request body")

	imagePayload := entity.GoogleRequest{
		Requests: []entity.ImageType{
			entity.ImageType{
				Image: entity.ImageContent{Content: requestBody.Contents},
				Features: []entity.FeatureContent{
					entity.FeatureContent{
						Type:       "LABEL_DETECTION",
						MaxResults: 15,
					},
					entity.FeatureContent{
						Type:       "OBJECT_LOCALIZATION",
						MaxResults: 15,
					},
				},
			},
		},
	}

	payload, err := json.Marshal(imagePayload)

	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Println("Done making request body")

	url := fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate?key=%s", os.Getenv("API_KEY"))

	//fmt.Println(url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var m entity.GoogleResponse

	body, err := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(body, &m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sb := string(body)

	var response entity.ClassifyResponse

	blueBinPossible, err := getBlueBin()
	redBinPossible, err := getRedBin()

	found := "Garbage"

	for _, s := range m.Responses[0].LocalizedObjectAnnotations {
		name := strings.ToLower(s.Name)
		blueExists := blueBinPossible[name]
		if blueExists {
			found = "blue"
			break
		}
		redExists := redBinPossible[name]
		if redExists {
			found = "red"
			break
		}
	}

	for _, s := range m.Responses[0].LabelAnnotations {
		name := strings.ToLower(s.Description)
		blueExists := blueBinPossible[name]
		if blueExists {
			found = "blue"
			break
		}
		redExists := redBinPossible[name]
		if redExists {
			found = "red"
			break
		}
	}
	response.Type = found

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
