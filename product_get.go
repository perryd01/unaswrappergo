package unaswrappergo

import (
	"encoding/xml"
)

type getProductRequest struct {
	Params GetProductRequestParams `xml:"Params"`
}

type GetProductRequestParams struct {
	StatusBase   *statusBaseType `xml:"StatusBase,omitempty"`   // termék alap státusza; 0 – nem aktív; 1 – aktív; 2 – aktív, új; 3 – aktív, nem vásárolható;
	ID           string          `xml:"Id,omitempty"`           // termék egyedi azonosítója, ha ezt megadtad, akkor az Sku értéket figyelmen kívül hagyjuk
	Sku          string          `xml:"Sku,omitempty"`          // termék cikkszáma
	Parent       string          `xml:"Parent,omitempty"`       // típus összevonás esetén ezen alap típushoz tartozó termékek listázása
	TimeStart    *UnasTimeStamp  `xml:"TimeStart,omitempty"`    // unix timestamp, ezen időpont után módosult termékek listázása
	TimeEnd      *UnasTimeStamp  `xml:"TimeEnd,omitempty"`      // unix timestamp, ezen időpont előtt módosult termékek listázása
	DateStart    *UnasDate       `xml:"DateStart,omitempty"`    // YYYY.MM.DD formátum, ezen dátum után módosult termékek listázása
	DateEnd      *UnasDate       `xml:"DateEnd,omitempty"`      // YYYY.MM.DD formátum, ezen dátum előtt módosult termékek listázása
	ContentType  string          `xml:"ContentType,omitempty"`  // azt határozhatod meg, hogy milyen adatok jelenjenek meg egy termékről, négy szint közül választhatsz: minimal - gyors adatlekérés, minimális termék adatokkal; short - szűkített lista, bővített törzsadatokkal; normal - normál lista a leggyakrabban használt adatokkal (alapértelmezett); full – teljes lista minden termék adattal (csak valódi szükség esetén használandó)
	ContentParam string          `xml:"ContentParam,omitempty"` // "full" lekérés esetén vesszővel elválasztva megadhatók paraméter azonosítók, így szűkíthető a lekérendő paraméter értékek köre
	LimitStart   string          `xml:"LimitStart,omitempty"`   // Ha nem az összes terméket szeretnéd letölteni, akkor itt adhatod meg, hogy hányadik terméktől induljon a letöltés, pozitív egész szám, csak a LimitNum paraméterrel együtt használható.
	LimitNum     string          `xml:"LimitNum,omitempty"`     // Ha nem az összes terméket szeretnéd letölteni, akkor itt adhatod meg, hogy hány termék kerüljön letöltésre.
}

func (uo *UnasObject) GetProduct(p *GetProductRequestParams) ([]*Product, error) {
	obj := getProductRequest{Params: *p}
	bodyMarshaled, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(endpointEnumType(GetProduct), bodyMarshaled)
	if err != nil {
		return nil, err
	}

	var products = make([]*Product, 0)

	err = xml.Unmarshal(resp, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
