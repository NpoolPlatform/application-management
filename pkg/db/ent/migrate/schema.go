// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApplicationsColumns holds the columns for the "applications" table.
	ApplicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "application_name", Type: field.TypeString},
		{Name: "application_owner", Type: field.TypeUUID},
		{Name: "homepage_url", Type: field.TypeString, Nullable: true},
		{Name: "redirect_url", Type: field.TypeString, Nullable: true},
		{Name: "client_secret", Type: field.TypeString, Unique: true},
		{Name: "application_logo", Type: field.TypeString, Nullable: true},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationsTable holds the schema information for the "applications" table.
	ApplicationsTable = &schema.Table{
		Name:       "applications",
		Columns:    ApplicationsColumns,
		PrimaryKey: []*schema.Column{ApplicationsColumns[0]},
	}
	// ApplicationGroupsColumns holds the columns for the "application_groups" table.
	ApplicationGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "group_name", Type: field.TypeString},
		{Name: "group_logo", Type: field.TypeString, Nullable: true},
		{Name: "group_owner", Type: field.TypeUUID},
		{Name: "annotation", Type: field.TypeString, Nullable: true},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationGroupsTable holds the schema information for the "application_groups" table.
	ApplicationGroupsTable = &schema.Table{
		Name:       "application_groups",
		Columns:    ApplicationGroupsColumns,
		PrimaryKey: []*schema.Column{ApplicationGroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationgroup_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupsColumns[1]},
			},
			{
				Name:    "applicationgroup_group_owner",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupsColumns[4]},
			},
		},
	}
	// ApplicationGroupUsersColumns holds the columns for the "application_group_users" table.
	ApplicationGroupUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "group_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "annotation", Type: field.TypeString, Nullable: true},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationGroupUsersTable holds the schema information for the "application_group_users" table.
	ApplicationGroupUsersTable = &schema.Table{
		Name:       "application_group_users",
		Columns:    ApplicationGroupUsersColumns,
		PrimaryKey: []*schema.Column{ApplicationGroupUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationgroupuser_group_id_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupUsersColumns[1], ApplicationGroupUsersColumns[2]},
			},
			{
				Name:    "applicationgroupuser_user_id_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupUsersColumns[3], ApplicationGroupUsersColumns[2]},
			},
			{
				Name:    "applicationgroupuser_user_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupUsersColumns[3]},
			},
			{
				Name:    "applicationgroupuser_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationGroupUsersColumns[2]},
			},
		},
	}
	// ApplicationResourcesColumns holds the columns for the "application_resources" table.
	ApplicationResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "resource_name", Type: field.TypeString},
		{Name: "resource_description", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString, Default: "API"},
		{Name: "creator", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationResourcesTable holds the schema information for the "application_resources" table.
	ApplicationResourcesTable = &schema.Table{
		Name:       "application_resources",
		Columns:    ApplicationResourcesColumns,
		PrimaryKey: []*schema.Column{ApplicationResourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationresource_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationResourcesColumns[1]},
			},
		},
	}
	// ApplicationRolesColumns holds the columns for the "application_roles" table.
	ApplicationRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "role_name", Type: field.TypeString},
		{Name: "creator", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
		{Name: "annotation", Type: field.TypeString, Nullable: true},
	}
	// ApplicationRolesTable holds the schema information for the "application_roles" table.
	ApplicationRolesTable = &schema.Table{
		Name:       "application_roles",
		Columns:    ApplicationRolesColumns,
		PrimaryKey: []*schema.Column{ApplicationRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationrole_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationRolesColumns[1]},
			},
			{
				Name:    "applicationrole_creator",
				Unique:  false,
				Columns: []*schema.Column{ApplicationRolesColumns[3]},
			},
		},
	}
	// ApplicationRoleUsersColumns holds the columns for the "application_role_users" table.
	ApplicationRoleUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationRoleUsersTable holds the schema information for the "application_role_users" table.
	ApplicationRoleUsersTable = &schema.Table{
		Name:       "application_role_users",
		Columns:    ApplicationRoleUsersColumns,
		PrimaryKey: []*schema.Column{ApplicationRoleUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationroleuser_app_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationRoleUsersColumns[1], ApplicationRoleUsersColumns[3]},
			},
			{
				Name:    "applicationroleuser_app_id_role_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationRoleUsersColumns[1], ApplicationRoleUsersColumns[2]},
			},
		},
	}
	// ApplicationUsersColumns holds the columns for the "application_users" table.
	ApplicationUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "original", Type: field.TypeBool, Default: true},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// ApplicationUsersTable holds the schema information for the "application_users" table.
	ApplicationUsersTable = &schema.Table{
		Name:       "application_users",
		Columns:    ApplicationUsersColumns,
		PrimaryKey: []*schema.Column{ApplicationUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationuser_app_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationUsersColumns[1], ApplicationUsersColumns[2]},
			},
			{
				Name:    "applicationuser_app_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationUsersColumns[1]},
			},
			{
				Name:    "applicationuser_user_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationUsersColumns[2]},
			},
			{
				Name:    "applicationuser_original",
				Unique:  false,
				Columns: []*schema.Column{ApplicationUsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicationsTable,
		ApplicationGroupsTable,
		ApplicationGroupUsersTable,
		ApplicationResourcesTable,
		ApplicationRolesTable,
		ApplicationRoleUsersTable,
		ApplicationUsersTable,
	}
)

func init() {
}
