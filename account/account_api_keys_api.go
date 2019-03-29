package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountApiKeysApi struct {
    apiClient *common.ApiClient
}

func NewAccountApiKeysApi(configs ...func(*common.ApiClient)) (*AccountApiKeysApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountApiKeysApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountApiKeysApi) Get(apiKeyId string) (*model.AccountApiKey, error) {
    var resp *model.AccountApiKey
    reqParams := func(params *common.RequestParams) {
        params.PathParams["api_key_id"] = apiKeyId
	}
    err := api.apiClient.Get("/account/api-keys/{api_key_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountApiKeysApi) List() (*pagination.AccountApiKeysListPagination, error) {
    var resp *pagination.AccountApiKeysListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/account/api-keys", &resp, reqParams)
    return resp, err
}
func (api *AccountApiKeysApi) Create() (*model.AccountApiKey, error) {
    
    var payload model.AccountApiKey
    err := api.apiClient.Post("/account/api-keys", &payload)
    return &payload, err
}
func (api *AccountApiKeysApi) Delete(apiKeyId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["api_key_id"] = apiKeyId
	}
    err := api.apiClient.Delete("/account/api-keys/{api_key_id}", &resp, reqParams)
    return resp, err
}
