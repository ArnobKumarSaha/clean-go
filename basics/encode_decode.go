package basics

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, welcome to %s!", r.URL.Path[1:])
	_, err := decodeRequest(w, r)
	if err != nil {
		return
	}
}

func EncodeDecodeLearn() {
	//stringEncodeDecode()

	// To illustrate the examples from https://medium.com/what-i-talk-about-when-i-talk-about-technology/go-code-snippet-json-encoder-and-json-decoder-818f81864614
	// it is necessary to run a go server first.
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func stringEncodeDecode() {
	str := "https://www.sourcecodeexamples.net/p/golang-source-code-examples.html"
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(enc)

	dec, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return
	}
	fmt.Println(string(dec))
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// these two functions from server side.

func decodeRequest(w http.ResponseWriter, r *http.Request) (*UserRequest, error) {
	var req UserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return nil, err
	}
	return &req, nil
}

func encodeResponse(w http.ResponseWriter, req UserRequest) {
	resp := UserResponse{
		ID:    123,
		Name:  req.Name,
		Email: req.Email,
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// these two from client side

func encodeRequest(url string, user UserRequest) (*http.Response, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(user)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func decodeResponse(resp *http.Response) (*UserResponse, error) {
	var userResp *UserResponse
	err := json.NewDecoder(resp.Body).Decode(userResp)
	if err != nil {
		return nil, err
	}
	return userResp, nil
}
