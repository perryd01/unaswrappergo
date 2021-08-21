package unaswrappergo

type getProductRequest struct {
	Params GetProductRequestParams `xml:"Params"`
}

type GetProductRequestParams struct {
	StatusBase   *string `xml:"StatusBase"`
	ID           *string `xml:"Id"`
	Sku          *string `xml:"Sku"`
	Parent       *string `xml:"Parent"`
	TimeStart    *string `xml:"TimeStart"`
	TimeEnd      *string `xml:"TimeEnd"`
	DateStart    *string `xml:"DateStart"`
	DateEnd      *string `xml:"DateEnd"`
	ContentType  *string `xml:"ContentType"`
	ContentParam *string `xml:"ContentParam"`
	LimitStart   *string `xml:"LimitStart"`
	LimitNum     *string `xml:"LimitNum"`
}

func (uo *UnasObject) GetProduct(p *GetProductRequestParams)
