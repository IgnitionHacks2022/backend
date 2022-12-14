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
)

func ClassifyHandler(w http.ResponseWriter, r *http.Request) {
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

	found := "garbage"
	identified := "garbage"

	for index, obj := range m.Responses[0].LocalizedObjectAnnotations {
		if index == 0 {
			identified = obj.Name
		}

		names := strings.Split(string(obj.Name), " ")

		for _, s := range names {
			name := strings.ToLower(s)
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
	}

	for index, obj := range m.Responses[0].LabelAnnotations {
		if index == 0 && identified == "Garbage" {
			identified = obj.Description
		}

		names := strings.Split(string(obj.Description), " ")

		for _, s := range names {

			name := strings.ToLower(s)
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
	}

	name := "You"
	bluetoothID := "None"
	uID := uint(0)
	conn, err := db.Connection()
	for _, s := range requestBody.BluetoothIDs {
		fmt.Println(s)
		uID, err = db.GetUserId(conn, s)
		if err == nil {
			bluetoothID = s
			break
		}
	}
	fmt.Println(bluetoothID)
	if bluetoothID != "None" {
		name = db.GetUserName(conn, uID)
		itemRecord := db.Item{UserID: uID, Type: found, Name: identified}
		err = db.AddItem(conn, &itemRecord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(requestBody.BluetoothIDs, identified)
	}
	audioMessage := fmt.Sprintf("%s has thrown away a %s. It will go into the %s bin.", name, identified, found)
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

	w.Write(jsonResponse)
}
