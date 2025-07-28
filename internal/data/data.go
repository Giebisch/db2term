package data

import (
	"bytes"
	"db2term/internal/config"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func GetData(conf config.Config, chosen string) (*TripResponse, error) {
	abfahrt, ankunft := get_locations(conf, chosen)
	request := NewRequest(abfahrt, ankunft, 0)

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(request)
	if err != nil {
		return nil, err
	}

	url := "https://www.bahn.de/web/api/angebote/fahrplan"

	req, _ := http.NewRequest("POST", url, payloadBuf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		slog.Error("Could not make requst", "error", err)
		return nil, e
	}

	defer res.Body.Close()

	if res.StatusCode != 200 && res.StatusCode != 201 {
		slog.Error("Request unsuccessful", "status", res.Status)
		return nil, errors.New("TripRequest unsuccessful")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Could not read body", "error", err)
		return nil, err
	}

	response := TripResponse{}
	err = json.Unmarshal([]byte(body), &response)

	if err != nil {
		slog.Error("Unable to unmarshal JSON", "error", err)
	}

	return &response, nil
}

func get_locations(conf config.Config, chosen string) (string, string) {
	strings := strings.Split(chosen, " 󱦰 ")
	for _, trips := range conf.Trips {
		if trips.From == strings[0] && trips.To == strings[1] {
			return trips.FromCode, trips.ToCode
		}
	}
	panic("Could not get Codes")
}
