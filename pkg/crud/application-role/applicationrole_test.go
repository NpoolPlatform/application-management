package applicationrole

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

func TestApplicationRoleCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-role" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	respApp, err := application.Create(context.Background(), &npool.CreateApplicationRequest{
		Info: applicationInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respApp.Info.ID, "")
		assert.NotEqual(t, respApp.Info.ClientSecret, "")
		assert.Equal(t, respApp.Info.ApplicationName, applicationInfo.ApplicationName)
		assert.Equal(t, respApp.Info.ApplicationOwner, applicationInfo.ApplicationOwner)
		applicationInfo.ID = respApp.Info.ID
	}

	applicationRole := &npool.RoleInfo{
		RoleName: "test" + uuid.New().String(),
		AppID:    applicationInfo.ID,
		Creator:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateRoleRequest{
		Info: applicationRole,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.AppID, applicationRole.AppID)
		assert.Equal(t, resp.Info.RoleName, applicationRole.RoleName)
		assert.Equal(t, resp.Info.Creator, applicationRole.Creator)
		applicationRole.ID = resp.Info.ID
	}

	resp1, err := Get(context.Background(), &npool.GetRoleRequest{
		RoleID: applicationRole.ID,
		AppID:  applicationRole.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, applicationRole.ID)
		assert.Equal(t, resp1.Info.AppID, applicationRole.AppID)
		assert.Equal(t, resp1.Info.RoleName, applicationRole.RoleName)
		assert.Equal(t, resp1.Info.Creator, applicationRole.Creator)
	}

	_, err = GetAll(context.Background(), &npool.GetRolesRequest{
		AppID: applicationRole.AppID,
	})
	assert.Nil(t, err)

	applicationRole.RoleName = "test" + uuid.New().String()
	resp2, err := Update(context.Background(), &npool.UpdateRoleRequest{
		Info: applicationRole,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, applicationRole.ID)
		assert.Equal(t, resp2.Info.AppID, applicationRole.AppID)
		assert.Equal(t, resp2.Info.RoleName, applicationRole.RoleName)
		assert.Equal(t, resp2.Info.Creator, applicationRole.Creator)
	}

	resp3, err := GetRoleByCreator(context.Background(), &npool.GetRoleByCreatorRequest{
		AppID:   applicationRole.AppID,
		Creator: applicationRole.Creator,
	})
	assert.Nil(t, err)
	fmt.Printf("get role by creator: %v", resp3)

	_, err = GetRoleByCreator(context.Background(), &npool.GetRoleByCreatorRequest{
		AppID:   uuid.New().String(),
		Creator: uuid.New().String(),
	})
	assert.NotNil(t, err)

	_, err = Delete(context.Background(), &npool.DeleteRoleRequest{
		RoleID: applicationRole.ID,
		AppID:  applicationRole.AppID,
	})
	assert.Nil(t, err)
}
