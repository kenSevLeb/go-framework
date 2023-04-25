package rbac

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rolePermissions = []RolePermission{
		{"r11", "/user/list", ""},
		{"r12", "/user/delete", ""},
		{"r13", "/user/info/:id", ""},
	}
	userRoles = []UserRole{
		{"1", "r11"},
		{"2", "r12"},
		{"3", "r13"},
		{"11", "r21"},
	}
)

func TestPermission_MustValidate(t *testing.T) {
	p := New()
	p.Load(func() []RolePermission {
		return rolePermissions
	}, func() []UserRole {
		return userRoles
	})
	assert.True(t, p.MustValidate("1", "/user/list"))
	assert.False(t, p.MustValidate("2", "/user/list"))
	assert.True(t, p.MustValidate("3", "/user/info/:1"))
	assert.False(t, p.MustValidate("11", "/user/list"))

}
