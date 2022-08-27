package sorter

import (
	"backend/internal/entity"
	"backend/pkg/db"
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
	userId := vars["userId"]
	var requestBody entity.Classify
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
		http.Error(w, "Error marshaling json", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate?key=%s", os.Getenv("API_KEY"))

	//fmt.Println(url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Error calling google vision api", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	var m entity.GoogleResponse

	body, err := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(body, &m); err != nil {
		http.Error(w, "Error unpacking google response", http.StatusInternalServerError)
		return
	}

	// sb := string(body)

	var response entity.ClassifyResponse

	blueBinPossible, err := getBlueBin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading blue bin file")
		return
	}
	redBinPossible, err := getRedBin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading red bin file")
		return
	}

	found := "Garbage"
	identified := "Garbage"

	for _, s := range m.Responses[0].LocalizedObjectAnnotations {
		name := strings.ToLower(s.Name)
		blueExists := blueBinPossible[name]
		if blueExists {
			found = "blue"
			identified = name
			break
		}
		redExists := redBinPossible[name]
		if redExists {
			found = "red"
			identified = name
			break
		}
	}

	for _, s := range m.Responses[0].LabelAnnotations {
		name := strings.ToLower(s.Description)
		blueExists := blueBinPossible[name]
		if blueExists {
			found = "blue"
			identified = name
			break
		}
		redExists := redBinPossible[name]
		if redExists {
			found = "red"
			identified = name
			break
		}
	}

	audioMessage := fmt.Sprintf("%s has thrown away a %s. It will go into the %s bin.", "Zhehai", identified, found)
	audio, err := textToAudio(audioMessage)
	if err != nil {
		http.Error(w, "Error generating audio message", http.StatusInternalServerError)
	}
	response.Type = found
	response.Audio = audio

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	conn, err := db.Connection()
	uID, err := db.GetUserId(conn, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	itemRecord := db.Item{UserID: uID, Type: found, Name: identified}
	err = db.AddItem(conn, &itemRecord)
	if err != nil {
		http.Error(w, "Error adding item to db", http.StatusInternalServerError)
		return
	}
	fmt.Println(userId, identified)

	w.Write(jsonResponse)
}
