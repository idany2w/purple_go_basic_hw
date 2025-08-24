package api

import (
	"bytes"
	"demo/go-json/bins"
	"demo/go-json/config"
	"demo/go-json/files"
	"encoding/json"
	"io"
	"net/http"
)

type Api struct {
	config config.Config
}

type CreateBinResponse struct {
	MetaData bins.Bin `json:"metadata"`
}

const apiUrl = "https://api.jsonbin.io/v3/b"

func NewApi(config *config.Config) *Api {
	if config == nil {
		panic("config cannot be nil")
	}

	return &Api{
		config: *config,
	}
}

func (api *Api) CreateBin(fromFile, name string) *bins.Bin {
	jsonDb := files.NewJsonDb(fromFile)
	data, err := jsonDb.Read()

	if err != nil {
		panic(err)
	}

	client, req := api.getHttpClient("POST", apiUrl, data)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var createBinResponse CreateBinResponse
	json.Unmarshal(body, &createBinResponse)

	bin := createBinResponse.MetaData
	bin.Name = name

	return &bin
}

func (api *Api) UpdateBin(fromFile, id string) bool {
	jsonDb := files.NewJsonDb(fromFile)
	data, err := jsonDb.Read()

	if err != nil {
		panic(err)
	}

	client, req := api.getHttpClient("PUT", apiUrl+"/"+id, data)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func (api *Api) GetBin(id string) *bins.Bin {
	client, req := api.getHttpClient("GET", apiUrl+"/"+id, nil)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var createBinResponse CreateBinResponse
	json.Unmarshal(body, &createBinResponse)

	return &createBinResponse.MetaData
}

func (api *Api) getHttpClient(method, url string, data []byte) (*http.Client, *http.Request) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", api.config.XMasterKey)

	return &http.Client{}, req
}

func (api *Api) DeleteBin(id string) bool {
	client, req := api.getHttpClient("DELETE", apiUrl+"/"+id, nil)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}
