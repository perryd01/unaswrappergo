package unaswrappergo

import (
	"encoding/xml"
	"strconv"
)

type checkCustomerParams struct {
	XMLName  xml.Name `xml:"Params"`
	User     string   `xml:"User"`
	Password string   `xml:"Password"`
}

type checkCustomerResponse struct {
	Result string `xml:"Result"`
}

func (uo *UnasObject) CheckCustomer(user *string, password *string) (*bool, error) {
	params := checkCustomerParams{User: *user, Password: *password, XMLName: xml.Name{Local: "Params"}}
	b, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	respbody, err := uo.makeRequest(CheckCustomer, b)
	if err != nil {
		return nil, err
	}
	resp := checkCustomerResponse{}
	err = xml.Unmarshal(respbody, &resp)
	if err != nil {
		return nil, err
	}
	result, err := strconv.ParseBool(resp.Result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
