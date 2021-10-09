package unaswrappergo

type GetStockParams struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"Id"`
	Sku        string `xml:"Sku"`
	Variant1   string `xml:"Variant1"`
	Variant2   string `xml:"Variant2"`
	Variant3   string `xml:"Variant3"`
	LimitStart string `xml:"LimitStart"`
	LimitNum   string `xml:"LimitNum"`
}

type getStockRequest struct {
	Params GetStockParams `xml:"Params"`
}

type getStockResponse struct {
	Products GetStockResponseProduct `xml:"Products"`
}

type GetStockResponseProduct struct {
}

// func (uo UnasObject) getStock(params GetStockParams) ([]*GetStockResponseProduct, error) {}
// func (uo UnasObject) setStock()                                                          {}
