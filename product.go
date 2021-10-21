package unaswrappergo

import (
	"encoding/xml"
	"jaytaylor.com/html2text"
	"regexp"
)

type Product struct {
	XMLName             xml.Name                       `xml:"Product"`
	Action              string                         `xml:"Action,omitempty"`
	State               string                         `xml:"State"`
	ID                  string                         `xml:"Id"`
	Sku                 string                         `xml:"Sku"`
	SkuNew              string                         `xml:"SkuNew,omitempty"`
	History             []ProductEventType             `xml:"History>Event,omitempty"`
	Statuses            []ProductStatusType            `xml:"Statuses>Status,omitempty"`
	NoList              bool                           `xml:"NoList,omitempty"`
	Inquire             bool                           `xml:"Inquire,omitempty"`
	CustDiscountDisable bool                           `xml:"CustDiscountDisable,omitempty"`
	Explicit            bool                           `xml:"Explicit,omitempty"`
	Export              *ProductExportType             `xml:"Export,omitempty"`
	PublicInterval      *ProductPublicIntervalType     `xml:"PublicInterval,omitempty"`
	Name                string                         `xml:"Name,omitempty"`
	Unit                string                         `xml:"Unit,omitempty"`
	MinimumQty          string                         `xml:"MinimumQty,omitempty"`
	MaximumQty          string                         `xml:"MaximumQty,omitempty"`
	AlertQty            *uint64                        `xml:"AlertQty,omitempty"`
	UnitStep            string                         `xml:"UnitStep,omitempty"`
	AlterUnit           *ProductAlterUnitType          `xml:"AlterUnit,omitempty"`
	Weight              *float64                       `xml:"Weight,omitempty"`
	Point               *uint64                        `xml:"Point,omitempty"`
	BuyableWithPoint    *string                        `xml:"BuyableWithPoint,omitempty"`
	Description         *ProductDescriptionType        `xml:"Description,omitempty"`
	Prices              []ProductPriceType             `xml:"Prices>Price,omitempty"`
	Categories          []ProductCategoryType          `xml:"Categories>Category,omitempty"`
	Url                 *string                        `xml:"Url,omitempty"`
	SefUrl              *string                        `xml:"SefUrl,omitempty"`
	Images              *ProductImagesType             `xml:"Images,omitempty"`
	Variants            []ProductVariantType           `xml:"Variants>Variant,omitempty"`
	Datas               []ProductDataType              `xml:"Datas>Data,omitempty"`
	Params              []ProductParamType             `xml:"Params>Param,omitempty"`
	AdditionalDatas     []ProductAdditionalDataType    `xml:"AdditionalDatas>AdditionalData"`
	PackageProduct      string                         `xml:"PackageProduct,omitempty"` // A termék csomagtermék e vagy sem
	Stocks              *ProductStocksType             `xml:"Stocks,omitempty"`
	OnlineContent       *ProductOnlineContentType      `xml:"OnlineContent,omitempty"`
	Types               *ProductTypesType              `xml:"Types,omitempty"`
	Meta                *ProductMetaType               `xml:"Meta,omitempty"`
	AutomaticMeta       *ProductAutomaticMetaType      `xml:"AutomaticMeta,omitempty"`
	QtyDiscount         *ProductQtyDiscountType        `xml:"QtyDiscount,omitempty"`
	SimilarProducts     []ProductSimilarProductType    `xml:"SimilarProducts>SimilarProduct,omitempty"`
	AdditionalProducts  []ProductAdditionalProductType `xml:"AdditionalProducts>AdditionalProduct,omitempty"`
}

type ProductEventType struct {
	Action string `xml:"Action"`
	Time   string `xml:"Time"`
	Sku    string `xml:"Sku"`
	SkuOld string `xml:"SkuOld,omitempty"`
}

type ProductCategoryType struct {
	Type string `xml:"Type"`
	ID   string `xml:"Id"`
	Name string `xml:"Name"`
}

type ProductDescriptionType struct {
	Short *ProductDescriptionLongType `xml:"Short"`
	Long  *ProductDescriptionLongType `xml:"Long"`
}

type ProductDescriptionLongType struct {
	Original  *string
	PlainText *string
}

func (desc *ProductDescriptionLongType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if s == "" {
		desc.Original = nil
		desc.PlainText = nil
	}
	if err == nil {
		desc.Original = &s
		text, err := html2text.FromString(s)
		if err != nil {
			return err
		}
		r := regexp.MustCompile("(\\t|\\r?\\n)+")
		description := r.ReplaceAllString(text, " ")
		desc.PlainText = &description
	}
	return err
}

type ProductPriceType struct {
	Type      string   `xml:"Type"`
	Area      string   `xml:"Area,omitempty"`
	AreaName  string   `xml:"AreaName,omitempty"`
	Group     string   `xml:"Group,omitempty"`
	GroupName string   `xml:"GroupName,omitempty"`
	Net       float64  `xml:"Net"`
	Gross     float64  `xml:"Gross"`
	Start     *string  `xml:"Start,omitempty"` // TODO: sajat tipus
	End       *string  `xml:"End,omitempty"`   // TODO: sajat tipus
	SaleNet   *float64 `xml:"SaleNet,omitempty"`
	SaleGross *float64 `xml:"SaleGross,omitempty"`
	SaleStart *string  `xml:"SaleStart,omitempty"`
	SaleEnd   *string  `xml:"SaleEnd,omitempty"`
	Percent   *string  `xml:"Percent,omitempty"`
}

