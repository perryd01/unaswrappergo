package unaswrappergo

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"sync"
)

type getProductRequestResponse struct {
	XMLName  xml.Name   `xml:"Products"`
	Products []*Product `xml:"Product"`
}

type GetProductRequestParams struct {
	XMLName      xml.Name         `xml:"Params"`
	StatusBase   statusBaseEnum   `xml:"StatusBase,omitempty"`   // termék alap státusza; 0 – nem aktív; 1 – aktív; 2 – aktív, új; 3 – aktív, nem vásárolható;
	ID           string           `xml:"Id,omitempty"`           // termék egyedi azonosítója, ha ezt megadtad, akkor az Sku értéket figyelmen kívül hagyjuk
	Sku          string           `xml:"Sku,omitempty"`          // termék cikkszáma
	Parent       string           `xml:"Parent,omitempty"`       // típus összevonás esetén ezen alap típushoz tartozó termékek listázása
	TimeStart    *UnasTimeStamp   `xml:"TimeStart,omitempty"`    // unix timestamp, ezen időpont után módosult termékek listázása
	TimeEnd      *UnasTimeStamp   `xml:"TimeEnd,omitempty"`      // unix timestamp, ezen időpont előtt módosult termékek listázása
	DateStart    *UnasDate        `xml:"DateStart,omitempty"`    // YYYY.MM.DD formátum, ezen dátum után módosult termékek listázása
	DateEnd      *UnasDate        `xml:"DateEnd,omitempty"`      // YYYY.MM.DD formátum, ezen dátum előtt módosult termékek listázása
	ContentType  string           `xml:"ContentType,omitempty"`  // azt határozhatod meg, hogy milyen adatok jelenjenek meg egy termékről, négy szint közül választhatsz: minimal - gyors adatlekérés, minimális termék adatokkal; short - szűkített lista, bővített törzsadatokkal; normal - normál lista a leggyakrabban használt adatokkal (alapértelmezett); full – teljes lista minden termék adattal (csak valódi szükség esetén használandó)
	ContentParam ContentParamList `xml:"ContentParam,omitempty"` // "full" lekérés esetén vesszővel elválasztva megadhatók paraméter azonosítók, így szűkíthető a lekérendő paraméter értékek köre
	LimitStart   string           `xml:"LimitStart,omitempty"`   // Ha nem az összes terméket szeretnéd letölteni, akkor itt adhatod meg, hogy hányadik terméktől induljon a letöltés, pozitív egész szám, csak a LimitNum paraméterrel együtt használható.
	LimitNum     string           `xml:"LimitNum,omitempty"`     // Ha nem az összes terméket szeretnéd letölteni, akkor itt adhatod meg, hogy hány termék kerüljön letöltésre.
}

func (uo *UnasObject) GetProduct(p *GetProductRequestParams) ([]*Product, error) {
	bodyMarshaled, err := xml.Marshal(p)
	if err != nil {
		return nil, err
	}
	resp, err := uo.makeRequest(GetProduct, bodyMarshaled)
	if err != nil {
		return nil, err
	}

	var products = getProductRequestResponse{}

	err = xml.Unmarshal(resp, &products)
	if err != nil {
		return nil, err
	}
	return products.Products, nil
}

func (uo *UnasObject) GetProductParallel(p *GetProductRequestParams, chunkSize uint32, allProductNumber uint32) ([]*Product, error) {
	numberOfChunks := allProductNumber / chunkSize
	if allProductNumber%chunkSize != 0 {
		numberOfChunks++
	}
	if numberOfChunks > 30 {
		fmt.Println("You might be exceeding Unas API limit")
	}

	ch := make(chan string) //p.LimitNum = strconv.Itoa(int(chunkSize))

	var wg sync.WaitGroup
	products := make([]*Product, 0)

	var parameters = make([]*GetProductRequestParams, 0)
	for i := 0; i < int(numberOfChunks); i++ {
		t := GetProductRequestParams{
			XMLName:      p.XMLName,
			StatusBase:   p.StatusBase,
			ID:           p.ID,
			Sku:          p.Sku,
			Parent:       p.Parent,
			TimeStart:    p.TimeStart,
			TimeEnd:      p.TimeEnd,
			DateStart:    p.DateStart,
			DateEnd:      p.DateEnd,
			ContentType:  p.ContentType,
			ContentParam: p.ContentParam,
			LimitNum:     strconv.Itoa(int(chunkSize)),
		}
		t.LimitStart = strconv.Itoa(i*int(chunkSize) + 1)
		parameters = append(parameters, &t)
	}

	for _, param := range parameters {
		bodyMarshaled, err := xml.Marshal(param)
		if err != nil {
			return nil, err
		}
		wg.Add(1)
		go uo.makeRequestParallel(GetProduct, bodyMarshaled, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		temp := getProductRequestResponse{}
		err := xml.Unmarshal([]byte(res), &temp)
		if err != nil {
			return nil, err
		}
		products = append(products, temp.Products...)
	}
	return products, nil
}
