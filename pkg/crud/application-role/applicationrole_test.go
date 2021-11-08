package applicationrole

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
	"github.com/thanhpk/randstr"
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

	applicationRole := &npool.RoleInfo{
		RoleName: "test" + uuid.New().String(),
		AppID:    randstr.Hex(10),
		Creator:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateRoleRequest{
		Request: applicationRole,
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
		Request: applicationRole,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, applicationRole.ID)
		assert.Equal(t, resp2.Info.AppID, applicationRole.AppID)
		assert.Equal(t, resp2.Info.RoleName, applicationRole.RoleName)
		assert.Equal(t, resp2.Info.Creator, applicationRole.Creator)
	}

	resp3, err := RoleNameExist(context.Background(), applicationRole)
	if assert.Nil(t, err) {
		assert.Equal(t, resp3, true)
	}

	_, err = Delete(context.Background(), &npool.DeleteRoleRequest{
		RoleID: applicationRole.ID,
		AppID:  applicationRole.AppID,
	})
	assert.Nil(t, err)
}
