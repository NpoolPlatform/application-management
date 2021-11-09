package applicationroleuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/crud/application"
	applicationrole "github.com/NpoolPlatform/application-management/pkg/crud/application-role"
	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
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

func TestApplicationRoleUserCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationInfo := &npool.ApplicationInfo{
		ApplicationName:  "test-role-user" + uuid.New().String(),
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

	respRole, err := applicationrole.Create(context.Background(), &npool.CreateRoleRequest{
		Info: applicationRole,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respRole.Info.ID, uuid.UUID{})
		assert.Equal(t, respRole.Info.AppID, applicationRole.AppID)
		assert.Equal(t, respRole.Info.RoleName, applicationRole.RoleName)
		assert.Equal(t, respRole.Info.Creator, applicationRole.Creator)
		applicationRole.ID = respRole.Info.ID
	}

	applicationUser := &npool.ApplicationUserInfo{
		AppID:    applicationInfo.ID,
		UserID:   uuid.New().String(),
		Original: true,
	}

	respUser, err := applicationuser.Create(context.Background(), &npool.AddUsersToApplicationRequest{
		AppID:    applicationUser.AppID,
		UserIDs:  []string{applicationUser.UserID},
		Original: applicationUser.Original,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, respUser.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, respUser.Infos[0].AppID, applicationUser.AppID)
		assert.Equal(t, respUser.Infos[0].UserID, applicationUser.UserID)
		assert.Equal(t, respUser.Infos[0].Original, applicationUser.Original)
		applicationUser.ID = respUser.Infos[0].ID
	}

	applicationRoleUser := &npool.RoleUserInfo{
		AppID:  applicationInfo.ID,
		RoleID: applicationRole.ID,
		UserID: applicationUser.UserID,
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

	resp2, err := GetUserRole(context.Background(), &npool.GetUserRoleRequest{
		UserID: applicationRoleUser.UserID,
		AppID:  applicationRoleUser.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("resp2 is %v", resp2)

	_, err = Delete(context.Background(), &npool.UnSetUserRoleRequest{
		AppID:   applicationRoleUser.AppID,
		RoleID:  applicationRoleUser.RoleID,
		UserIDs: []string{applicationRoleUser.UserID},
	})
	assert.Nil(t, err)
}
