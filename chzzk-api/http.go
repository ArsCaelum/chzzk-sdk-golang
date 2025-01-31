package chzzk_api

import (
	"io"
	"log"
	"net/http"
)

func sendGet(url string) (error, interface{}) {
	var request *http.Request
	var err error

	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err, ""
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err, ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(response.Body)

	// TODO : check status code

	bodyBytes, err := io.ReadAll(response.Body)

	if nil != err {
		log.Print(err.Error())
	}

	return nil, string(bodyBytes)
}
