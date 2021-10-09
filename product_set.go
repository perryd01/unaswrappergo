package unaswrappergo

import "encoding/xml"

type setProductRequestParams struct {
	Products []*Product `xml:"Product"`
}

type setProductRequest struct {
	P setProductRequestParams `xml:"Products"`
}

type setProductResponse struct {
	setProductStatuses `xml:"Products"`
}

type setProductStatuses struct {
	Products []*productStatus `xml:"Product"`
}

type productStatus struct {
	ID     string `xml:"Id,omitempty"`
	Sku    string `xml:"Sku,omitempty"`
	Action string `xml:"Action,omitempty"`
	Status string `xml:"Status,omitempty"`
	Error  string `xml:"Error,omitempty"`
}

func (uo *UnasObject) SetProduct(products []*Product) ([]*productStatus, error) {
	body := setProductRequest{P: setProductRequestParams{Products: products}}
	b, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}
	r, err := uo.makeRequest(SetProduct, b)
	if err != nil {
		return nil, err
	}

	spResponse := setProductResponse{}

	err = xml.Unmarshal(r, &spResponse)
	if err != nil {
		return nil, err
	}

	return spResponse.setProductStatuses.Products, nil
}
