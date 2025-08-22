package main

import (
	"errors"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	if res.StatusCode >= 400 {
		return "", errors.New("can´t get url")
	}
	head := res.Header
	if head.Get("content-type") != "text/html" {
		return "", errors.New("not text/html")
	}
	byte, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}
