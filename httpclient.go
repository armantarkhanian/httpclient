package httpclient

import (
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
)

func Get(url string, values map[string]string) ([]byte, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	baseURL, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	params := neturl.Values{}
	for key, value := range values {
		params.Add(key, value)
	}

	baseURL.RawQuery = params.Encode()

	url = baseURL.String()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Post(url string, values map[string]string) ([]byte, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	params := neturl.Values{}
	for key, value := range values {
		params.Add(key, value)
	}

	resp, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
