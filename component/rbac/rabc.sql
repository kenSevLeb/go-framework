
drop table if exists rbac_permission;
create table rbac_permission(
  id int unsigned not null primary key auto_increment comment '主键ID',
  p_name varchar(100) not null default '' comment '名称',
  p_type varchar(100) not null default '' comment '类型(module,page,btn)',
  p_path varchar(100) not null default '' comment '前端路由',
  unique_key varchar(100) not null default '' comment '唯一key(映射前端的路由)',
  is_del tinyint unsigned not null default 0 comment '是否删除(1:是,0:否)',
  created_at int not null default 0 comment '创建时间',
  updated_at int not null default 0 comment '更新时间',
  key `idx_unique_key`(`unique_key`),
  key `idx_is_del`(`is_del`),
  key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='权限信息';

drop table if exists rbac_router;
create table rbac_router(
  id int unsigned not null primary key auto_increment comment '主键ID',
  router varchar(100) not null default '' comment '名称',
  domain varchar(100) not null default '' comment '域名',
  unique_key varchar(100) not null default '' comment '唯一key',
  is_del tinyint unsigned not null default 0 comment '是否删除(1:是,0:否)',
  created_at int not null default 0 comment '创建时间',
  updated_at int not null default 0 comment '更新时间',
  key `idx_unique_key`(`unique_key`),
  key `idx_is_del`(`is_del`),
  key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='权限分组信息';

drop table if exists rbac_role;
create table rbac_role(
  id int unsigned not null primary key auto_increment comment '主键ID',
  role_name varchar(255) not null default '' comment '角色名称',
  is_del tinyint unsigned not null default 0 comment '是否删除(1:是,0:否)',
  created_at int not null default 0 comment '创建时间',
  updated_at int not null default 0 comment '更新时间',
  key `idx_is_del`(`is_del`),
  key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='角色信息表';

drop table if exists rbac_role_permission;
create table rbac_role_permission(
  id int unsigned not null primary key auto_increment comment '主键ID',
  role_id int not null default 0 comment '角色ID',
  permission_id int not null default 0 comment '权限ID',
  is_del tinyint unsigned not null default 0 comment '是否删除(1:是,0:否)',
  created_at int not null default 0 comment '创建时间',
  updated_at int not null default 0 comment '更新时间',
  key `idx_is_del`(`is_del`),
  key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='角色权限信息';

drop table if exists rbac_user_role;
create table rbac_user_role(
   id int unsigned not null primary key auto_increment comment '主键ID',
   user_id int not null default 0 comment '用户ID',
   role_id int not null default 0 comment '角色ID',
   is_del tinyint unsigned not null default 0 comment '是否删除(1:是,0:否)',
   created_at int not null default 0 comment '创建时间',
   updated_at int not null default 0 comment '更新时间',
   key `idx_is_del`(`is_del`),
   key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='用户角色信息';

/* 万物始于权限同步:1.同步权限树,2:授权角色 */
insert into rbac_permission(p_name, p_type, unique_key) values( '权限同步', 'page',  'role_tree');
insert into rbac_router(router, domain, unique_key) values('/protected/permission/sync', 'app', 'permission.sync');

insert into rbac_role(role_name) values('admin');
insert into rbac_role_permission(role_id, permission_id) values(1, 1);
insert into rbac_user_role(user_id, role_id) values(1,1);