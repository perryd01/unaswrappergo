package unaswrappergo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Func that handles making requests to specific endpointEnumType endpoints.
func (uo *UnasObject) makeRequest(endpoint endpointEnumType, body []byte) ([]byte, error) {
	/*
		if tokenExpired(*uo.Login.Expire.ToTime()) {
			return nil, errors.New("login token already expired")
		}
	*/

	reqBuf := bytes.NewBuffer([]byte(xml.Header))
	_, err := reqBuf.Write(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", string(endpoint), reqBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", "Bearer "+uo.Login.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200 on " + string(endpoint))
	}

	returnable, _ := ioutil.ReadAll(resp.Body)

	return returnable, nil
}
