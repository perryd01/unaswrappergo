package unaswrappergo

import "encoding/xml"

type getStorageType string

const (
	StorageTypeAll    = "all"
	StorageTypeFile   = "file"
	StorageTypeFolder = "folder"
)

type getStorageGetInfo string

const (
	StorageGetInfoNo   = "no"
	StorageGetInfoYes  = "yes"
	StorageGetInfoOnly = "only"
)

type GetStorageParams struct {
	Type    getStorageType    `xml:"Params>Type,omitempty"`
	GetInfo getStorageGetInfo `xml:"Params>GetInfo,omitempty"`
	Folder  string            `xml:"Params>Folder,omitempty"`
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

type setStorageItemAction string

const (
	CreateFolder setStorageItemAction = "create_folder"
	DeleteFolder setStorageItemAction = "create_folder"
	UploadFile   setStorageItemAction = "upload_file"
	DeleteFile   setStorageItemAction = "delete_file"
)

type setStorageItemRequest struct {
	XMLName     xml.Name         `xml:"StorageItems"`
	StorageItem []SetStorageItem `xml:"StorageItem"`
}

type SetStorageItem struct {
	Action setStorageItemAction `xml:"Action,omitempty"`
	Folder string               `xml:"Folder,omitempty"`
	Url    string               `xml:"Url,omitempty"`
	Name   string               `xml:"Name,omitempty"`
}

type setStorageResponse struct {
	XMLName     xml.Name                 `xml:"StorageItems"`
	StorageItem []SetStorageResponseItem `xml:"StorageItem"`
}

type SetStorageResponseItem struct {
	Action string `xml:"Action"`
	Status string `xml:"Status"`
}

//SetStorage https://unas.hu/tudastar/api/storage#setstorage-funkcio
func (uo *UnasObject) SetStorage(items []SetStorageItem) ([]SetStorageResponseItem, error) {
	b, err := xml.Marshal(setStorageItemRequest{StorageItem: items})
	if err != nil {
		return nil, err
	}
	respBody, err := uo.makeRequest(SetStorage, b)
	if err != nil {
		return nil, err
	}

	sStorageResponse := setStorageResponse{}
	err = xml.Unmarshal(respBody, &sStorageResponse)
	if err != nil {
		return nil, err
	}
	return sStorageResponse.StorageItem, nil
}
