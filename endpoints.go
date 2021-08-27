package unaswrappergo

type endpointEnumType string

const (
	GetNewsletter endpointEnumType = "https://api.unas.eu/shop/getNewsletter"
	SetNewsletter                  = "https://api.unas.eu/shop/setNewsletter"
	LoginEnd                       = "https://api.unas.eu/shop/login"
	GetProduct                     = "https://api.unas.eu/shop/getProduct"
	SetProduct                     = "https://api.unas.eu/shop/setProduct"
	GetProductDB                   = "https://api.unas.eu/shop/getProductDB"
	SetProductDB                   = "https://api.unas.eu/shop/setProductDB"
)
