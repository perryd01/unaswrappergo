package unaswrappergo

type Product struct {
	Action      string       `json:"Action"`
	Sku         string       `json:"Sku"`
	Name        *string      `json:"Name,omitempty"`
	Unit        *string      `json:"Unit,omitempty"`
	MinimumQty  *string      `json:"MinimumQty,omitempty"`
	MaximumQty  *string      `json:"MaximumQty,omitempty"`
	Description *Description `json:"Description,omitempty"`
	Prices      *Prices      `json:"Prices,omitempty"`
	Categories  *Categories  `json:"Categories,omitempty"`
	Statuses    *Statuses    `json:"Statuses,omitempty"`
}

type Categories struct {
	Category Category `json:"Category"`
}

type Category struct {
	Type string `json:"Type"`
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type Description struct {
	Short string `json:"Short"`
	Long  string `json:"Long"`
}

type Prices struct {
	Price []Price `json:"Price"`
}

type Price struct {
	Type  string  `json:"Type"`
	Net   string  `json:"Net"`
	Gross string  `json:"Gross"`
	Start *string `json:"Start,omitempty"`
	End   *string `json:"End,omitempty"`
}

type Statuses struct {
	Status Status `json:"Status"`
}

type Status struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}
