package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsEmailsEncodingEncodingsErrorApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsEmailsEncodingEncodingsErrorApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsEncodingEncodingsErrorApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsEncodingEncodingsErrorApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsEncodingEncodingsErrorApi) Update(notificationId string, emailNotification *model.EmailNotification) (*model.EmailNotification, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }
    payload := model.EmailNotification(*emailNotification)
    
    err := api.apiClient.Put("/notifications/emails/encoding/encodings/error/{notification_id}", &payload, reqParams)
    return &payload, err
}
func (api *NotificationsEmailsEncodingEncodingsErrorApi) CreateByEncodingId(encodingId string, emailNotification model.EmailNotification) (*model.EmailNotification, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.EmailNotification(emailNotification)
    
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/{encoding_id}/error", &payload, reqParams)
    return &payload, err
}
