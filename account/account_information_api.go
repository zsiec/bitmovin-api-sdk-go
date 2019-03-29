package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AccountInformationApi struct {
    apiClient *common.ApiClient
}

func NewAccountInformationApi(configs ...func(*common.ApiClient)) (*AccountInformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountInformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountInformationApi) Get() (*model.AccountInformation, error) {
    var resp *model.AccountInformation
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/account/information", &resp, reqParams)
    return resp, err
}
