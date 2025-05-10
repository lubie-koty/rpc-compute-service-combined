package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func ValidateRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}
}

func GetRequestBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return body, err
	}
	if err := Validate.Struct(body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return body, err
	}
	return body, nil
}

func WriteResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetResponseBody[T any](response *http.Response) (T, error) {
	defer response.Body.Close()
	var body T
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return body, errors.Join(err, errors.New("Invalid response body"))
	}
	if err := Validate.Struct(body); err != nil {
		return body, errors.Join(err, errors.New("Invalid response body"))
	}
	return body, nil
}

func CreateRequest[T any](method string, url string, data T) (*http.Request, error) {
	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(data); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
