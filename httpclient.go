package httpclient

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(URL string, values map[string]string) ([]byte, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}

	baseURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	for key, value := range values {
		params.Add(key, value)
	}

	baseURL.RawQuery = params.Encode()

	URL = baseURL.String()

	resp, err := http.Get(URL)
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

func Post(URL string, values map[string]string) ([]byte, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}

	params := url.Values{}
	for key, value := range values {
		params.Add(key, value)
	}

	resp, err := http.PostForm(URL, params)
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
