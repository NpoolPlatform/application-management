// Code generated by entc, DO NOT EDIT.

package application

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ApplicationName applies equality check predicate on the "application_name" field. It's identical to ApplicationNameEQ.
func ApplicationName(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationName), v))
	})
}

// ApplicationOwner applies equality check predicate on the "application_owner" field. It's identical to ApplicationOwnerEQ.
func ApplicationOwner(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationOwner), v))
	})
}

// HomepageURL applies equality check predicate on the "homepage_url" field. It's identical to HomepageURLEQ.
func HomepageURL(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHomepageURL), v))
	})
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// ApplicationLogo applies equality check predicate on the "application_logo" field. It's identical to ApplicationLogoEQ.
func ApplicationLogo(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationLogo), v))
	})
}

// SmsLogin applies equality check predicate on the "sms_login" field. It's identical to SmsLoginEQ.
func SmsLogin(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmsLogin), v))
	})
}

// GoogleRecaptcha applies equality check predicate on the "google_recaptcha" field. It's identical to GoogleRecaptchaEQ.
func GoogleRecaptcha(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoogleRecaptcha), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// ApplicationNameEQ applies the EQ predicate on the "application_name" field.
func ApplicationNameEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameNEQ applies the NEQ predicate on the "application_name" field.
func ApplicationNameNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameIn applies the In predicate on the "application_name" field.
func ApplicationNameIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApplicationName), v...))
	})
}

// ApplicationNameNotIn applies the NotIn predicate on the "application_name" field.
func ApplicationNameNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApplicationName), v...))
	})
}

// ApplicationNameGT applies the GT predicate on the "application_name" field.
func ApplicationNameGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameGTE applies the GTE predicate on the "application_name" field.
func ApplicationNameGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameLT applies the LT predicate on the "application_name" field.
func ApplicationNameLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameLTE applies the LTE predicate on the "application_name" field.
func ApplicationNameLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameContains applies the Contains predicate on the "application_name" field.
func ApplicationNameContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameHasPrefix applies the HasPrefix predicate on the "application_name" field.
func ApplicationNameHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameHasSuffix applies the HasSuffix predicate on the "application_name" field.
func ApplicationNameHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameEqualFold applies the EqualFold predicate on the "application_name" field.
func ApplicationNameEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldApplicationName), v))
	})
}

// ApplicationNameContainsFold applies the ContainsFold predicate on the "application_name" field.
func ApplicationNameContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldApplicationName), v))
	})
}

// ApplicationOwnerEQ applies the EQ predicate on the "application_owner" field.
func ApplicationOwnerEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationOwner), v))
	})
}

// ApplicationOwnerNEQ applies the NEQ predicate on the "application_owner" field.
func ApplicationOwnerNEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplicationOwner), v))
	})
}

// ApplicationOwnerIn applies the In predicate on the "application_owner" field.
func ApplicationOwnerIn(vs ...uuid.UUID) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApplicationOwner), v...))
	})
}

// ApplicationOwnerNotIn applies the NotIn predicate on the "application_owner" field.
func ApplicationOwnerNotIn(vs ...uuid.UUID) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApplicationOwner), v...))
	})
}

// ApplicationOwnerGT applies the GT predicate on the "application_owner" field.
func ApplicationOwnerGT(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplicationOwner), v))
	})
}

// ApplicationOwnerGTE applies the GTE predicate on the "application_owner" field.
func ApplicationOwnerGTE(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplicationOwner), v))
	})
}

// ApplicationOwnerLT applies the LT predicate on the "application_owner" field.
func ApplicationOwnerLT(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplicationOwner), v))
	})
}

// ApplicationOwnerLTE applies the LTE predicate on the "application_owner" field.
func ApplicationOwnerLTE(v uuid.UUID) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplicationOwner), v))
	})
}

// HomepageURLEQ applies the EQ predicate on the "homepage_url" field.
func HomepageURLEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLNEQ applies the NEQ predicate on the "homepage_url" field.
func HomepageURLNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLIn applies the In predicate on the "homepage_url" field.
func HomepageURLIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHomepageURL), v...))
	})
}

