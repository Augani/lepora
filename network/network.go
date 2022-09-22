package network

import (
	"bytes"
	"errors"
	"net/http"
)

type NetworkOptions struct {
	//URL for sending logs
	URL string
	//the access token to be used to send logs
	AccessToken string
	//app key
	AppKey string
}

//make a post request to server to send logs
func PostLogs(log string, network NetworkOptions) error {
	//check if the url is valid
	if network.URL == "" {
		return errors.New("url is required")
	}
	//check if the access token is valid
	if network.AccessToken == "" {
		return errors.New("access token is required")
	}
	//send the logs to the server
	//return the response
	//create json payload of appkey and logs
	payload := []byte(`{"appkey":"` + network.AppKey + `","logs":"` + log + `"}`)
	request, err := http.NewRequest("POST", network.URL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+network.AccessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil

}
