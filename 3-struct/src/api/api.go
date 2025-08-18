package api

type ApiProvider struct {
	apiUrl string
}

func NewApiProvider(apiUrl string) *ApiProvider {
	return &ApiProvider{apiUrl: apiUrl}
}

func (a *ApiProvider) Read() ([]byte, error) {
	return nil, nil
}

func (a *ApiProvider) Save(data []byte) error {
	return nil
}
