package unaswrappergo

type endpointEnumType string

const (
	GetNewsletter endpointEnumType = "https://api.unas.eu/shop/getNewsletter"
	SetNewsletter                  = "https://api.unas.eu/shop/setNewsletter"
	GetProduct                     = "https://api.unas.eu/shop/getProduct"
	SetProduct                     = "https://api.unas.eu/shop/setProduct"
	LoginEnd                       = "https://api.unas.eu/shop/login"
)
