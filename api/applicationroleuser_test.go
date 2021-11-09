package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApplicationRoleUserAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	application := npool.ApplicationInfo{
		ApplicationName:  "test-role-user" + uuid.New().String(),
		ApplicationOwner: uuid.New().String(),
	}

	response := npool.CreateApplicationResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateApplicationRequest{
			Info: &application,
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

	roleInfo := npool.RoleInfo{
		RoleName: "test-role-user" + uuid.New().String(),
		Creator:  uuid.New().String(),
		AppID:    application.ID,
	}

	response1 := npool.CreateRoleResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateRoleRequest{
			Info: &roleInfo,
		}).Post("http://localhost:32759/v1/create/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response1.Info.ID, uuid.UUID{})
			assert.Equal(t, response1.Info.RoleName, roleInfo.RoleName)
			assert.Equal(t, response1.Info.AppID, roleInfo.AppID)
			assert.Equal(t, response1.Info.Creator, roleInfo.Creator)
			roleInfo.ID = response1.Info.ID
		}
	}

	appUserInfo := npool.ApplicationUserInfo{
		AppID:  roleInfo.AppID,
		UserID: uuid.New().String(),
	}

	responseUser := npool.AddUsersToApplicationResponse{}
	respUser, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AddUsersToApplicationRequest{
			AppID:   appUserInfo.AppID,
			UserIDs: []string{appUserInfo.UserID},
		}).Post("http://localhost:32759/v1/add/users/to/app")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, respUser.StatusCode())
		err := json.Unmarshal(respUser.Body(), &responseUser)
		if assert.Nil(t, err) {
			assert.NotEqual(t, responseUser.Infos[0].ID, uuid.UUID{})
			assert.Equal(t, responseUser.Infos[0].UserID, appUserInfo.UserID)
			assert.Equal(t, responseUser.Infos[0].AppID, appUserInfo.AppID)
		}
	}

	response2 := npool.SetUserRoleResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetUserRoleRequest{
			UserIDs: []string{appUserInfo.UserID},
			AppID:   appUserInfo.AppID,
			RoleID:  roleInfo.ID,
		}).Post("http://localhost:32759/v1/set/user/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response2.Infos[0].ID, uuid.UUID{})
			assert.Equal(t, response2.Infos[0].AppID, appUserInfo.AppID)
			assert.Equal(t, response2.Infos[0].UserID, appUserInfo.UserID)
			assert.Equal(t, response2.Infos[0].RoleID, roleInfo.ID)
		}
	}

	response3 := npool.GetUserRoleResponse{}
	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetUserRoleRequest{
			UserID: appUserInfo.UserID,
			AppID:  appUserInfo.AppID,
		}).Post("http://localhost:32759/v1/get/user/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		err := json.Unmarshal(resp3.Body(), &response3)
		if assert.Nil(t, err) {
			assert.Equal(t, response3.Infos[0].AppID, appUserInfo.AppID)
			assert.Equal(t, response3.Infos[0].ID, roleInfo.ID)
		}
	}

	response4 := npool.GetRoleUsersResponse{}
	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRoleUsersRequest{
			AppID:  appUserInfo.AppID,
			RoleID: roleInfo.ID,
		}).Post("http://localhost:32759/v1/get/role/users")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		err := json.Unmarshal(resp4.Body(), &response4)
		if assert.Nil(t, err) {
			assert.Equal(t, response4.Infos[0].AppID, appUserInfo.AppID)
			assert.Equal(t, response4.Infos[0].RoleID, roleInfo.ID)
			assert.Equal(t, response4.Infos[0].UserID, appUserInfo.UserID)
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnSetUserRoleRequest{
			UserIDs: []string{appUserInfo.UserID},
			AppID:   appUserInfo.AppID,
			RoleID:  roleInfo.ID,
		}).Post("http://localhost:32759/v1/unset/user/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		fmt.Printf("unset user role resp5 is %v", resp5)
	}
}
