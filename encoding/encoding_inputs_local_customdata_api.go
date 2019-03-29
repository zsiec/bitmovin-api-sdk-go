package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsLocalCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsLocalCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsLocalCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsLocalCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsLocalCustomdataApi) GetCustomData(inputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/local/{input_id}/customData", &resp, reqParams)
    return resp, err
}
