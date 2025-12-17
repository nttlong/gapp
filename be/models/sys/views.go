package sys

import (
	"time"
)

// SysView represents the sys_views table
type SysView struct {
	Id          string     `json:"id" db:"pk;size:36;default:uuid()"`
	Name        string     `json:"name" db:"size:100;uk"`
	Path        string     `json:"path" db:"size:255"`
	Icon        string     `json:"icon" db:"size:100"`
	SortOrder   int        `json:"sortOrder" db:"default:0"`
	ParentId    string     `json:"parentId" db:"size:36"`
	CreatedOn   time.Time  `json:"createdOn" db:"auto_now_add"`
	CreatedBy   string     `json:"createdBy" db:"size:36"`
	ModifiedOn  *time.Time `json:"modifiedOn" db:"auto_now"`
	ModifiedBy  *string    `json:"modifiedBy" db:"size:36"`
	Description *string    `json:"description" db:"size:500"`
}

func (SysView) TableName() string {
	return "sys_views"
}
