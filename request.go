package unaswrappergo

import (
	"bytes"
	"errors"
	"net/http"
)

func (uo *UnasObject) makeRequest(endpoint endpointEnumType, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", string(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization": "Bearer " + uo.Login.Login.Token,)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200 on " + string(endpoint))
	}
	return resp, nil
}
