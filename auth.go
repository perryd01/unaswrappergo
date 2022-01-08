package unaswrappergo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
)

type loginAPIParams struct {
	APIKey  string `xml:"ApiKey"`
	XMLName xml.Name
}

// Struct which contains data for request authentication
type login struct {
	XMLName     xml.Name     `xml:"Login"`
	Token       string       `xml:"Token"`
	Expire      UnasTimeDate `xml:"Expire"`
	Permissions permissions  `xml:"Permissions"`
	Status      string       `xml:"Status"`
}

// Permissions for allowed methods towards Unas
type permissions struct {
	Permission []string `xml:"Permission"`
}

// Auth Struct for login with user:pass
type Auth struct {
	Username      string `xml:"Username"`
	PasswordCrypt string `xml:"PasswordCrypt"`
	ShopID        string `xml:"ShopId"`
	AuthCode      string `xml:"AuthCode"`
}

// AuthwithAPIKey Authenticating using an API key.
// https://unas.hu/tudastar/api/authorization#api-kulcs-alapu-azonositas
func AuthwithAPIKey(apikey string) (*UnasObject, error) {

	if apikey == "" {
		return nil, errors.New("empty apiKey")
	}

	payload := loginAPIParams{
		APIKey:  apikey,
		XMLName: xml.Name{Local: "Params"},
	}
	xmlpayload, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}

	reqBuf := bytes.NewBuffer([]byte(xml.Header))
	_, err = reqBuf.Write(xmlpayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", string(LoginEndPoint), reqBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mime.TypeByExtension(".xml"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful post")
	}

	xmlresponse := login{}
	err = xml.Unmarshal(body, &xmlresponse)
	if err != nil {
		return nil, err
	}

	uo := UnasObject{
		Login: xmlresponse,
	}

	if uo.Login.Status != "ok" {
		return nil, errors.New("auth status not OK")
	}

	return &uo, nil
}
