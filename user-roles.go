package authority

import "github.com/google/uuid"

// UserRoleInt represents the relationship between users and roles
type UserRoleInt struct {
	ID     uint
	UserID uint
	RoleID uint
}

type UserRoleUUID struct {
	ID     uint
	UserID uuid.UUID
	RoleID uint
}

// TableName sets the table name
func (u UserRoleInt) TableName() string {
	return tablePrefix + "user_roles_int"
}

// TableName sets the table name
func (u UserRoleUUID) TableName() string {
	return tablePrefix + "user_roles_uuid"
}
