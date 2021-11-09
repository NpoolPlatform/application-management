package application

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
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

func TestApplicationCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	application := &npool.ApplicationInfo{
		ApplicationName:  "test" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateApplicationRequest{
		Info: application,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, "")
		assert.NotEqual(t, resp.Info.ClientSecret, "")
		assert.Equal(t, resp.Info.ApplicationName, application.ApplicationName)
		assert.Equal(t, resp.Info.ApplicationOwner, application.ApplicationOwner)
		application.ID = resp.Info.ID
		application.ClientSecret = resp.Info.ClientSecret
	}

	resp1, err := Get(context.Background(), &npool.GetApplicationRequest{
		AppID: application.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, application.ID)
		assert.Equal(t, resp1.Info.ClientSecret, application.ClientSecret)
		assert.Equal(t, resp1.Info.ApplicationName, application.ApplicationName)
		assert.Equal(t, resp1.Info.ApplicationOwner, application.ApplicationOwner)
	}

	_, err = GetAll(context.Background(), &npool.GetApplicationsRequest{})
	assert.Nil(t, err)

	_, err = GetApplicationsByOwner(context.Background(), application.ApplicationOwner)
	assert.Nil(t, err)

	_, err = Delete(context.Background(), &npool.DeleteApplicationRequest{
		AppID: application.ID,
	})
	assert.Nil(t, err)
}
