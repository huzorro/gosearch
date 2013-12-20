package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Searcher interface {
	Search(r *http.Request, w http.ResponseWriter) (int, string)
}

const (
	BING_KEY string = `BbqNHTFiVOCkt1myItsU8STfPErvGyy/UYgmQKicwmU`
	BING_URL string = `https://api.datamarket.azure.com/Bing/Search/v1`
)

type BingSearcher struct{}

func (_ BingSearcher) Search(r *http.Request, w http.ResponseWriter) (int, string) {
	u := fmt.Sprintf(`%s/Web?$format=json&Query=%%27%s%%27`, BING_URL, url.QueryEscape(r.URL.Query().Get("Query")))
	client := &http.Client{}

	request, err := http.NewRequest("GET", u, nil)
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(BING_KEY+":"+BING_KEY))))
	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("the request %s", r.URL.Query().Get("Query"))
	}

	response, err := client.Do(request)

	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("the request %s", r.URL.Query().Get("Query"))
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("the request %s ", r.URL.Query().Get("Query"))
	}

	return http.StatusOK, string(body)
}
