package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi) GetCustomData(encodingId string, muxingId string, captionsId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/webvtt/{captions_id}/customData", &resp, reqParams)
    return resp, err
}