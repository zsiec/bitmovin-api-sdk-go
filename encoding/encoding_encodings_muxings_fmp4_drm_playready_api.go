package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4DrmPlayreadyApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmPlayreadyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmPlayreadyApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyApi) Get(encodingId string, muxingId string, drmId string) (*model.PlayReadyDrm, error) {
    var resp *model.PlayReadyDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyApi) List(encodingId string, muxingId string, queryParams ...func(*query.PlayReadyDrmListQueryParams)) (*pagination.PlayReadyDrmsListPagination, error) {
    queryParameters := &query.PlayReadyDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.PlayReadyDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyApi) Create(encodingId string, muxingId string, playReadyDrm model.PlayReadyDrm) (*model.PlayReadyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.PlayReadyDrm(playReadyDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready", &payload, reqParams)
    return &payload, err
}
