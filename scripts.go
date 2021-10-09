package unaswrappergo

import "encoding/xml"

type getScriptTagRequest struct {
	Params GetScriptTagParams `xml:"Params,omitempty"`
}

type GetScriptTagParams struct {
	ID string `xml:"Id,omitempty"`
}

type ScriptTag struct {
	Action string `xml:"Action"`
	ID     string `xml:"Id"`
	Status string `xml:"Status"`
	Dates  struct {
		Creation     UnasTimeDate `xml:"Creation"`
		Modification UnasTimeDate `xml:"Modification"`
	} `xml:"Dates"`
	Pages struct {
		Page []struct {
			ID string `xml:"Id"`
		} `xml:"Page"`
	} `xml:"Pages"`
	Type    string `xml:"Type"`
	Title   string `xml:"Title"`
	Content string `xml:"Content"`
}

func (uo UnasObject) GetScriptTag(Params *GetScriptTagParams) ([]*ScriptTag, error) {
	requestObject := getScriptTagRequest{Params: *Params}
	b, err := xml.Marshal(requestObject)
	if err != nil {
		return nil, err
	}
	respBody, err := uo.makeRequest(GetScriptTag, b)
	if err != nil {
		return nil, err
	}
	scriptTags := make([]*ScriptTag, 0)
	err = xml.Unmarshal(respBody, scriptTags)
	if err != nil {
		return nil, err
	}
	return scriptTags, nil
}

// func (uo UnasObject) SetScriptTag() ([]*SetScriptTagStatus, error) {}

type SetScriptTagElement struct {
	ID     string `xml:"ScriptTag>Id,omitempty"`
	Action string `xml:"ScriptTag>Action,omitempty"`
	Status string `xml:"ScriptTag>Status,omitempty"`
}

type setScriptTagResponse struct {
	Statuses []*SetScriptTagStatus `xml:"ScriptTag"`
}

type SetScriptTagStatus struct {
	ID     string `xml:"Id"`
	Action string `xml:"Action"`
	Status string `xml:"Status"`
}
