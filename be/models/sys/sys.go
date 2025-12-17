package sys

import "github.com/vn-go/dx"

func init() {
	// User -> Role
	// Note: create foreign key use field of struct not column name in db
	// RoleId in struct is role_id in db
	dx.AddForeignKey[SysUser]("RoleId", &SysRole{}, "Id", &dx.FkOpt{
		OnDelete: true, //"cascade",
		OnUpdate: true, //"cascade",
	})

	// RoleView -> Role
	dx.AddForeignKey[SysRoleView]("RoleId", &SysRole{}, "Id", &dx.FkOpt{
		OnDelete: true,
		OnUpdate: true,
	})
	// RoleView -> View
	dx.AddForeignKey[SysRoleView]("ViewId", &SysView{}, "Id", &dx.FkOpt{
		OnDelete: true,
		OnUpdate: true,
	})

	// View -> Parent View
	dx.AddForeignKey[SysView]("ParentId", &SysView{}, "Id", &dx.FkOpt{
		OnDelete: true,
		OnUpdate: true,
	})

	// ViewApi -> View
	dx.AddForeignKey[SysViewApi]("ViewId", &SysView{}, "Id", &dx.FkOpt{
		OnDelete: true,
		OnUpdate: true,
	})
	// ViewApi -> Api
	dx.AddForeignKey[SysViewApi]("ApiId", &SysApi{}, "Id", &dx.FkOpt{
		OnDelete: true,
		OnUpdate: true,
	})
}
