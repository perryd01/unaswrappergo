package unaswrappergo

import "encoding/xml"

type GetStorageParams struct {
	Type    string `xml:"Params>Type,omitempty"`    // TODO enum
	GetInfo string `xml:"Params>GetInfo,omitempty"` //TODO enum
	Folder  string `xml:"Params>Folder,omitempty"`
}

type getStorageResponse struct {
	StorageItems []*StorageItem `xml:"StorageItems"`
}

type StorageItem struct {
	Type        string       `xml:"Type,omitempty"`
	Name        string       `xml:"Name,omitempty"`
	ActualSpace string       `xml:"ActualSpace,omitempty"` //TODO type
	SpaceLimit  string       `xml:"SpaceLimit,omitempty"`  //TODO type
	ModifyDate  UnasTimeDate `xml:"ModifyDate,omitempty"`
	FileType    string       `xml:"FileType,omitempty"` //TODO file extension type?
	Path        string       `xml:"Path,omitempty"`
}

func (uo *UnasObject) GetStorage(params *GetStorageParams) ([]*StorageItem, error) {
	b, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}
	respBody, err := uo.makeRequest(GetStorage, b)
	if err != nil {
		return nil, err
	}

	gStorageResponse := getStorageResponse{}

	err = xml.Unmarshal(respBody, &gStorageResponse)
	if err != nil {
		return nil, err
	}

	return gStorageResponse.StorageItems, nil
}
