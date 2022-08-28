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

func textToAudio(message string) (string, error) {
	requestPayload := entity.SpeechRequest{
		AudioConfig: entity.AudioConfig{
			AudioEncoding: "LINEAR16",
			Pitch:         0,//7.2,
			SpeakingRate:  0.96,
		},
		Voice: entity.Voice{
			LanguageCode: "en-US",
			Name:         "en-US-Wavenet-H",
		},
		Input: entity.Input{
			Text: message,
		},
	}

	payload, err := json.Marshal(requestPayload)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://texttospeech.googleapis.com/v1beta1/text:synthesize?key=%s", os.Getenv("API_KEY"))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var m entity.SpeechResponse

	body, err := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(body, &m); err != nil {
		return "", err
	}

	return m.AudioContent, nil

}

func createAudioMessage(name string, identified string, found string) (string, error) {
	audioMessage := fmt.Sprintf("%s has thrown away a %s. It will go into the %s bin.", "Zhehai", identified, found)
	return textToAudio(audioMessage)
}
