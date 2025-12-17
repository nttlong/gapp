package sys

import (
	"time"
)

// SysApi represents the sys_api table
type SysApi struct {
	Id          string     `json:"id" db:"pk;size:36;default:uuid()"`
	Method      string     `json:"method" db:"size:10"`
	Path        string     `json:"path" db:"size:255"`
	Code        string     `json:"code" db:"size:50;uk"`
	CreatedOn   time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy   string     `json:"createdBy" db:"size:36"`
	ModifiedOn  *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy  *string    `json:"modifiedBy" db:"size:36"`
	Description *string    `json:"description" db:"size:500"`
}

func (SysApi) TableName() string {
	return "sys_api"
}
