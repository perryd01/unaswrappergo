package unaswrappergo

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestUnasObject_GetStorage(t *testing.T) {
	var apiKey = os.Getenv("apiKey")
	var unasObject, err = AuthwithAPIKey(apiKey)
	if err != nil {
		t.Error(err)
	}

	params := GetStorageParams{
		Type:    StorageTypeAll,
		GetInfo: StorageGetInfoYes,
		Folder:  "",
	}
	storageItems, err := unasObject.GetStorage(&params)
	if err != nil {
		t.Error(err)
	}
	if len(storageItems) != 0 {
		if storageItems[0].Name == "" {
			t.Error("empty storageitem element name")
		}
	}
}

func TestUnasObject_SetStorage(t *testing.T) {
	var apiKey = os.Getenv("apiKey")
	var unasObject, err = AuthwithAPIKey(apiKey)
	if err != nil {
		t.Error(err)
	}

	randomNumber := rand.Intn(1000000)
	textAsFolder := strconv.Itoa(randomNumber)

	item := SetStorageItem{
		Action: UploadFile,
		Url:    "https://pastebin.com/raw/tDgYwTj7",
		Name:   textAsFolder + ".txt",
	}

	params := []*SetStorageItem{&item}

	responseItems, err := unasObject.SetStorage(params)
	if err != nil {
		t.Error(err)
	}

	if responseItems[0].Status != "ok" {
		t.Error("setStorage response not OK")
	}
}
