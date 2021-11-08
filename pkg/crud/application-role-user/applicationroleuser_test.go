package applicationroleuser

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

func TestApplicationRoleUserCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationRoleUser := &npool.RoleUserInfo{
		AppID:  randstr.Hex(10),
		RoleID: uuid.New().String(),
		UserID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.SetUserRoleRequest{
		UserIDs: []string{applicationRoleUser.UserID},
		AppID:   applicationRoleUser.AppID,
		RoleID:  applicationRoleUser.RoleID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, resp.Infos[0].UserID, applicationRoleUser.UserID)
		assert.Equal(t, resp.Infos[0].AppID, applicationRoleUser.AppID)
		assert.Equal(t, resp.Infos[0].RoleID, applicationRoleUser.RoleID)
		applicationRoleUser.ID = resp.Infos[0].ID
	}

	resp1, err := GetRoleUsers(context.Background(), &npool.GetRoleUsersRequest{
		RoleID: applicationRoleUser.RoleID,
		AppID:  applicationRoleUser.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("resp1 is %v", resp1)

	_, err = Delete(context.Background(), &npool.UnSetUserRoleRequest{
		AppID:   applicationRoleUser.AppID,
		RoleID:  applicationRoleUser.RoleID,
		UserIDs: []string{applicationRoleUser.UserID},
	})
	assert.Nil(t, err)
}
