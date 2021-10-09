package unaswrappergo

import "encoding/xml"

func (uo *UnasObject) GetCategory(params *GetCategoryRequestParams) ([]*Category, error) {
	bodyMarshaled, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(GetCategory, bodyMarshaled)
	if err != nil {
		return nil, err
	}
	var categories = getCategoryResponse{}
	err = xml.Unmarshal(resp, &categories)
	if err != nil {
		return nil, err
	}
	return categories.Categories, nil
}

func (uo *UnasObject) SetCategory(params []*SetCategoryRequestTypes) ([]*SetCategoryResponseTypes, error) {
	bodyMarshaled, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(SetCategory, bodyMarshaled)
	if err != nil {
		return nil, err
	}
	var categoryResponse = setCategoryResponse{}
	err = xml.Unmarshal(resp, &categoryResponse)
	if err != nil {
		return nil, err
	}
	return categoryResponse.Categories, nil
}

// https://unas.hu/tudastar/api/category#getcategory-keres
type GetCategoryRequestParams struct {
	XMLName     xml.Name       `xml:"Params"`
	ID          string         `xml:"Id,omitempty"`          // Category Unique ID
	Name        string         `xml:"Name,omitempty"`        // Category Name
	Parent      string         `xml:"Parent,omitempty"`      // Category's parent's ID
	TimeStart   *UnasTimeStamp `xml:"TimeStart,omitempty"`   // Categories modified after this timestamp
	TimeEnd     *UnasTimeStamp `xml:"TimeEnd,omitempty"`     // Categories modified before this timestamp
	DateStart   *UnasDate      `xml:"DateStart,omitempty"`   // Categories modified after this date
	DateEnd     *UnasDate      `xml:"DateEnd,omitempty"`     // Categories modified before this date
	ContentType string         `xml:"ContentType,omitempty"` // Level of details: "minimal", "normal","full"
	LimitStart  string         `xml:"LimitStart,omitempty"`  // Shift starting point of query by N items
	LimitNum    string         `xml:"LimitNum,omitempty"`    // Limit Maxmimum Category data queried overall
	History     string         `xml:"History,omitempty"`     // Log Category editing history data. "no", "yes"
}

type getCategoryResponse struct {
	XMLName    xml.Name    `xml:"Categories"`
	Categories []*Category `xml:"Category"`
}

type SetCategoryRequestTypes struct {
	XMLName xml.Name `xml:"Category"`
	ID      string   `xml:"Id,omitempty"`
	Name    string   `xml:"Name,omitempty"`
	Action  string   `xml:"Action,omitempty"`
}

type setCategoryResponse struct {
	XMLName    xml.Name                    `xml:"Categories"`
	Categories []*SetCategoryResponseTypes `xml:"Category"`
}

type SetCategoryResponseTypes struct {
	XMLName xml.Name `xml:"Category"`
	ID      string   `xml:"Id"`
	Action  string   `xml:"Action"`
	Status  string   `xml:"Status"`
	Error   string   `xml:"Error"`
}

type Category struct {
	XMLName xml.Name `xml:"Category,omitempty"`
	Action  string   `xml:"Action,omitempty"`
	ID      string   `xml:"Id,omitempty"`
	History struct {
		Event struct {
			Action string `xml:"Action,omitempty"`
			Time   string `xml:"Time,omitempty"`
			ID     string `xml:"Id,omitempty"`
		} `xml:"Event"`
	} `xml:"History"`
	Name        string `xml:"Name,omitempty"`
	URL         string `xml:"Url,omitempty"`
	SefUrl      string `xml:"SefUrl,omitempty"`
	AltUrl      string `xml:"AltUrl,omitempty"`
	AltUrlBlank string `xml:"AltUrlBlank,omitempty"`
	Display     struct {
		Page string `xml:"Page,omitempty"`
		Menu string `xml:"Menu,omitempty"`
	} `xml:"Display"`
	PageLayout struct {
		ProductList  string `xml:"ProductList,omitempty"`
		CategoryList string `xml:"CategoryList,omitempty"`
	} `xml:"PageLayout"`
	Parent struct {
		ID   string `xml:"Id,omitempty"`
		Tree string `xml:"Tree,omitempty"`
	} `xml:"Parent"`
	Order    string `xml:"Order,omitempty"`
	Products struct {
		All string `xml:"All,omitempty"`
		New string `xml:"New,omitempty"`
	} `xml:"Products"`
	Texts struct {
		Top    string `xml:"Top,omitempty"`
		Bottom string `xml:"Bottom,omitempty"`
		Menu   string `xml:"Menu,omitempty"`
	} `xml:"Texts"`
	Meta struct {
		Keywords    string `xml:"Keywords,omitempty"`
		Description string `xml:"Description,omitempty"`
		Title       string `xml:"Title,omitempty"`
		Robots      string `xml:"Robots,omitempty"`
	} `xml:"Meta"`
	AutomaticMeta struct {
		Keywords    string `xml:"Keywords,omitempty"`
		Description string `xml:"Description,omitempty"`
		Title       string `xml:"Title,omitempty"`
	} `xml:"AutomaticMeta"`
	LastModTime string `xml:"LastModTime,omitempty"`
	Image       struct {
		URL string `xml:"Url,omitempty"`
		OG  string `xml:"OG,omitempty"`
	} `xml:"Image"`
	Tags []string `xml:"Tags>Tag,omitempty"`
}
