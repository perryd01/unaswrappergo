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

type GetCustomerParams struct {
	XMLName        xml.Name `xml:"Params"`
	ID             string   `xml:"Id,omitempty"`
	Email          string   `xml:"Email,omitempty"`
	Username       string   `xml:"Username,omitempty"`
	RegTimeStart   string   `xml:"RegTimeStart,omitempty"`
	RegTimeEnd     string   `xml:"RegTimeEnd,omitempty"`
	ModTimeStart   string   `xml:"ModTimeStart,omitempty"`
	ModTimeEnd     string   `xml:"ModTimeEnd,omitempty"`
	LoginTimeStart string   `xml:"LoginTimeStart,omitempty"`
	LoginTimeEnd   string   `xml:"LoginTimeEnd,omitempty"`
	LimitStart     string   `xml:"LimitStart,omitempty"`
	LimitNum       string   `xml:"LimitNum,omitempty"`
}

type Customer struct {
	Action          string `xml:"Action"`
	ID              string `xml:"Id"`
	Email           string `xml:"Email"`
	Username        string `xml:"Username"`
	Password        string `xml:"Password"`
	PasswordCrypted string `xml:"PasswordCrypted"`
	Contact         struct {
		Name   string `xml:"Name"`
		Phone  string `xml:"Phone"`
		Mobile string `xml:"Mobile"`
		Lang   string `xml:"Lang"`
	} `xml:"Contact"`
	Addresses struct {
		Invoice struct {
			Name         string `xml:"Name"`
			ZIP          string `xml:"ZIP"`
			City         string `xml:"City"`
			Street       string `xml:"Street"`
			StreetName   string `xml:"StreetName"`
			StreetType   string `xml:"StreetType"`
			StreetNumber string `xml:"StreetNumber"`
			County       string `xml:"County"`
			Country      string `xml:"Country"`
			CountryCode  string `xml:"CountryCode"`
			TaxNumber    string `xml:"TaxNumber"`
			EUTaxNumber  string `xml:"EUTaxNumber"`
			CustomerType string `xml:"CustomerType"`
		} `xml:"Invoice"`
		Shipping struct {
			Name         string `xml:"Name"`
			ZIP          string `xml:"ZIP"`
			City         string `xml:"City"`
			Street       string `xml:"Street"`
			StreetName   string `xml:"StreetName"`
			StreetType   string `xml:"StreetType"`
			StreetNumber string `xml:"StreetNumber"`
			County       string `xml:"County"`
			Country      string `xml:"Country"`
			CountryCode  string `xml:"CountryCode"`
		} `xml:"Shipping"`
		Other []struct {
			Name        string `xml:"Name"`
			ZIP         string `xml:"ZIP"`
			City        string `xml:"City"`
			Street      string `xml:"Street"`
			County      string `xml:"County"`
			Country     string `xml:"Country"`
			CountryCode string `xml:"CountryCode"`
		} `xml:"Other"`
	} `xml:"Addresses"`
	Param []struct {
		ID    string `xml:"Id"`
		Name  string `xml:"Name"`
		Value string `xml:"Value"`
	} `xml:"Params>Param"`
	Dates struct {
		Registration string `xml:"Registration"`
		Modification string `xml:"Modification"`
		Login        string `xml:"Login"`
	} `xml:"Dates"`
	Group struct {
		ID   string `xml:"Id"`
		Name string `xml:"Name"`
	} `xml:"Group"`
	Authorize struct {
		Customer string `xml:"Customer"`
		Admin    string `xml:"Admin"`
	} `xml:"Authorize"`
	Discount struct {
		Total  string `xml:"Total"`
		Direct string `xml:"Direct"`
	} `xml:"Discount"`
	PointsAccount struct {
		Balance string `xml:"Balance"`
	} `xml:"PointsAccount"`
	Newsletter struct {
		Subscribed string `xml:"Subscribed"`
		Authorized string `xml:"Authorized"`
	} `xml:"Newsletter"`
	Comment      string `xml:"Comment"`
	Restrictions struct {
		Restriction []struct {
			Type string `xml:"Type"`
			ID   string `xml:"Id"`
			Name string `xml:"Name"`
		} `xml:"Restriction"`
	} `xml:"Restrictions"`
	Others struct {
		FacebookConnect string `xml:"FacebookConnect"`
		Ip              string `xml:"Ip"`
		Referer         string `xml:"Referer"`
	} `xml:"Others"`
}