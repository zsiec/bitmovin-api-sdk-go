package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerChannelsApi struct {
    apiClient *common.ApiClient
    Versions *PlayerChannelsVersionsApi
}

func NewPlayerChannelsApi(configs ...func(*common.ApiClient)) (*PlayerChannelsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerChannelsApi{apiClient: apiClient}

    versionsApi, err := NewPlayerChannelsVersionsApi(configs...)
    api.Versions = versionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerChannelsApi) List() (*pagination.PlayerChannelsListPagination, error) {
    var resp *pagination.PlayerChannelsListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/player/channels", &resp, reqParams)
    return resp, err
}