type ProductStatusType struct {
	Type  string         `xml:"Type"`
	Id    string         `xml:"Id,omitempty"`
	Name  string         `xml:"Name,omitempty"`
	Value statusBaseEnum `xml:"Value"`
}

type ProductExportType struct {
	Status    bool     `xml:"Status"` // Engedélyezve(1) vagy Tiltva(0)
	Forbidden []string `xml:"Forbidden>Format"`
}

type ProductPublicIntervalType struct {
	Start string `xml:"Start,omitempty"` // TODO: format miatt kulon tipus
	End   string `xml:"End,omitempty"`   // TODO: format miatt kulon tipus
}

type ProductAlterUnitType struct {
	Quantity uint64 `xml:"Qty"`
	Unit     string `xml:"Unit"`
}

type ProductImagesType struct {
	DefaultFilename string             `xml:"DefaultFilename,omitempty"`
	DefaultAlt      string             `xml:"DefaultAlt,omitempty"`
	Og              *uint64            `xml:"OG,omitempty"`
	Version         string             `xml:"Version,omitempty"`
	Image           []ProductImageType `xml:"Image,omitempty"`
}

type ProductImageType struct {
	Type     string                  `xml:"Type,omitempty"`
	Id       *uint64                 `xml:"Id,omitempty"`
	SefUrl   string                  `xml:"SefUrl,omitempty"`
	FileName string                  `xml:"FileName,omitempty"`
	Alt      string                  `xml:"Alt,omitempty"`
	Import   *ProductImageImportType `xml:"Import,omitempty"`
}

type ProductImageImportType struct {
	Url     string `xml:"Url"`
	Encoded string `xml:"Encoded"`
}

type ProductVariantType struct {
	Name   string                    `xml:"Name"`
	Values []ProductVariantValueType `xml:"Values>Value,omitempty"`
}

type ProductVariantValueType struct {
	Name       string   `xml:"Name"`
	ExtraPrice *float64 `xml:"ExtraPrice,omitempty"`
}

type ProductDataType struct {
	Id    uint64 `xml:"Id"`
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type ProductParamType struct {
	Id     string `xml:"Id,omitempty"`
	Type   string `xml:"Type,omitempty"`
	Name   string `xml:"Name,omitempty"`
	Group  string `xml:"Group,omitempty"`
	Value  string `xml:"Value,omitempty"`
	Before string `xml:"Before,omitempty"`
	After  string `xml:"After,omitempty"`
}

type ProductStocksType struct {
	Status struct {
		Active  string `xml:"Active,omitempty"`
		Empty   string `xml:"Empty,omitempty"`
		Variant string `xml:"Variant,omitempty"`
	} `xml:"Status"`
	Stock []struct {
		Variant []string `xml:"Variants>Variant,omitempty"`
		Qty     *uint64  `xml:"Qty,omitempty"`
	} `xml:"Stock,omitempty"`
}

type ProductAdditionalDataType struct {
	ID      string `xml:"Id,omitempty"`
	Title   string `xml:"Title,omitempty"`
	Content string `xml:"Content,omitempty"`
}

type ProductOnlineContentType struct {
	URL      string `xml:"Url,omitempty"`
	Limit    string `xml:"Limit,omitempty"`
	Filename string `xml:"Filename,omitempty"`
}

type ProductTypesType struct {
	Type    string `xml:"Type,omitempty"`
	Parent  string `xml:"Parent,omitempty"`
	Display string `xml:"Display,omitempty"`
	Order   string `xml:"Order,omitempty"`
}

type ProductMetaType struct {
	Keywords    string `xml:"Keywords,omitempty"`
	Description string `xml:"Description,omitempty"`
	Title       string `xml:"Title,omitempty"`
}
type ProductAutomaticMetaType struct {
	Keywords    string `xml:"Keywords,omitempty"`
	Description string `xml:"Description,omitempty"`
	Title       string `xml:"Title,omitempty"`
}

type ProductQtyDiscountType struct {
	Step []struct {
		Limit struct {
			Lower string `xml:"Lower,omitempty"`
			Upper string `xml:"Upper,omitempty"`
		} `xml:"Limit,omitempty"`
		Discount string `xml:"Discount,omitempty"`
	} `xml:"Step,omitempty"`
}
type ProductSimilarProductType struct {
	ID   string `xml:"Id,omitempty"`
	Sku  string `xml:"Sku,omitempty"`
	Name string `xml:"Name,omitempty"`
}

type ProductAdditionalProductType struct {
	ID   string `xml:"Id,omitempty"`
	Sku  string `xml:"Sku,omitempty"`
	Name string `xml:"Name,omitempty"`
}
