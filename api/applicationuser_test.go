package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationUserAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	application := npool.ApplicationInfo{
		ApplicationName:  "test-role" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	response := npool.CreateApplicationResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateApplicationRequest{
			Request: &application,
		}).Post("http://localhost:32759/v1/create/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())

		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Info.ID, uuid.UUID{})
			assert.Equal(t, response.Info.ApplicationName, application.ApplicationName)
			assert.Equal(t, response.Info.ApplicationOwner, application.ApplicationOwner)
			application.ID = response.Info.ID
		}
	}

	appUser := npool.ApplicationUserInfo{
		AppID:  application.ID,
		UserID: uuid.New().String(),
	}

	responseUser := npool.AddUsersToApplicationResponse{}
	respUser, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AddUsersToApplicationRequest{
			AppID:   appUser.AppID,
			UserIDs: []string{appUser.UserID},
		}).Post("http://localhost:32759/v1/add/users/to/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, respUser.StatusCode())
		err := json.Unmarshal(respUser.Body(), &responseUser)
		if assert.Nil(t, err) {
			assert.NotEqual(t, responseUser.Infos[0].ID, uuid.UUID{})
			assert.Equal(t, responseUser.Infos[0].UserID, appUser.UserID)
			assert.Equal(t, responseUser.Infos[0].AppID, appUser.AppID)
		}
	}
}
