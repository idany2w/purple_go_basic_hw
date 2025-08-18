package api

import "demo/struct/src/config"

type ApiProvider struct {
	apiUrl string
	config *config.Config
}

func NewApiProvider(apiUrl string, config *config.Config) *ApiProvider {
	return &ApiProvider{
		apiUrl: apiUrl,
		config: config,
	}
}

func (a *ApiProvider) Read() ([]byte, error) {
	return nil, nil
}

func (a *ApiProvider) Save(data []byte) error {
	return nil
}
