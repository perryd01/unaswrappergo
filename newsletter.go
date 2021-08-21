package unaswrappergo

type GetNewsletterRequest struct {
	Params NewsletterParams `json:"Params"`
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
	Subscribers newsletterSubscriberList `json:"Subscribers"`
}

type newsletterSubscriberList struct {
	Subscriber []Subscriber `json:"Subscriber"`
}

type newsletterSubscriber struct {
	Email      *string `json:"Email,omitempty"`
	Type       *string `json:"Type,omitempty"`
	Time       *string `json:"Time,omitempty"`
	Name       *string `json:"Name,omitempty"`
	Address    *string `json:"Address,omitempty"`
	Lang       *string `json:"Lang,omitempty"`
	Authorized *string `json:"Authorized,omitempty"`
}

type NewsletterParams struct {
	Type newsletterIdentityType `json:"Type"`
	Auth string                 `json:"Auth"`
	// Unix timestmap, optional
	TimeStart *string `json:"TimeStart"`
	// Unix timestamp, optional
	TimeEnd *string `json:"TimeEnd"`
}

func (uo *UnasObject) getNewsletter(*np NewsletterParams)(*newsletterSubscriberList, error){

}
