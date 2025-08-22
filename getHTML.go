package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return "", errors.New("can´t get url")
	}
	head := res.Header
	contentType := head.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", errors.New("unexpected content type: " + contentType)
	}
	byte, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}
