package unaswrappergo

import "encoding/xml"

// GetNewsletterParams Params to set for querying newsletter subscribers
// https://unas.hu/tudastar/api/newsletter#getnewsletter-keres
type GetNewsletterParams struct {
	Type      string        `xml:"Params>Type,omitempty"`
	Auth      string        `xml:"Params>Auth,omitempty"` //TODO bool
	TimeStart UnasTimeStamp `xml:"Params>TimeStart,omitempty"`
	TimeEnd   UnasTimeStamp `xml:"Params>TimeEnd,omitempty"`
}

type getNewsletterResponse struct {
	Subscribers []*NewsletterSubscriber `xml:"Subscribers"`
}

// NewsletterSubscriber One subscriber of the newsletter
type NewsletterSubscriber struct {
	Email      string        `xml:"Email,omitempty"`
	Type       string        `xml:"Type,omitempty"`
	Time       UnasTimeStamp `xml:"Time,omitempty"`
	Name       string        `xml:"Name,omitempty"`
	Address    string        `xml:"Address,omitempty"`
	Lang       string        `xml:"Lang,omitempty"`
	Authorized string        `xml:"Authorized,omitempty"` //TODO bool
}

type SetNewsletterSubscriberParams struct {
	Action string `xml:"Action,omitempty"`
	Email string `xml:"Email,omitempty"`
	Name string `xml:"Name,omitempty"`
}

type SetNewsletterSubscriberStatus struct {
	Action string `xml:"Action,omitempty"`
	Email string `xml:"Email,omitempty"`
	Name string `xml:"Name,omitempty"`
	Status string `xml:"Status,omitempty"`
}

// SetNewsletterRequest
// https://unas.hu/tudastar/api/newsletter#setnewsletter-keres
type setNewsletterRequest struct {
	Subscribers []*SetNewsletterSubscriberParams `xml:"Subscribers"`
}
// SetNewsletterResponse
// https://unas.hu/tudastar/api/newsletter#setnewsletter-valasz
type setNewsletterResponse struct {
	Subscribers []*SetNewsletterSubscriberStatus `xml:"Subscribers"`
}

// GetNewsletter Queries newsletter subscribers of the webshop
// https://unas.hu/tudastar/api/newsletter#getnewsletter-funkcio
func (uo UnasObject) GetNewsletter(params *GetNewsletterParams) ([]*NewsletterSubscriber, error) {
	body, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := uo.makeRequest(GetNewsletter, body)
	if err != nil {
		return nil, err
	}

	gnlresp := getNewsletterResponse{}

	err = xml.Unmarshal(resp, &gnlresp)
	if err != nil {
		return nil, err
	}

	return gnlresp.Subscribers, nil
}

// SetNewsletter Add/Modify/Delete data of newsletter subscribers
// https://unas.hu/tudastar/api/newsletter#setnewsletter-funkcio
func (uo UnasObject) SetNewsletter(subscribers []*SetNewsletterSubscriberParams) ([]*SetNewsletterSubscriberStatus, error) {
	snreq := setNewsletterRequest{Subscribers: subscribers}

	b, err := xml.Marshal(snreq)
	if err != nil {
		return nil, err
	}

	r, err := uo.makeRequest(SetNewsletter, b)
	if err != nil {
		return nil, err
	}

	snresp := setNewsletterResponse{}

	err = xml.Unmarshal(r, snresp)
	if err != nil {
		return nil, err
	}

	return snresp.Subscribers, nil
}