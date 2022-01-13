package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/application"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	application := npool.ApplicationInfo{
		ApplicationName:  "test-app" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	response := npool.CreateApplicationResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateApplicationRequest{
			Info: &application,
		}).Post("http://localhost:50080/v1/create/app")
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

	response1 := npool.UpdateApplicationResponse{}
	application.ApplicationLogo = "test"
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateApplicationRequest{
			Info: &application,
		}).Post("http://localhost:50080/v1/update/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.Equal(t, response1.Info.ID, application.ID)
			assert.Equal(t, response1.Info.ApplicationName, application.ApplicationName)
			assert.Equal(t, response1.Info.ApplicationOwner, application.ApplicationOwner)
			assert.Equal(t, response1.Info.ApplicationLogo, application.ApplicationLogo)
		}
	}

	response2 := npool.GetApplicationResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetApplicationRequest{
			AppID: application.ID,
		}).Post("http://localhost:50080/v1/get/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.Equal(t, response2.Info.ID, application.ID)
			assert.Equal(t, response2.Info.ApplicationName, application.ApplicationName)
			assert.Equal(t, response2.Info.ApplicationOwner, application.ApplicationOwner)
			assert.Equal(t, response2.Info.ApplicationLogo, application.ApplicationLogo)
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetApplicationsRequest{}).
		Post("http://localhost:50080/v1/get/apps")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetApplicationByOwnerRequest{
			Owner: application.ApplicationOwner,
		}).
		Post("http://localhost:50080/v1/get/app/by/owner")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteApplicationRequest{
			AppID: application.ID,
		}).Post("http://localhost:50080/v1/delete/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
	}
}
