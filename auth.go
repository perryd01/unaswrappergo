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

type authPassRequest struct {
	XMLName       xml.Name `xml:"Auth"`
	Username      string   `xml:"Username"`
	PasswordCrypt string   `xml:"PasswordCrypt"`
	ShopId        string   `xml:"ShopId"`
	AuthCode      string   `xml:"AuthCode"`
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

	return &uo, nil
}

// AuthwithPass Authenticating using a User:Pass combo.
// https://unas.hu/tudastar/api/authorization#felhasznalonev-alapu-azonositas

func AuthwithPass(a Auth) (*UnasObject, error) {
	xmlpayload, err := xml.Marshal(authPassRequest{XMLName: xml.Name{Local: "Auth"}, Username: a.Username, ShopId: a.ShopID, PasswordCrypt: a.PasswordCrypt, AuthCode: a.AuthCode})
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

	return &uo, nil
}
