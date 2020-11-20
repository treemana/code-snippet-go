package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	type args struct {
		url     string
		body    interface{}
		params  map[string]string
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Post(tt.args.url, tt.args.body, tt.args.params, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
	var (
		url      = "http://google.com"
		response *http.Response
		err      error
	)

	response, err = Post(url, nil, nil, nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	defer func() {
		_ = response.Body.Close()
	}()

	if err != nil {
		log.Panicln(err)
		return
	}

	log.Println(string(body))
}
