package unaswrappergo

import (
	"encoding/xml"
	"net/url"
)

type getProductDBRequest struct {
	Params *GetProductDBParameters `xml:"Params"`
}

// GetProductDBParameters Parameters required for GetProductDB request,
// more info at: https://unas.hu/tudastar/api/product#getproductdb-keres
type GetProductDBParameters struct {
	Format              string `xml:"Format,omitempty"`
	Compress            string `xml:"Compress,omitempty"`
	LimitStart          string `xml:"LimitStart,omitempty"`
	LimitNum            string `xml:"LimitNum,omitempty"`
	Category            string `xml:"Category,omitempty"`
	Lang                string `xml:"Lang,omitempty"`
	GetName             string `xml:"GetName,omitempty"`
	GetStatus           string `xml:"GetStatus,omitempty"`
	GetPrice            string `xml:"GetPrice,omitempty"`
	GetPriceSale        string `xml:"GetPriceSale,omitempty"`
	GetPriceSpecial     string `xml:"GetPriceSpecial,omitempty"`
	GetCategory         string `xml:"GetCategory,omitempty"`
	GetDescriptionShort string `xml:"GetDescriptionShort,omitempty"`
	GetDescriptionLong  string `xml:"GetDescriptionLong,omitempty"`
	GetLink             string `xml:"GetLink,omitempty"`
	GetMinQty           string `xml:"GetMinQty,omitempty"`
	GetStock            string `xml:"GetStock,omitempty"`
	GetUnit             string `xml:"GetUnit,omitempty"`
	GetAlterUnit        string `xml:"GetAlterUnit,omitempty"`
	GetWeight           string `xml:"GetWeight,omitempty"`
	GetPoint            string `xml:"GetPoint,omitempty"`
	GetParam            string `xml:"GetParam,omitempty"`
	GetData             string `xml:"GetData,omitempty"`
	GetAttach           string `xml:"GetAttach,omitempty"`
	GetPack             string `xml:"GetPack,omitempty"`
	GetVariant          string `xml:"GetVariant,omitempty"`
	GetAlterCategory    string `xml:"GetAlterCategory,omitempty"`
	GetImage            string `xml:"GetImage,omitempty"`
	GetURL              string `xml:"GetURL,omitempty"`
	GetExport           string `xml:"GetExport,omitempty"`
	GetOrder            string `xml:"GetOrder,omitempty"`
	GetExplicit         string `xml:"GetExplicit,omitempty"`
	GetOnlineContent    string `xml:"GetOnlineContent,omitempty"`
	GetSEO              string `xml:"GetSeo,omitempty"`
	GetType             string `xml:"GetType,omitempty"`
	GetImageConnect     string `xml:"GetImageConnect,omitempty"`
	GetDiscount         string `xml:"GetDiscount,omitempty"`
	GetUnitStep         string `xml:"GetUnitStep,omitempty"`
	GetShipping         string `xml:"GetShipping,omitempty"`
	GetPayment          string `xml:"GetPayment,omitempty"`
	GetCustomerGroup    string `xml:"GetCustomerGroup,omitempty"`
	GetAddModDate       string `xml:"GetAddModDate,omitempty"`
	GetService          string `xml:"GetService,omitempty"`
}

// Simplified response from GetProductDB request,
// more info at: https://unas.hu/tudastar/api/product#getproductdb-valasz
type getProductDBResponse struct {
	URL string `xml:"getProductDB>Url"`
}

// GetProductDB Returns an url where all the products can be downloaded in a .csv file,
// the link expires after 1 hour
// more info at: https://unas.hu/tudastar/api/product#getproductdb-funkcio
func (uo UnasObject) GetProductDB(params *GetProductDBParameters) (*url.URL, error) {
	requestBody := getProductDBRequest{Params: params}

	b, err := xml.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := uo.makeRequest(endpointEnumType(GetProductDB), b)
	if err != nil {
		return nil, err
	}

	gpDBresponse := getProductDBResponse{}

	err = xml.Unmarshal(response, &gpDBresponse)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(gpDBresponse.URL)
	if err != nil {
		return nil, err
	}

	return url, nil
}
