package unaswrappergo

import (
	"encoding/xml"
	"strconv"
	"time"
)

// YYYY.MM.DD Date format
type UnasDate time.Time

func (date *UnasDate) ToTime() *time.Time {
	return (*time.Time)(date)
}

func (date UnasDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(date).Format("2006.01.02"), start)
}

func (date *UnasDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err == nil {
		*(*time.Time)(date), err = time.Parse("2006.01.02", s)
	}
	return err
}

// Unix timestamp format
type UnasTimeStamp time.Time

func (timestamp *UnasTimeStamp) ToTime() *time.Time {
	return (*time.Time)(timestamp)
}

func (timestamp UnasTimeStamp) MarhsalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(timestamp).Unix(), start)
}

func (timestamp *UnasTimeStamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err == nil {
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			*(*time.Time)(timestamp) = time.Unix(i, 0)
		}
	}
	return err
}

type statusBaseType int

const (
	StatusBaseNotActive statusBaseType = iota
	StatusBaseActive
	StatusBaseActiveNew
	StatusBaseActiveNotBuyable
)
