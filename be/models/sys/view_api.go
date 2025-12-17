package sys

import (
	"time"
)

// SysViewApi represents the sys_view_api table
type SysViewApi struct {
	Id          string     `json:"id" db:"pk;size:36;default:uuid()"`
	ViewId      string     `json:"viewId" db:"size:36;uk:view_api"`
	ApiId       string     `json:"apiId" db:"size:36;uk:view_api"`
	CreatedOn   time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy   string     `json:"createdBy" db:"size:36"`
	ModifiedOn  *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy  *string    `json:"modifiedBy" db:"size:36"`
	Description *string    `json:"description" db:"size:500"`
}

func (SysViewApi) TableName() string {
	return "sys_view_api"
}
