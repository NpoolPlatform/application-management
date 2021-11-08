package applicationgroup

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

func TestApplicationGroupCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationGroup := &npool.GroupInfo{
		AppID:      randstr.Hex(10),
		GroupName:  "test" + uuid.New().String(),
		GroupOwner: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateGroupRequest{
		Request: applicationGroup,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp.Info.GroupOwner, applicationGroup.GroupOwner)
		applicationGroup.ID = resp.Info.ID
	}

	applicationGroup.GroupName = "test-update" + uuid.New().String()
	resp1, err := Update(context.Background(), &npool.UpdateGroupRequest{
		Request: applicationGroup,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, applicationGroup.ID)
		assert.Equal(t, resp1.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp1.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp1.Info.GroupOwner, applicationGroup.GroupOwner)
	}

	resp2, err := Get(context.Background(), &npool.GetGroupRequest{
		GroupID: applicationGroup.ID,
		AppID:   applicationGroup.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, applicationGroup.ID)
		assert.Equal(t, resp2.Info.AppID, applicationGroup.AppID)
		assert.Equal(t, resp2.Info.GroupName, applicationGroup.GroupName)
		assert.Equal(t, resp2.Info.GroupOwner, applicationGroup.GroupOwner)
	}

	resp3, err := GetAll(context.Background(), &npool.GetAllGroupsRequest{
		AppID: applicationGroup.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get all group is resp3: %v", resp3)

	_, err = Delete(context.Background(), &npool.DeleteGroupRequest{
		GroupID: applicationGroup.ID,
		AppID:   applicationGroup.AppID,
	})
	assert.Nil(t, err)
}
