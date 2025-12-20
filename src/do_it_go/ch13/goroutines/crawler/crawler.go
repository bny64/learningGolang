package main

import (
	"io"
	"net/http"
	"sync"
)

const UserAgent = "Mozilla/5.0"

func Crawler(index NewsIndex, useGoroutine bool) error {
	var wg sync.WaitGroup
	for i, link := range index.Link {
		wg.Add(1)
		if useGoroutine {
			go GetContent(link.URI, &index.Link[i], &wg)
		} else {
			err := GetContent(link.URI, &index.Link[i], &wg)
			if err != nil {
				return err
			}
		}
	}
	wg.Wait()
	return nil
}

func GetContent(url string, link *NewsLink, wg *sync.WaitGroup) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Add("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		wg.Done()
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		wg.Done()
		return err
	}

	link.Content = string(body)
	wg.Done()
	return nil
}
