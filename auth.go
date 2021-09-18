package unaswrappergo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
)

type loginAPIRequest struct {
	Params loginAPIParams `xml:"Params"`
}

type loginAPIParams struct {
	APIKey  string `xml:"ApiKey"`
	XMLName xml.Name
}

type loginAPIResponse struct {
	Login login `xml:"Login"`
}

// Struct which contains data for request authentication
type login struct {
	Token       string       `xml:"Login>Token"`
	Expire      UnasTimeDate `xml:"Login>Expire"`
	Permissions permissions  `xml:"Login>Permissions"`
	Status      string       `xml:"Login>Status"`
}

// Permissions for allowed methods towards Unas
type permissions struct {
	Permission []string `xml:"Permission"`
}

type authPassRequest struct {
	Auth Auth `xml:"Auth"`
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

	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful post")
	}
	xmlresponse := loginAPIResponse{}
	err = xml.Unmarshal(body, &xmlresponse)
	if err != nil {
		return nil, err
	}

	uo := UnasObject{
		Login: xmlresponse.Login,
	}

	return &uo, nil
}

// AuthwithPass Authenticating using a User:Pass combo.
// https://unas.hu/tudastar/api/authorization#felhasznalonev-alapu-azonositas
func AuthwithPass(a Auth) (*UnasObject, error) {
	xmlpayload, err := xml.Marshal(authPassRequest{Auth: a})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", string(LoginEndPoint), bytes.NewBuffer(xmlpayload))
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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful post")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	xmlresponse := loginAPIResponse{}
	err = xml.Unmarshal(body, &xmlresponse)
	if err != nil {
		return nil, err
	}

	uo := UnasObject{
		Login: xmlresponse.Login,
	}

	return &uo, nil
}
