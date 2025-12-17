package sys

import (
	"time"
)

// SysRole represents the sys_roles table
type SysRole struct {
	Id          string     `json:"id" db:"pk;size:36;default:uuid()"`
	Name        string     `json:"name" db:"size:100;uk"`
	Code        string     `json:"code" db:"size:50;uk"` // use uk insted of unique
	CreatedOn   time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy   string     `json:"createdBy" db:"size:36"`
	ModifiedOn  *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy  *string    `json:"modifiedBy" db:"size:36"`
	Description *string    `json:"description" db:"size:500"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
