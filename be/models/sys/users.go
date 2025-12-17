package sys

import (
	"time"
)

// SysUser represents the sys_users table
type SysUser struct {
	Id           string     `json:"id" db:"pk;size:36;default:uuid()"`
	Username     string     `json:"username" db:"size:50;uk"`
	HashPassword string     `json:"-" db:"size:255"` // hash password
	Email        string     `json:"email" db:"size:100;uk"`
	Phone        string     `json:"phone" db:"size:20"`
	Avatar       string     `json:"avatar" db:"size:255"`
	Status       int        `json:"status" db:"default:1"`
	CreatedOn    time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy    string     `json:"createdBy" db:"size:36"`
	ModifiedOn   *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy   *string    `json:"modifiedBy" db:"size:36"`
	Description  *string    `json:"description" db:"size:500"`
	RoleId       *string    `json:"roleId" db:"size:36"` //<-- columname in db automatically convert to role_id
}

func (SysUser) TableName() string {
	return "sys_users"
}
