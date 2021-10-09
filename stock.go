package unaswrappergo

import "encoding/xml"

type GetStockParams struct {
	XMLName    xml.Name `xml:"Params"`
	ID         string   `xml:"Id"`
	Sku        string   `xml:"Sku"`
	Variant1   string   `xml:"Variant1"`
	Variant2   string   `xml:"Variant2"`
	Variant3   string   `xml:"Variant3"`
	LimitStart string   `xml:"LimitStart"`
	LimitNum   string   `xml:"LimitNum"`
}

type SetStockParams struct {
	XMLName xml.Name `xml:"Products"`
	Product []struct {
		Action string `xml:"Action,omitempty"`
		Sku    string `xml:"Sku,omitempty"`
		Stocks struct {
			Stock []struct {
				Qty      string `xml:"Qty,omitempty"`
				Price    string `xml:"Price,omitempty"`
				Variants struct {
					Variant []string `xml:"Variant,omitempty"`
				} `xml:"Variants,omitempty"`
			} `xml:"Stock,omitempty"`
		} `xml:"Stocks,omitempty"`
	} `xml:"Product"`
}

type SetStockProductStatus struct {
	ID     string `xml:"Id"`
	Sku    string `xml:"Sku"`
	Action string `xml:"Action"`
	Status string `xml:"Status"`
	Error  string `xml:"Error"`
}

type setStockResponse struct {
	XMLName xml.Name                 `xml:"Products"`
	Stocks  []*SetStockProductStatus `xml:"Product"`
}

type getStockResponse struct {
	XMLName xml.Name            `xml:"Products"`
	Product []*ProductStockData `xml:"Product,omitempty"`
}

type ProductStockData struct {
	Sku   string `xml:"Sku,omitempty"`
	Stock []struct {
		Variant []string `xml:"Variants>Variant"`
		Qty     string   `xml:"Qty"`
	} `xml:"Stocks>Stock"`
}

func (uo UnasObject) getStock(params GetStockParams) ([]*ProductStockData, error) {
	bodyMarshaled, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(GetStock, bodyMarshaled)
	if err != nil {
		return nil, err
	}
	var stocks = getStockResponse{}
	err = xml.Unmarshal(resp, &stocks)
	if err != nil {
		return nil, err
	}
	return stocks.Product, nil
}

func (uo UnasObject) setStock(params *SetStockParams) ([]*SetStockProductStatus, error) {
	bodyMarshaled, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(SetStock, bodyMarshaled)
	if err != nil {
		return nil, err
	}
	var stocks = setStockResponse{}
	err = xml.Unmarshal(resp, &stocks)
	if err != nil {
		return nil, err
	}
	return stocks.Stocks, nil
}
