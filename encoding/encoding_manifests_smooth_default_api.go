package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsSmoothDefaultApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothDefaultApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothDefaultApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothDefaultApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothDefaultApi) Create(smoothManifestDefault model.SmoothManifestDefault) (*model.SmoothManifestDefault, error) {
    payload := model.SmoothManifestDefault(smoothManifestDefault)
    
    err := api.apiClient.Post("/encoding/manifests/smooth/default", &payload)
    return &payload, err
}
