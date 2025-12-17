package sys

import (
	"time"
)

// SysRoleView represents the sys_role_view table
type SysRoleView struct {
	Id          string     `json:"id" db:"pk;size:36;default:uuid()"`
	RoleId      string     `json:"roleId" db:"size:36;uk:role_view"`
	ViewId      string     `json:"viewId" db:"size:36;uk:role_view"`
	CreatedOn   time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy   string     `json:"createdBy" db:"size:36"`
	ModifiedOn  *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy  *string    `json:"modifiedBy" db:"size:36"`
	Description *string    `json:"description" db:"size:500"`
}

func (SysRoleView) TableName() string {
	return "sys_role_view"
}
