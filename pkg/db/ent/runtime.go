// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/application-management/pkg/db/ent/application"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroup"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	applicationFields := schema.Application{}.Fields()
	_ = applicationFields
	// applicationDescClientSecret is the schema descriptor for client_secret field.
	applicationDescClientSecret := applicationFields[5].Descriptor()
	// application.DefaultClientSecret holds the default value on creation for the client_secret field.
	application.DefaultClientSecret = applicationDescClientSecret.Default.(func() string)
	// applicationDescSmsLogin is the schema descriptor for sms_login field.
	applicationDescSmsLogin := applicationFields[7].Descriptor()
	// application.DefaultSmsLogin holds the default value on creation for the sms_login field.
	application.DefaultSmsLogin = applicationDescSmsLogin.Default.(bool)
	// applicationDescGoogleRecaptcha is the schema descriptor for google_recaptcha field.
	applicationDescGoogleRecaptcha := applicationFields[8].Descriptor()
	// application.DefaultGoogleRecaptcha holds the default value on creation for the google_recaptcha field.
	application.DefaultGoogleRecaptcha = applicationDescGoogleRecaptcha.Default.(bool)
	// applicationDescCreateAt is the schema descriptor for create_at field.
	applicationDescCreateAt := applicationFields[9].Descriptor()
	// application.DefaultCreateAt holds the default value on creation for the create_at field.
	application.DefaultCreateAt = applicationDescCreateAt.Default.(func() uint32)
	// applicationDescUpdateAt is the schema descriptor for update_at field.
	applicationDescUpdateAt := applicationFields[10].Descriptor()
	// application.DefaultUpdateAt holds the default value on creation for the update_at field.
	application.DefaultUpdateAt = applicationDescUpdateAt.Default.(func() uint32)
	// application.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	application.UpdateDefaultUpdateAt = applicationDescUpdateAt.UpdateDefault.(func() uint32)
	// applicationDescDeleteAt is the schema descriptor for delete_at field.
	applicationDescDeleteAt := applicationFields[11].Descriptor()
	// application.DefaultDeleteAt holds the default value on creation for the delete_at field.
	application.DefaultDeleteAt = applicationDescDeleteAt.Default.(func() uint32)
	// applicationDescID is the schema descriptor for id field.
	applicationDescID := applicationFields[0].Descriptor()
	// application.DefaultID holds the default value on creation for the id field.
	application.DefaultID = applicationDescID.Default.(func() uuid.UUID)
	applicationgroupFields := schema.ApplicationGroup{}.Fields()
	_ = applicationgroupFields
	// applicationgroupDescCreateAt is the schema descriptor for create_at field.
	applicationgroupDescCreateAt := applicationgroupFields[6].Descriptor()
	// applicationgroup.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationgroup.DefaultCreateAt = applicationgroupDescCreateAt.Default.(func() uint32)
	// applicationgroupDescUpdateAt is the schema descriptor for update_at field.
	applicationgroupDescUpdateAt := applicationgroupFields[7].Descriptor()
	// applicationgroup.DefaultUpdateAt holds the default value on creation for the update_at field.
	applicationgroup.DefaultUpdateAt = applicationgroupDescUpdateAt.Default.(func() uint32)
	// applicationgroup.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	applicationgroup.UpdateDefaultUpdateAt = applicationgroupDescUpdateAt.UpdateDefault.(func() uint32)
	// applicationgroupDescDeleteAt is the schema descriptor for delete_at field.
	applicationgroupDescDeleteAt := applicationgroupFields[8].Descriptor()
	// applicationgroup.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationgroup.DefaultDeleteAt = applicationgroupDescDeleteAt.Default.(func() uint32)
	// applicationgroupDescID is the schema descriptor for id field.
	applicationgroupDescID := applicationgroupFields[0].Descriptor()
	// applicationgroup.DefaultID holds the default value on creation for the id field.
	applicationgroup.DefaultID = applicationgroupDescID.Default.(func() uuid.UUID)
	applicationgroupuserFields := schema.ApplicationGroupUser{}.Fields()
	_ = applicationgroupuserFields
	// applicationgroupuserDescCreateAt is the schema descriptor for create_at field.
	applicationgroupuserDescCreateAt := applicationgroupuserFields[5].Descriptor()
	// applicationgroupuser.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationgroupuser.DefaultCreateAt = applicationgroupuserDescCreateAt.Default.(func() uint32)
	// applicationgroupuserDescDeleteAt is the schema descriptor for delete_at field.
	applicationgroupuserDescDeleteAt := applicationgroupuserFields[6].Descriptor()
	// applicationgroupuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationgroupuser.DefaultDeleteAt = applicationgroupuserDescDeleteAt.Default.(func() uint32)
	// applicationgroupuserDescID is the schema descriptor for id field.
	applicationgroupuserDescID := applicationgroupuserFields[0].Descriptor()
	// applicationgroupuser.DefaultID holds the default value on creation for the id field.
	applicationgroupuser.DefaultID = applicationgroupuserDescID.Default.(func() uuid.UUID)
	applicationresourceFields := schema.ApplicationResource{}.Fields()
	_ = applicationresourceFields
	// applicationresourceDescType is the schema descriptor for type field.
	applicationresourceDescType := applicationresourceFields[4].Descriptor()
	// applicationresource.DefaultType holds the default value on creation for the type field.
	applicationresource.DefaultType = applicationresourceDescType.Default.(string)
	// applicationresourceDescCreateAt is the schema descriptor for create_at field.
	applicationresourceDescCreateAt := applicationresourceFields[6].Descriptor()
	// applicationresource.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationresource.DefaultCreateAt = applicationresourceDescCreateAt.Default.(func() uint32)
	// applicationresourceDescUpdateAt is the schema descriptor for update_at field.
	applicationresourceDescUpdateAt := applicationresourceFields[7].Descriptor()
	// applicationresource.DefaultUpdateAt holds the default value on creation for the update_at field.
	applicationresource.DefaultUpdateAt = applicationresourceDescUpdateAt.Default.(func() uint32)
	// applicationresource.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	applicationresource.UpdateDefaultUpdateAt = applicationresourceDescUpdateAt.UpdateDefault.(func() uint32)
	// applicationresourceDescDeleteAt is the schema descriptor for delete_at field.
	applicationresourceDescDeleteAt := applicationresourceFields[8].Descriptor()
	// applicationresource.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationresource.DefaultDeleteAt = applicationresourceDescDeleteAt.Default.(func() uint32)
	// applicationresourceDescID is the schema descriptor for id field.
	applicationresourceDescID := applicationresourceFields[0].Descriptor()
	// applicationresource.DefaultID holds the default value on creation for the id field.
	applicationresource.DefaultID = applicationresourceDescID.Default.(func() uuid.UUID)
	applicationroleFields := schema.ApplicationRole{}.Fields()
	_ = applicationroleFields
	// applicationroleDescCreateAt is the schema descriptor for create_at field.
	applicationroleDescCreateAt := applicationroleFields[4].Descriptor()
	// applicationrole.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationrole.DefaultCreateAt = applicationroleDescCreateAt.Default.(func() uint32)
	// applicationroleDescUpdateAt is the schema descriptor for update_at field.
	applicationroleDescUpdateAt := applicationroleFields[5].Descriptor()
	// applicationrole.DefaultUpdateAt holds the default value on creation for the update_at field.
	applicationrole.DefaultUpdateAt = applicationroleDescUpdateAt.Default.(func() uint32)
	// applicationrole.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	applicationrole.UpdateDefaultUpdateAt = applicationroleDescUpdateAt.UpdateDefault.(func() uint32)
	// applicationroleDescDeleteAt is the schema descriptor for delete_at field.
	applicationroleDescDeleteAt := applicationroleFields[6].Descriptor()
	// applicationrole.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationrole.DefaultDeleteAt = applicationroleDescDeleteAt.Default.(func() uint32)
	// applicationroleDescID is the schema descriptor for id field.
	applicationroleDescID := applicationroleFields[0].Descriptor()
	// applicationrole.DefaultID holds the default value on creation for the id field.
	applicationrole.DefaultID = applicationroleDescID.Default.(func() uuid.UUID)
	applicationroleuserFields := schema.ApplicationRoleUser{}.Fields()
	_ = applicationroleuserFields
	// applicationroleuserDescCreateAt is the schema descriptor for create_at field.
	applicationroleuserDescCreateAt := applicationroleuserFields[4].Descriptor()
	// applicationroleuser.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationroleuser.DefaultCreateAt = applicationroleuserDescCreateAt.Default.(func() uint32)
	// applicationroleuserDescDeleteAt is the schema descriptor for delete_at field.
	applicationroleuserDescDeleteAt := applicationroleuserFields[5].Descriptor()
	// applicationroleuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationroleuser.DefaultDeleteAt = applicationroleuserDescDeleteAt.Default.(func() uint32)
	// applicationroleuserDescID is the schema descriptor for id field.
	applicationroleuserDescID := applicationroleuserFields[0].Descriptor()
	// applicationroleuser.DefaultID holds the default value on creation for the id field.
	applicationroleuser.DefaultID = applicationroleuserDescID.Default.(func() uuid.UUID)
	applicationuserFields := schema.ApplicationUser{}.Fields()
	_ = applicationuserFields
	// applicationuserDescOriginal is the schema descriptor for original field.
	applicationuserDescOriginal := applicationuserFields[3].Descriptor()
	// applicationuser.DefaultOriginal holds the default value on creation for the original field.
	applicationuser.DefaultOriginal = applicationuserDescOriginal.Default.(bool)
	// applicationuserDescKycVerify is the schema descriptor for kyc_verify field.
	applicationuserDescKycVerify := applicationuserFields[4].Descriptor()
	// applicationuser.DefaultKycVerify holds the default value on creation for the kyc_verify field.
	applicationuser.DefaultKycVerify = applicationuserDescKycVerify.Default.(bool)
	// applicationuserDescGaVerify is the schema descriptor for ga_verify field.
	applicationuserDescGaVerify := applicationuserFields[5].Descriptor()
	// applicationuser.DefaultGaVerify holds the default value on creation for the ga_verify field.
	applicationuser.DefaultGaVerify = applicationuserDescGaVerify.Default.(bool)
	// applicationuserDescGaLogin is the schema descriptor for ga_login field.
	applicationuserDescGaLogin := applicationuserFields[6].Descriptor()
	// applicationuser.DefaultGaLogin holds the default value on creation for the ga_login field.
	applicationuser.DefaultGaLogin = applicationuserDescGaLogin.Default.(bool)
	// applicationuserDescLoginNumber is the schema descriptor for Login_number field.
	applicationuserDescLoginNumber := applicationuserFields[7].Descriptor()
	// applicationuser.DefaultLoginNumber holds the default value on creation for the Login_number field.
	applicationuser.DefaultLoginNumber = applicationuserDescLoginNumber.Default.(uint32)
	// applicationuserDescCreateAt is the schema descriptor for create_at field.
	applicationuserDescCreateAt := applicationuserFields[8].Descriptor()
	// applicationuser.DefaultCreateAt holds the default value on creation for the create_at field.
	applicationuser.DefaultCreateAt = applicationuserDescCreateAt.Default.(func() uint32)
	// applicationuserDescDeleteAt is the schema descriptor for delete_at field.
	applicationuserDescDeleteAt := applicationuserFields[9].Descriptor()
	// applicationuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applicationuser.DefaultDeleteAt = applicationuserDescDeleteAt.Default.(func() uint32)
	// applicationuserDescID is the schema descriptor for id field.
	applicationuserDescID := applicationuserFields[0].Descriptor()
	// applicationuser.DefaultID holds the default value on creation for the id field.
	applicationuser.DefaultID = applicationuserDescID.Default.(func() uuid.UUID)
}
