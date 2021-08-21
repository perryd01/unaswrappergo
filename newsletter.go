package unaswrappergo

import (
	"encoding/xml"
)

type GetNewsletterRequest struct {
	Params NewsletterParams `xml:"Params"`
}

// igénylő típusa, lehetséges értékek:
// 	customer - vásárló;
//	subscriber - csak igénylő
type newsletterIdentityType string

const (
	Customer   = "customer"
	Subscriber = "subscriber"
)

//a hírlevél igénylő visszaigazolta a feliratkozását, "yes" vagy "1" érték: feliratkozás visszaigazolva; "no" vagy "0" érték: feliratkozás nincs visszaigazolva
type newsletterJoinAcceptType string

const (
	Yes = "yes"
	No  = "no"
)

type GetNewsletterResponse struct {
	Subscribers newsletterSubscriberList `xml:"Subscribers"`
}

type newsletterSubscriberList struct {
	Subscriber []newsletterSubscriber `xml:"Subscriber"`
}

type newsletterSubscriber struct {
	Email      *string `xml:"Email,omitempty"`
	Type       *string `xml:"Type,omitempty"`
	Time       *string `xml:"Time,omitempty"`
	Name       *string `xml:"Name,omitempty"`
	Address    *string `xml:"Address,omitempty"`
	Lang       *string `xml:"Lang,omitempty"`
	Authorized *string `xml:"Authorized,omitempty"`
}

type NewsletterParams struct {
	Type      newsletterIdentityType `xml:"Type"`
	Auth      string                 `xml:"Auth"`
	TimeStart *string                `xml:"TimeStart"`
	TimeEnd   *string                `xml:"TimeEnd"`
}

func (uo *UnasObject) getNewsletter(np *NewsletterParams) (*newsletterSubscriberList, error) {
	getnwr := GetNewsletterRequest{Params: *np}
	b, err := xml.Marshal(getnwr)
	if err != nil {
		return nil, err
	}
	body, err := uo.makeRequest(endpointEnumType(GetNewsletter), b)
	if err != nil {
		return nil, err
	}

	nlsl := newsletterSubscriberList{}

	err = xml.Unmarshal(body, &nlsl)
	if err != nil {
		return nil, err
	}

	return &nlsl, nil
}

type setNewsletterSubscriberList struct {
	Subscribers setNewsletterSubscribers `xml:"Subscribers"`
}

type setNewsletterSubscribers struct {
	Subscriber []setNewsletterSubscriber `xml:"Subscriber"`
}

type setNewsletterSubscriber struct {
	Action string `xml:"Action"`
	Email  string `xml:"Email"`
	Name   string `xml:"Name"`
}

//func (uo *UnasObject) setNewsletter()
