package applicationresource

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	testinit "github.com/NpoolPlatform/application-management/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestApplicationResourceCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-resource" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	respApp, err := application.Create(context.Background(), &npool.CreateApplicationRequest{
		Request: applicationInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respApp.Info.ID, "")
		assert.NotEqual(t, respApp.Info.ClientSecret, "")
		assert.Equal(t, respApp.Info.ApplicationName, applicationInfo.ApplicationName)
		assert.Equal(t, respApp.Info.ApplicationOwner, applicationInfo.ApplicationOwner)
		applicationInfo.ID = respApp.Info.ID
	}

	resource := &npool.ResourceInfo{
		AppID:        applicationInfo.ID,
		ResourceName: "test" + uuid.New().String(),
		Creator:      uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateResourceRequest{
		Request: resource,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.AppID, resource.AppID)
		assert.Equal(t, resp.Info.ResourceName, resource.ResourceName)
		assert.Equal(t, resp.Info.Creator, resource.Creator)
		resource.ID = resp.Info.ID
	}

	resp1, err := Get(context.Background(), &npool.GetResourceRequest{
		ResourceID: resource.ID,
		AppID:      resource.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resource.ID)
		assert.Equal(t, resp1.Info.AppID, resource.AppID)
		assert.Equal(t, resp1.Info.ResourceName, resource.ResourceName)
		assert.Equal(t, resp1.Info.Creator, resource.Creator)
	}

	resp2, err := GetAll(context.Background(), &npool.GetResourcesRequest{
		AppID: resource.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get all resource resp2 is: %v", resp2)

	resource.ResourceName = "test-update" + uuid.New().String()
	resp3, err := Update(context.Background(), &npool.UpdateResourceRequest{
		Request: resource,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resource.ID)
		assert.Equal(t, resp3.Info.AppID, resource.AppID)
		assert.Equal(t, resp3.Info.ResourceName, resource.ResourceName)
		assert.Equal(t, resp3.Info.Creator, resource.Creator)
	}

	_, err = Delete(context.Background(), &npool.DeleteResourceRequest{
		ResourceID: resource.ID,
		AppID:      resource.AppID,
	})
	assert.Nil(t, err)
}
