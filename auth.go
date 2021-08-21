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
	Token       string      `xml:"Token"`
	Expire      string      `xml:"Expire"`
	Permissions Permissions `xml:"Permissions"`
	Status      string      `xml:"Status"`
}

type Permissions struct {
	Permission []string `xml:"Permission"`
}

type AuthPassRequest struct {
	Auth Auth `json:"Auth"`
}

type Auth struct {
	Username      string `json:"Username"`
	PasswordCrypt string `json:"PasswordCrypt"`
	ShopID        string `json:"ShopId"`
	AuthCode      string `json:"AuthCode"`
}

const loginEndpoint = "https://api.unas.eu/shop/login"

func AuthAPI(apikey string) (*LoginAPIResponse, error) {
	payload := LoginAPIRequest{Params: Params{APIKey: apikey}}
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

	return &xmlresponse, nil
}

func AuthPass(a Auth) (*LoginAPIResponse, error) {
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

	return &xmlresponse, nil

}
