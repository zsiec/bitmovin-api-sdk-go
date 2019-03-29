package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsAsperaApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsAsperaCustomdataApi
}

func NewEncodingInputsAsperaApi(configs ...func(*common.ApiClient)) (*EncodingInputsAsperaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsAsperaApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsAsperaCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsAsperaApi) Delete(inputId string) (*model.AsperaInput, error) {
    var resp *model.AsperaInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/aspera/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsAsperaApi) List(queryParams ...func(*query.AsperaInputListQueryParams)) (*pagination.AsperaInputsListPagination, error) {
    queryParameters := &query.AsperaInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AsperaInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/aspera", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsAsperaApi) Get(inputId string) (*model.AsperaInput, error) {
    var resp *model.AsperaInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/aspera/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsAsperaApi) Create(asperaInput model.AsperaInput) (*model.AsperaInput, error) {
    payload := model.AsperaInput(asperaInput)
    
    err := api.apiClient.Post("/encoding/inputs/aspera", &payload)
    return &payload, err
}
