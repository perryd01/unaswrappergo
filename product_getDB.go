package unaswrappergo

import (
	"encoding/xml"
	"net/url"
)

type getProductDBRequest struct {
	Params *GetProductDBParameters `xml:"Params"`
}

// Parameters required for GetProductDB request,
// more info at: https://unas.hu/tudastar/api/product#getproductdb-keres
type GetProductDBParameters struct {
	Format              string `xml:"Format"`
	Compress            string `xml:"Compress"`
	LimitStart          string `xml:"LimitStart"`
	LimitNum            string `xml:"LimitNum"`
	Category            string `xml:"Category"`
	Lang                string `xml:"Lang"`
	GetName             string `xml:"GetName"`
	GetStatus           string `xml:"GetStatus"`
	GetPrice            string `xml:"GetPrice"`
	GetPriceSale        string `xml:"GetPriceSale"`
	GetPriceSpecial     string `xml:"GetPriceSpecial"`
	GetCategory         string `xml:"GetCategory"`
	GetDescriptionShort string `xml:"GetDescriptionShort"`
	GetDescriptionLong  string `xml:"GetDescriptionLong"`
	GetLink             string `xml:"GetLink"`
	GetMinQty           string `xml:"GetMinQty"`
	GetStock            string `xml:"GetStock"`
	GetUnit             string `xml:"GetUnit"`
	GetAlterUnit        string `xml:"GetAlterUnit"`
	GetWeight           string `xml:"GetWeight"`
	GetPoint            string `xml:"GetPoint"`
	GetParam            string `xml:"GetParam"`
	GetData             string `xml:"GetData"`
	GetAttach           string `xml:"GetAttach"`
	GetPack             string `xml:"GetPack"`
	GetVariant          string `xml:"GetVariant"`
	GetAlterCategory    string `xml:"GetAlterCategory"`
	GetImage            string `xml:"GetImage"`
	GetURL              string `xml:"GetURL"`
	GetExport           string `xml:"GetExport"`
	GetOrder            string `xml:"GetOrder"`
	GetExplicit         string `xml:"GetExplicit"`
	GetOnlineContent    string `xml:"GetOnlineContent"`
	GetSEO              string `xml:"GetSeo"`
	GetType             string `xml:"GetType"`
	GetImageConnect     string `xml:"GetImageConnect"`
	GetDiscount         string `xml:"GetDiscount"`
	GetUnitStep         string `xml:"GetUnitStep"`
	GetShipping         string `xml:"GetShipping"`
	GetPayment          string `xml:"GetPayment"`
	GetCustomerGroup    string `xml:"GetCustomerGroup"`
	GetAddModDate       string `xml:"GetAddModDate"`
	GetService          string `xml:"GetService"`
}

// Simplified response from GetProductDB request,
// more info at: https://unas.hu/tudastar/api/product#getproductdb-valasz
type getProductDBResponse struct {
	URL string `xml:"getProductDB>Url"`
}

// Returns an url where all of the products can be downloaded in a .csv file,
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
