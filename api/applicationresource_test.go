package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/application"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationResourceAPI(t *testing.T) { // nolint
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

	resourceInfo := npool.ResourceInfo{
		ResourceName: "test-resource" + uuid.New().String(),
		AppID:        application.ID,
		Creator:      uuid.New().String(),
	}

	response1 := npool.CreateResourceResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateResourceRequest{
			Info: &resourceInfo,
		}).Post("http://localhost:50080/v1/create/resource")
	if assert.Nil(t, err) {
		fmt.Println(err, resp1)
		assert.Equal(t, 200, resp1.StatusCode())

		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response1.Info.ID, uuid.UUID{})
			assert.Equal(t, response1.Info.ResourceName, resourceInfo.ResourceName)
			assert.Equal(t, response1.Info.AppID, resourceInfo.AppID)
			resourceInfo.ID = response1.Info.ID
		}
	}

	response2 := npool.GetResourceResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetResourceRequest{
			ResourceID: resourceInfo.ID,
			AppID:      resourceInfo.AppID,
		}).Post("http://localhost:50080/v1/get/resource")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())

		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.Equal(t, response2.Info.ID, resourceInfo.ID)
			assert.Equal(t, response2.Info.ResourceName, resourceInfo.ResourceName)
			assert.Equal(t, response2.Info.AppID, resourceInfo.AppID)
		}
	}

	response6 := npool.GetResourceResponse{}
	resp6, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetResourceByNameRequest{
			ResourceName: resourceInfo.ResourceName,
			AppID:        resourceInfo.AppID,
		}).Post("http://localhost:50080/v1/get/resource/by/name")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp6.StatusCode())

		err := json.Unmarshal(resp6.Body(), &response6)
		if assert.Nil(t, err) {
			assert.Equal(t, response6.Info.ID, resourceInfo.ID)
			assert.Equal(t, response6.Info.ResourceName, resourceInfo.ResourceName)
			assert.Equal(t, response6.Info.AppID, resourceInfo.AppID)
		}
	}

	response3 := npool.GetResourcesResponse{}
	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetResourcesRequest{
			AppID: resourceInfo.AppID,
		}).Post("http://localhost:50080/v1/get/resources")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())

		err := json.Unmarshal(resp3.Body(), &response3)
		assert.Nil(t, err)
	}

	resourceInfo.ResourceDescription = "test resource api"
	response4 := npool.UpdateResourceResponse{}
	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateResourceRequest{
			Info: &resourceInfo,
		}).Post("http://localhost:50080/v1/update/resource")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())

		err := json.Unmarshal(resp4.Body(), &response4)
		if assert.Nil(t, err) {
			assert.Equal(t, response4.Info.ID, resourceInfo.ID)
			assert.Equal(t, response4.Info.ResourceName, resourceInfo.ResourceName)
			assert.Equal(t, response4.Info.AppID, resourceInfo.AppID)
			assert.Equal(t, response4.Info.ResourceDescription, resourceInfo.ResourceDescription)
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetResourceByCreatorRequest{
			AppID:   resourceInfo.AppID,
			Creator: resourceInfo.Creator,
		}).Post("http://localhost:50080/v1/get/resource/by/creator")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
	}

	resp7, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppID:      resourceInfo.AppID,
			ResourceID: resourceInfo.ID,
		}).Post("http://localhost:50080/v1/delete/resource")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp7.StatusCode())
	}
}
