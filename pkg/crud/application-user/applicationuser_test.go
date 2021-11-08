package applicationuser

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

func TestApplicationUserCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	applicationUser := &npool.ApplicationUserInfo{
		AppID:    randstr.Hex(10),
		UserID:   uuid.New().String(),
		Original: true,
	}

	resp, err := Create(context.Background(), &npool.AddUsersToApplicationRequest{
		AppID:    applicationUser.AppID,
		UserIDs:  []string{applicationUser.UserID},
		Original: applicationUser.Original,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Infos[0].ID, uuid.UUID{})
		assert.Equal(t, resp.Infos[0].AppID, applicationUser.AppID)
		assert.Equal(t, resp.Infos[0].UserID, applicationUser.UserID)
		assert.Equal(t, resp.Infos[0].Original, applicationUser.Original)
		applicationUser.ID = resp.Infos[0].ID
	}

	resp1, err := Get(context.Background(), &npool.GetUserFromApplicationRequest{
		AppID:  applicationUser.AppID,
		UserID: applicationUser.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, applicationUser.ID)
		assert.Equal(t, resp1.Info.UserID, applicationUser.UserID)
		assert.Equal(t, resp1.Info.AppID, applicationUser.AppID)
		assert.Equal(t, resp1.Info.Original, applicationUser.Original)
	}

	_, err = GetAll(context.Background(), &npool.GetUsersRequest{
		AppID: applicationUser.AppID,
	})
	assert.Nil(t, err)

	_, err = Delete(context.Background(), &npool.RemoveUsersFromApplicationRequest{
		UserIDs: []string{applicationUser.UserID},
		AppID:   applicationUser.AppID,
	})
	assert.Nil(t, err)
}
