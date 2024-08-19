package constant

// Known roleIds.
const (
	RoleIdSuperadmin uint64 = iota + 1
	RoleIdAdmin
	RoleIdUser
)

const DefaultRoleId = RoleIdUser

// Role permission.
var (
	// Registration.
	AllowedRolesRegister    = []uint64{RoleIdSuperadmin, RoleIdAdmin}
	AllowedRolesRegisterMap = map[uint64][]uint64{
		RoleIdSuperadmin: {RoleIdSuperadmin, RoleIdAdmin, RoleIdUser},
		RoleIdAdmin:      {RoleIdAdmin, RoleIdUser},
	}
)
