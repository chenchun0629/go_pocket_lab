// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "admin_son", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admins_admins_son",
				Columns:    []*schema.Column{AdminsColumns[3]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BeesColumns holds the columns for the "bees" table.
	BeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// BeesTable holds the schema information for the "bees" table.
	BeesTable = &schema.Table{
		Name:       "bees",
		Columns:    BeesColumns,
		PrimaryKey: []*schema.Column{BeesColumns[0]},
	}
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "model", Type: field.TypeString},
		{Name: "registered_at", Type: field.TypeTime},
		{Name: "user_cars", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_users_cars",
				Columns:    []*schema.Column{CarsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CatsColumns holds the columns for the "cats" table.
	CatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "admin_cats", Type: field.TypeInt, Nullable: true},
	}
	// CatsTable holds the schema information for the "cats" table.
	CatsTable = &schema.Table{
		Name:       "cats",
		Columns:    CatsColumns,
		PrimaryKey: []*schema.Column{CatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cats_admins_cats",
				Columns:    []*schema.Column{CatsColumns[3]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FieldTestsColumns holds the columns for the "field_tests" table.
	FieldTestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_d", Type: field.TypeTime},
		{Name: "date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "title", Type: field.TypeString, Size: 64},
		{Name: "j_f", Type: field.TypeJSON},
		{Name: "j_s_f", Type: field.TypeJSON},
		{Name: "app_id", Type: field.TypeString},
	}
	// FieldTestsTable holds the schema information for the "field_tests" table.
	FieldTestsTable = &schema.Table{
		Name:       "field_tests",
		Columns:    FieldTestsColumns,
		PrimaryKey: []*schema.Column{FieldTestsColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// AdminBeesColumns holds the columns for the "admin_bees" table.
	AdminBeesColumns = []*schema.Column{
		{Name: "admin_id", Type: field.TypeInt},
		{Name: "bee_id", Type: field.TypeInt},
	}
	// AdminBeesTable holds the schema information for the "admin_bees" table.
	AdminBeesTable = &schema.Table{
		Name:       "admin_bees",
		Columns:    AdminBeesColumns,
		PrimaryKey: []*schema.Column{AdminBeesColumns[0], AdminBeesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admin_bees_admin_id",
				Columns:    []*schema.Column{AdminBeesColumns[0]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "admin_bees_bee_id",
				Columns:    []*schema.Column{AdminBeesColumns[1]},
				RefColumns: []*schema.Column{BeesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CatCbColumns holds the columns for the "cat_cb" table.
	CatCbColumns = []*schema.Column{
		{Name: "cat_id", Type: field.TypeInt},
		{Name: "bee_id", Type: field.TypeInt},
	}
	// CatCbTable holds the schema information for the "cat_cb" table.
	CatCbTable = &schema.Table{
		Name:       "cat_cb",
		Columns:    CatCbColumns,
		PrimaryKey: []*schema.Column{CatCbColumns[0], CatCbColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cat_cb_cat_id",
				Columns:    []*schema.Column{CatCbColumns[0]},
				RefColumns: []*schema.Column{CatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cat_cb_bee_id",
				Columns:    []*schema.Column{CatCbColumns[1]},
				RefColumns: []*schema.Column{BeesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		BeesTable,
		CarsTable,
		CatsTable,
		FieldTestsTable,
		GroupsTable,
		UsersTable,
		AdminBeesTable,
		CatCbTable,
		GroupUsersTable,
	}
)

func init() {
	AdminsTable.ForeignKeys[0].RefTable = AdminsTable
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	CatsTable.ForeignKeys[0].RefTable = AdminsTable
	FieldTestsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	AdminBeesTable.ForeignKeys[0].RefTable = AdminsTable
	AdminBeesTable.ForeignKeys[1].RefTable = BeesTable
	CatCbTable.ForeignKeys[0].RefTable = CatsTable
	CatCbTable.ForeignKeys[1].RefTable = BeesTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}
