package unaswrappergo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func (uo *UnasObject) makeRequest(endpoint endpointEnumType, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", string(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", "Bearer "+uo.Login.Login.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200 on " + string(endpoint))
	}

	returnbody, _ := ioutil.ReadAll(resp.Body)

	return returnbody, nil
}
