package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsHlsDefaultApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsDefaultApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsDefaultApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsDefaultApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsDefaultApi) Create(hlsManifestDefault model.HlsManifestDefault) (*model.HlsManifestDefault, error) {
    payload := model.HlsManifestDefault(hlsManifestDefault)
    
    err := api.apiClient.Post("/encoding/manifests/hls/default", &payload)
    return &payload, err
}
