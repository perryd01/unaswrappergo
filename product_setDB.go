package unaswrappergo

import "encoding/xml"

type setProductDBRequest struct {
	Params *SetProductDBParams `xml:"Params"`
}

// SetProductDBParams Request params for SetProductDB request, where a URL can be set for import,
// more info at: https://unas.hu/tudastar/api/product#setproductdb-keres
type SetProductDBParams struct {
	URL     string `xml:"Url"`
	DB      string `xml:"DB"`
	DelType string `xml:"DelType"`
	Lang    string `xml:"Lang"`
}

// SetProductDBResponse Response from a SetProductDB request that contains logs from the result of the import,
// more info at: https://unas.hu/tudastar/api/product#setproductdb-valasz
type SetProductDBResponse struct {
	ModifiedProducts uint64                     `xml:"setProductDB>Ok>Modify,omitempty"`
	AddedProducts    uint64                     `xml:"setProductDB>Ok>Add,omitempty"`
	DeletedProducts  uint64                     `xml:"setProductDB>Ok>Delete,omitempty"`
	Errors           SetProductDBResponseErrors `xml:"setProductDB>Error,omitempty"`
}

type SetProductDBResponseErrors struct {
	UnknownColumns    uint64 `xml:"UnknownColumns,omitempty"`
	FaultyProducts    uint64 `xml:"FaultyProducts,omitempty"`
	SkuDuplicity      uint64 `xml:"SKU_Duplicity,omitempty"`
	NewProductFewData uint64 `xml:"NewProductFewData,omitempty"`
	LimitError        uint64
}

// SetProductDB Importing a csv? from an url into the webshop database,
// more info at: https://unas.hu/tudastar/api/product#setproductdb-funkcio
func (uo UnasObject) SetProductDB(params *SetProductDBParams) (*SetProductDBResponse, error) {
	reqBody := setProductDBRequest{Params: params}
	b, err := xml.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(SetProductDB, b)
	if err != nil {
		return nil, err
	}

	setPDBResponse := SetProductDBResponse{}
	err = xml.Unmarshal(resp, &setPDBResponse)
	if err != nil {
		return nil, err
	}

	return &setPDBResponse, nil
}
