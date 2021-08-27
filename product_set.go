package unaswrappergo

type SetProductRequestParams struct {
	Products []Product
}

type SetProductResponse struct{}

func (uo UnasObject) SetProduct(products []*Product)
