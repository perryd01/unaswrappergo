package unaswrappergo

import "encoding/xml"

type Product struct {
	XMLName             xml.Name                   `xml:"Product"`
	Action              string                     `xml:"Action,omitempty"`
	State               string                     `xml:"State"`
	Id                  string                     `xml:"ID"`
	Sku                 string                     `xml:"Sku"`
	SkuNew              string                     `xml:"SkuNew,omitempty"`
	History             []ProductEventType         `xml:"History>Event,omitempty"`
	Statuses            []ProductStatusType        `xml:"Statuses>Status,omitempty"`
	NoList              bool                       `xml:"NoList,omitempty"`
	Inquire             bool                       `xml:"Inquire,omitempty"`
	CustDiscountDisable bool                       `xml:"CustDiscountDisable,omitempty"`
	Explicit            bool                       `xml:"Explicit,omitempty"`
	Export              *ProductExportType         `xml:"Export,omitempty"`
	PublicInterval      *ProductPublicIntervalType `xml:"PublicInterval,omitempty"`
	Name                string                     `xml:"Name,omitempty"`
	Unit                string                     `xml:"Unit,omitempty"`
	MinimumQty          string                     `xml:"MinimumQty,omitempty"`
	MaximumQty          string                     `xml:"MaximumQty,omitempty"`
	AlertQty            *uint64                    `xml:"AlertQty,omitempty"`
	UnitStep            string                     `xml:"UnitStep,omitempty"`
	AlterUnit           *ProductAlterUnitType      `xml:"AlterUnit,omitempty"`
	Weight              *float64                   `xml:"Weight,omitempty"`
	Point               *uint64                    `xml:"Point,omitempty"`
	BuyableWithPoint    *string                    `xml:"BuyableWithPoint,omitempty"`
	Description         *ProductDescriptionType    `xml:"Description,omitempty"`
	Prices              []ProductPriceType         `xml:"Prices>Price,omitempty"`
	Categories          []ProductCategoryType      `xml:"Categories>Category,omitempty"`
	Url                 *string                    `xml:"Url,omitempty"`
	SefUrl              *string                    `xml:"SefUrl,omitempty"`
	Images              *ProductImagesType         `xml:"Images,omitempty"`
	Variants            []ProductVariantType       `xml:"Variants>Variant,omitempty"`
	Datas               []ProductDataType          `xml:"Datas>Data,omitempty"`
	Params              []ProductParamType         `xml:"Params>Param,omitempty"`
	// AdditionalDatas
	// QtyDsicount
	// AdditionalProducts
	// SimilarProducts
	// Stocks
	// OnlineContent
	// Types
	// PackageProduct
	// Meta
	// AutoMaticMeta
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
	Short string `xml:"Short"`
	Long  string `xml:"Long"`
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