// HomepageURLNotIn applies the NotIn predicate on the "homepage_url" field.
func HomepageURLNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHomepageURL), v...))
	})
}

// HomepageURLGT applies the GT predicate on the "homepage_url" field.
func HomepageURLGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLGTE applies the GTE predicate on the "homepage_url" field.
func HomepageURLGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLLT applies the LT predicate on the "homepage_url" field.
func HomepageURLLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLLTE applies the LTE predicate on the "homepage_url" field.
func HomepageURLLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLContains applies the Contains predicate on the "homepage_url" field.
func HomepageURLContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLHasPrefix applies the HasPrefix predicate on the "homepage_url" field.
func HomepageURLHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLHasSuffix applies the HasSuffix predicate on the "homepage_url" field.
func HomepageURLHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLIsNil applies the IsNil predicate on the "homepage_url" field.
func HomepageURLIsNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHomepageURL)))
	})
}

// HomepageURLNotNil applies the NotNil predicate on the "homepage_url" field.
func HomepageURLNotNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHomepageURL)))
	})
}

// HomepageURLEqualFold applies the EqualFold predicate on the "homepage_url" field.
func HomepageURLEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHomepageURL), v))
	})
}

// HomepageURLContainsFold applies the ContainsFold predicate on the "homepage_url" field.
func HomepageURLContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHomepageURL), v))
	})
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIsNil applies the IsNil predicate on the "redirect_url" field.
func RedirectURLIsNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRedirectURL)))
	})
}

// RedirectURLNotNil applies the NotNil predicate on the "redirect_url" field.
func RedirectURLNotNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRedirectURL)))
	})
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRedirectURL), v))
	})
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientSecret), v))
	})
}

// ApplicationLogoEQ applies the EQ predicate on the "application_logo" field.
func ApplicationLogoEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoNEQ applies the NEQ predicate on the "application_logo" field.
func ApplicationLogoNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoIn applies the In predicate on the "application_logo" field.
func ApplicationLogoIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApplicationLogo), v...))
	})
}

// ApplicationLogoNotIn applies the NotIn predicate on the "application_logo" field.
func ApplicationLogoNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApplicationLogo), v...))
	})
}

// ApplicationLogoGT applies the GT predicate on the "application_logo" field.
func ApplicationLogoGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoGTE applies the GTE predicate on the "application_logo" field.
func ApplicationLogoGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoLT applies the LT predicate on the "application_logo" field.
func ApplicationLogoLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoLTE applies the LTE predicate on the "application_logo" field.
func ApplicationLogoLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoContains applies the Contains predicate on the "application_logo" field.
func ApplicationLogoContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoHasPrefix applies the HasPrefix predicate on the "application_logo" field.
func ApplicationLogoHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoHasSuffix applies the HasSuffix predicate on the "application_logo" field.
func ApplicationLogoHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoIsNil applies the IsNil predicate on the "application_logo" field.
func ApplicationLogoIsNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldApplicationLogo)))
	})
}

// ApplicationLogoNotNil applies the NotNil predicate on the "application_logo" field.
func ApplicationLogoNotNil() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldApplicationLogo)))
	})
}

// ApplicationLogoEqualFold applies the EqualFold predicate on the "application_logo" field.
func ApplicationLogoEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldApplicationLogo), v))
	})
}

// ApplicationLogoContainsFold applies the ContainsFold predicate on the "application_logo" field.
func ApplicationLogoContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldApplicationLogo), v))
	})
}

// SmsLoginEQ applies the EQ predicate on the "sms_login" field.
func SmsLoginEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmsLogin), v))
	})
}

// SmsLoginNEQ applies the NEQ predicate on the "sms_login" field.
func SmsLoginNEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSmsLogin), v))
	})
}

// GoogleRecaptchaEQ applies the EQ predicate on the "google_recaptcha" field.
func GoogleRecaptchaEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoogleRecaptcha), v))
	})
}

// GoogleRecaptchaNEQ applies the NEQ predicate on the "google_recaptcha" field.
func GoogleRecaptchaNEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoogleRecaptcha), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		p(s.Not())
	})
}
