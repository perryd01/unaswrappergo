package unaswrappergo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type LoginAPIRequest struct {
	Params LoginAPIParams `xml:"Params"`
}

type LoginAPIParams struct {
	APIKey string `xml:"ApiKey"`
}

type LoginAPIResponse struct {
	Login Login `xml:"Login"`
}

type Login struct {
	Token       string      `xml:"Login>Token"`
	Expire      string      `xml:"Login>Expire"`
	Permissions Permissions `xml:"Login>Permissions"`
	Status      string      `xml:"Login>Status"`
}

type Permissions struct {
	Permission []string `xml:"Permission"`
}

type AuthPassRequest struct {
	Auth Auth `xml:"Auth"`
}

type Auth struct {
	Username      string `xml:"Username"`
	PasswordCrypt string `xml:"PasswordCrypt"`
	ShopID        string `xml:"ShopId"`
	AuthCode      string `xml:"AuthCode"`
}

const loginEndpoint = "https://api.unas.eu/shop/login"

func AuthwithAPIKey(apikey string) (*UnasObject, error) {
	payload := LoginAPIRequest{Params: LoginAPIParams{APIKey: apikey}}
	xmlpayload, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", loginEndpoint, bytes.NewBuffer(xmlpayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful post")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	xmlresponse := LoginAPIResponse{}
	err = xml.Unmarshal(body, &xmlresponse)
	if err != nil {
		return nil, err
	}

	uo := UnasObject{
		Login: xmlresponse,
	}

	return &uo, nil
}

func AuthwithPass(a Auth) (*UnasObject, error) {
	xmlpayload, err := xml.Marshal(AuthPassRequest{Auth: a})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", loginEndpoint, bytes.NewBuffer(xmlpayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful post")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	xmlresponse := LoginAPIResponse{}
	err = xml.Unmarshal(body, &xmlresponse)
	if err != nil {
		return nil, err
	}

	uo := UnasObject{
		Login: xmlresponse,
	}

	return &uo, nil
}
