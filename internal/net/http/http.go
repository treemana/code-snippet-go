package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	// add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	return do(req, params, headers)
}

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return do(req, params, headers)
}

func do(request *http.Request, params, headers map[string]string) (*http.Response, error) {
	// add params
	if params != nil && len(params) > 0 {
		q := request.URL.Query()
		for key, val := range params {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()
	}

	// add headers
	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	log.Println(http.MethodPost, request.URL.String())
	return client.Do(request)
}

// GetContentType 获取 url 对应资源的 ContentType
func GetContentType(url string) (string, error) {

	if len(url) == 0 {
		return "", errors.New("empty url")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var data = make([]byte, 512)
	if _, err = resp.Body.Read(data); err != nil {
		return "", err
	}

	return http.DetectContentType(data), nil
}
