package exist

import (
	"context"

	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/application"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroup"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func Application(ctx context.Context, appID string) (bool, error) {
	info, err := db.Client().
		Application.
		Query().
		Where(
			application.And(
				application.DeleteAt(0),
				application.ID(appID),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get application: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func ApplicationRole(ctx context.Context, roleID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.ID(roleID),
				applicationrole.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to query role: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func RoleName(ctx context.Context, roleName, appID string) (int, error) {
	info, err := db.Client().
		ApplicationRole.
		Query().
		Where(
			applicationrole.And(
				applicationrole.DeleteAt(0),
				applicationrole.RoleName(roleName),
				applicationrole.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return -1, xerrors.Errorf("fail to query role: %v", err)
	}

	if len(info) == 0 {
		return 0, nil
	} else if len(info) == 1 {
		return 1, nil
	}

	return -1, nil
}

func UserRole(ctx context.Context, userID, roleID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationRoleUser.
		Query().
		Where(
			applicationroleuser.And(
				applicationroleuser.UserID(userID),
				applicationroleuser.RoleID(roleID),
				applicationroleuser.AppID(appID),
				applicationroleuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get user role: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func ApplicationUser(ctx context.Context, appID string, userID uuid.UUID) (bool, error) {
	info, err := db.Client().
		ApplicationUser.
		Query().
		Where(
			applicationuser.And(
				applicationuser.UserID(userID),
				applicationuser.AppID(appID),
				applicationuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get application user: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func ApplicationGroup(ctx context.Context, groupID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.ID(groupID),
				applicationgroup.AppID(appID),
				applicationgroup.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get application group: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func GroupName(ctx context.Context, groupName, appID string) (int, error) {
	info, err := db.Client().
		ApplicationGroup.
		Query().
		Where(
			applicationgroup.And(
				applicationgroup.DeleteAt(0),
				applicationgroup.GroupName(groupName),
				applicationgroup.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return -1, xerrors.Errorf("fail to query group by name: %v", err)
	}

	if len(info) == 0 {
		return 0, nil
	} else if len(info) == 1 {
		return 1, nil
	}

	return -1, nil
}

func GroupUser(ctx context.Context, userID, groupID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationGroupUser.
		Query().
		Where(
			applicationgroupuser.And(
				applicationgroupuser.GroupID(groupID),
				applicationgroupuser.AppID(appID),
				applicationgroupuser.UserID(userID),
				applicationgroupuser.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get application group user: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func ApplicationResource(ctx context.Context, resourceID uuid.UUID, appID string) (bool, error) {
	info, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.ID(resourceID),
				applicationresource.AppID(appID),
				applicationresource.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return true, xerrors.Errorf("fail to get resource: %v", err)
	}

	if len(info) == 0 {
		return false, nil
	}

	return true, nil
}

func ResourceName(ctx context.Context, resourceName, appID string) (int, error) {
	info, err := db.Client().
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.DeleteAt(0),
				applicationresource.ResourceName(resourceName),
				applicationresource.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return -1, xerrors.Errorf("fail to query resource by name: %v", err)
	}

	if len(info) == 0 {
		return 0, nil
	} else if len(info) == 1 {
		return 1, nil
	}

	return -1, nil
}
