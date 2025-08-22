package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	r := strings.NewReader(htmlBody)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return walk(doc, base), nil
}

func walk(n *html.Node, base *url.URL) []string {
	var urls []string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				val := strings.TrimSpace(attr.Val)
				if val == "" {
					continue
				}
				u, err := url.Parse(val)
				if err != nil {
					continue // ignorar href inválido
				}
				abs := base.ResolveReference(u)
				urls = append(urls, abs.String())
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childURLs := walk(c, base)
		urls = append(urls, childURLs...)
	}

	return urls
}
