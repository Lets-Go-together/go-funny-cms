create table funy_cms.admins
(
    id int auto_increment comment '主键ID'
        primary key,
    account varchar(60) null comment '用户名称',
    password varchar(60) null comment '密码',
    description varchar(255) default '' null comment '用户描述',
    avatar varchar(255) null comment '头像',
    phone varchar(60) null comment '电话号码',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null comment '软删除时间戳',
    email varchar(60) null comment '邮箱账号',
    constraint admins_account_uindex
        unique (account)
)
    comment '管理员表';

create table funy_cms.casbin_rule
(
    id bigint unsigned auto_increment
        primary key,
    p_type varchar(40) null,
    v0 varchar(40) null,
    v1 varchar(40) null,
    v2 varchar(40) null,
    v3 varchar(40) null,
    v4 varchar(40) null,
    v5 varchar(40) null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    constraint unique_index
        unique (p_type, v0, v1, v2, v3, v4, v5)
);

create table funy_cms.email_record
(
    id bigint unsigned auto_increment
        primary key,
    email_id tinyint not null comment '关联的邮件ID',
    status int default 1 null comment '状态: 1: 启用; 2:禁用;',
    submitter_id tinyint null comment '关联的管理员ID',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null
)
    comment '邮件记录关联';

create table funy_cms.email_tasks
(
    id bigint unsigned auto_increment
        primary key,
    title varchar(60) default '' not null comment '邮件描述',
    mailer text not null comment '邮件配置',
    email varchar(255) not null comment '接受人',
    subject varchar(255) default '' not null comment '主题',
    content text not null comment '发送内容',
    attachments json not null comment '附件',
    status int default 1 null comment '状态: 1: 启用; 2:禁用;',
    submitter_id tinyint null comment '关联的管理员ID',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    send_at timestamp null comment '发送时间',
    remark varchar(255) null comment '备注'
)
    comment '邮件内容' engine=MyISAM collate=utf8mb4_unicode_ci;

create table funy_cms.menus
(
    id bigint unsigned auto_increment
        primary key,
    name varchar(60) default '' not null comment '角色名',
    status int default 1 null comment '状态: 1: 启用; 2:禁用;',
    description varchar(255) null comment '角色描述',
    p_id int default 1 null comment '父ID',
    weight int default 1 null comment '排序',
    url varchar(60) null,
    hidden int default 2 null comment '菜单隐藏',
    component varchar(60) null comment '组件',
    icon varchar(60) default '' null comment 'ICON',
    deleted_at timestamp null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
)
    comment '菜单表' engine=MyISAM collate=utf8mb4_unicode_ci;

create table funy_cms.notification
(
    id bigint unsigned auto_increment
        primary key,
    title varchar(255) default '' not null comment '标题',
    tag varchar(255) not null comment '分类: 根据业务类型进行定义',
    description json not null comment '详细信息',
    submitter_id tinyint null comment '发送者',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    read_at timestamp null comment '阅读时间'
)
    comment '消息通知' engine=MyISAM collate=utf8mb4_unicode_ci;

create table funy_cms.notification_user
(
    id bigint unsigned auto_increment
        primary key,
    notification_id varchar(255) default '' not null comment '标题',
    follow_id tinyint null comment '跟进人',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    read_at timestamp null comment '阅读时间'
)
    comment '消息通知关联用户表' engine=MyISAM collate=utf8mb4_unicode_ci;

create table funy_cms.permissions
(
    id bigint unsigned auto_increment
        primary key,
    name varchar(60) default '' not null comment '权限名',
    icon varchar(60) default 'link' null comment '权限图标',
    url varchar(60) default '' null comment '路径',
    status tinyint default 1 not null comment '状态: 1:正常; 2禁用',
    method varchar(255) default 'GET' null comment '方法名称',
    p_id int default 1 null comment '节点位置 1:根结点;',
    hidden tinyint default 2 null comment '是否隐藏 1:是 2否',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null
)
    comment '权限表' engine=MyISAM collate=utf8mb4_unicode_ci;

create table funy_cms.roles
(
    id bigint unsigned auto_increment
        primary key,
    name varchar(60) default '' not null comment '角色名',
    description varchar(255) null comment '角色描述',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    status int default 1 null comment '状态: 1: 启用; 2:禁用;',
    menu_ids varchar(255) null comment '拥有的菜单'
)
    comment '角色表' engine=MyISAM collate=utf8mb4_unicode_ci;

