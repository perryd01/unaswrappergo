package unaswrappergo

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// UnasDate YYYY.MM.DD Date format
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

// UnasTimeStamp Unix timestamp format
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

// Base type of the Product
type statusBaseEnum int

func (statusBase statusBaseEnum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(statusBase), start)
}

func (statusBase *statusBaseEnum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err == nil {
		*statusBase = statusBaseEnum(i)
	}
	return err
}

type ContentParamList []string

func (contentParam ContentParamList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(contentParam, ","), start)
}

func (contentParam *ContentParamList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err == nil {
		*contentParam = strings.Split(s, ",")
	}
	return err
}
