// Code generated by entc, DO NOT EDIT.

package bee

const (
	// Label holds the string label denoting the bee type in the database.
	Label = "bee"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeBc holds the string denoting the bc edge name in mutations.
	EdgeBc = "bc"
	// Table holds the table name of the bee in the database.
	Table = "bees"
	// OwnerTable is the table that holds the owner relation/edge. The primary key declared below.
	OwnerTable = "admin_bees"
	// OwnerInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	OwnerInverseTable = "admins"
	// BcTable is the table that holds the bc relation/edge. The primary key declared below.
	BcTable = "cat_cb"
	// BcInverseTable is the table name for the Cat entity.
	// It exists in this package in order to avoid circular dependency with the "cat" package.
	BcInverseTable = "cats"
)

// Columns holds all SQL columns for bee fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// OwnerPrimaryKey and OwnerColumn2 are the table columns denoting the
	// primary key for the owner relation (M2M).
	OwnerPrimaryKey = []string{"admin_id", "bee_id"}
	// BcPrimaryKey and BcColumn2 are the table columns denoting the
	// primary key for the bc relation (M2M).
	BcPrimaryKey = []string{"cat_id", "bee_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
