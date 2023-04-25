package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2/persist"
)

// 可选项
type options struct {
	// 管理员角色，可选
	adminRole string
	// 路由过滤，可选
	pathFilter func(objects []string, path string) bool
	// 监视器，可选，分布式同步时需要
	watcher persist.Watcher
	// 域名,默认为空
	domain string
	// 角色前缀
	rolePrefix string
}

type Option interface {
	apply(*options)
}

type adminRoleOption string

func (o adminRoleOption) apply(opts *options) {
	opts.adminRole = string(o)
}

// 设置超管角色选项
func WithAdminRole(a string) Option {
	return adminRoleOption(a)
}

type rolePrefixOption string

func (o rolePrefixOption) apply(opts *options) {
	opts.rolePrefix = string(o)
}

func roleWithPrefix(prefix, role string) string {
	return fmt.Sprintf("%s_%s", prefix, role)
}

// 设置角色前缀选项
func WithRolePrefix(a string) Option {
	return rolePrefixOption(a)
}

type pathFilterOption func(objects []string, path string) bool

func (o pathFilterOption) apply(opts *options) {
	opts.pathFilter = o
}

// 设置路由过滤函数
func WithPathFilter(f func(objects []string, path string) bool) Option {
	return pathFilterOption(f)
}

type watcherOption struct {
	persist.Watcher
}

func (o watcherOption) apply(opts *options) {
	opts.watcher = o
}

// 设置分布式watcher
func WithWatcher(w persist.Watcher) Option {
	return watcherOption{w}
}

type domainOption string

func (o domainOption) apply(opts *options) {
	opts.domain = string(o)
}

// 设置域名
func WithDomain(d string) Option {
	return domainOption(d)
}
