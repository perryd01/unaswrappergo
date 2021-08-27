package unaswrappergo

import "encoding/xml"

type setProductDBRequest struct {
	Params *SetProductDBParams `json:"Params"`
}

type SetProductDBParams struct {
	URL     string `json:"Url"`
	DB      string `json:"DB"`
	DelType string `json:"DelType"`
	Lang    string `json:"Lang"`
}

type SetProductDBResponse struct {
	ModifiedProducts uint64                     `xml:"setProductDB>Ok>Modify,omitempty"`
	AddedProducts    uint64                     `xml:"setProductDB>Ok>Add,omitempty"`
	DeletedProducts  uint64                     `xml:"setProductDB>Ok>Delete,omitempty"`
	Errors           SetProductDBResponseErrors `xml:"setProductDB>Error,omitempty"`
}

type SetProductDBResponseErrors struct {
	UnknownColumns uint64 `xml:"UnknownColumns,omitempty"`
	FaultyProducts uint64  `xml:"FaultyProducts,omitempty"`
	SKU_Duplicity uint64 `xml:"SKU_Duplicity,omitempty"`
	NewProductFewData uint64 `xml:"NewProductFewData,omitempty"`
	LimitError uint64
}

func (uo UnasObject) SetProductDB(params *SetProductDBParams) (*SetProductDBResponse, error){
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
