package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsTypeApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsTypeApi) Get(configurationId string) (*model.CodecConfigTypeResponse, error) {
    var resp *model.CodecConfigTypeResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/{configuration_id}/type", &resp, reqParams)
    return resp, err
}
