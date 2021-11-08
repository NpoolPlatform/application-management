package applicationgroupuser

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

func TestApplicationGroupUserCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	groupUser := &npool.GroupUserInfo{
		AppID:   randstr.Hex(10),
		GroupID: uuid.New().String(),
		UserID:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.AddGroupUsersRequest{
		UserIDs: []string{groupUser.UserID},
		AppID:   groupUser.AppID,
		GroupID: groupUser.GroupID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, resp.Infos[0].AppID, groupUser.AppID)
		assert.Equal(t, resp.Infos[0].UserID, groupUser.UserID)
		assert.Equal(t, resp.Infos[0].GroupID, groupUser.GroupID)
		groupUser.ID = resp.Infos[0].ID
	}

	resp1, err := Get(context.Background(), &npool.GetGroupUsersRequest{
		GroupID: groupUser.GroupID,
		AppID:   groupUser.AppID,
	})
	assert.Nil(t, err)
	fmt.Printf("get all group users resp1 is: %v", resp1)

	_, err = Delete(context.Background(), &npool.RemoveGroupUsersRequest{
		GroupID: groupUser.GroupID,
		AppID:   groupUser.AppID,
		UserIDs: []string{groupUser.UserID},
	})
	assert.Nil(t, err)
}
