// Code generated by entc, DO NOT EDIT.

package ent

import (
	"pocket_lab/ent/ent/admin"
	"pocket_lab/ent/ent/bee"
	"pocket_lab/ent/ent/cat"
	"pocket_lab/ent/ent/group"
	"pocket_lab/ent/ent/schema"
	"pocket_lab/ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescAge is the schema descriptor for age field.
	adminDescAge := adminFields[0].Descriptor()
	// admin.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	admin.AgeValidator = adminDescAge.Validators[0].(func(int) error)
	// adminDescName is the schema descriptor for name field.
	adminDescName := adminFields[1].Descriptor()
	// admin.DefaultName holds the default value on creation for the name field.
	admin.DefaultName = adminDescName.Default.(string)
	beeFields := schema.Bee{}.Fields()
	_ = beeFields
	// beeDescAge is the schema descriptor for age field.
	beeDescAge := beeFields[0].Descriptor()
	// bee.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	bee.AgeValidator = beeDescAge.Validators[0].(func(int) error)
	// beeDescName is the schema descriptor for name field.
	beeDescName := beeFields[1].Descriptor()
	// bee.DefaultName holds the default value on creation for the name field.
	bee.DefaultName = beeDescName.Default.(string)
	catFields := schema.Cat{}.Fields()
	_ = catFields
	// catDescAge is the schema descriptor for age field.
	catDescAge := catFields[0].Descriptor()
	// cat.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	cat.AgeValidator = catDescAge.Validators[0].(func(int) error)
	// catDescName is the schema descriptor for name field.
	catDescName := catFields[1].Descriptor()
	// cat.DefaultName holds the default value on creation for the name field.
	cat.DefaultName = catDescName.Default.(string)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
