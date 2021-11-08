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

func TestApplicationRoleAPI(t *testing.T) { // nolint
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

	roleInfo := npool.RoleInfo{
		RoleName: "test-role" + uuid.New().String(),
		Creator:  uuid.New().String(),
		AppID:    application.ID,
	}

	response1 := npool.CreateRoleResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateRoleRequest{
			Request: &roleInfo,
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

	roleInfo.Annotation = "test role"
	response2 := npool.UpdateRoleResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateRoleRequest{
			Request: &roleInfo,
		}).Post("http://localhost:32759/v1/update/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.Equal(t, response2.Info.ID, roleInfo.ID)
			assert.Equal(t, response2.Info.RoleName, roleInfo.RoleName)
			assert.Equal(t, response2.Info.AppID, roleInfo.AppID)
			assert.Equal(t, response2.Info.Creator, roleInfo.Creator)
		}
	}

	response3 := npool.GetRoleResponse{}
	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRoleRequest{
			AppID:  roleInfo.AppID,
			RoleID: roleInfo.ID,
		}).Post("http://localhost:32759/v1/get/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		err := json.Unmarshal(resp3.Body(), &response3)
		if assert.Nil(t, err) {
			assert.Equal(t, response3.Info.ID, roleInfo.ID)
			assert.Equal(t, response3.Info.RoleName, roleInfo.RoleName)
			assert.Equal(t, response3.Info.AppID, roleInfo.AppID)
			assert.Equal(t, response3.Info.Creator, roleInfo.Creator)
		}
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRolesRequest{
			AppID: roleInfo.AppID,
		}).Post("http://localhost:32759/v1/get/roles")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteRoleRequest{
			AppID:  roleInfo.AppID,
			RoleID: roleInfo.ID,
		}).Post("http://localhost:32759/v1/delete/role")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
	}
}
