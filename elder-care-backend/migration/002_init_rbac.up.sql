CREATE TABLE IF NOT EXISTS permissions
(
    id              VARCHAR(32)  NOT NULL,
    code            VARCHAR(128) NOT NULL COMMENT '编码',
    name            VARCHAR(128) NOT NULL COMMENT '名称',
    permission_type INT          NOT NULL COMMENT '权限类型：1-菜单，2-按钮，3-API，4-数据',
    group_id        VARCHAR(32)  NULL COMMENT '权限分组ID',
    resource_code   VARCHAR(128) NOT NULL COMMENT '资源编码',
    action_code     VARCHAR(128) NOT NULL COMMENT '动作编码',
    weight          INT          NOT NULL COMMENT '权重，用于排序',
    status          TINYINT(1)   NOT NULL COMMENT '状态：1-正常，0-禁用',
    remarks         TEXT         NULL,
    created_at      DATETIME     NOT NULL,
    updated_at      DATETIME     NOT NULL,
    deleted_at      DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX idx_code (code),
    INDEX idx_name (name),
    INDEX idx_group (group_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='权限';

CREATE TABLE IF NOT EXISTS role_permissions
(
    id            VARCHAR(32) NOT NULL,
    role_id       VARCHAR(32) NOT NULL COMMENT '角色ID',
    permission_id VARCHAR(32) NOT NULL COMMENT '权限ID',
    created_at    DATETIME    NOT NULL,
    updated_at    DATETIME    NOT NULL,
    deleted_at    DATETIME    NULL,
    PRIMARY KEY (id),
    INDEX idx_role_permission (role_id, permission_id),
    INDEX idx_permission (permission_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='角色权限关系';

CREATE TABLE IF NOT EXISTS menus
(
    id         VARCHAR(32)  NOT NULL,
    parent_id  VARCHAR(32)  NOT NULL COMMENT '父级ID',
    code       VARCHAR(128) NOT NULL COMMENT '编码',
    name       VARCHAR(128) NOT NULL COMMENT '名称',
    menu_type  INT          NOT NULL COMMENT '菜单类型 0-CATALOG，1-MENU，2-HIDDEN_ROUTE，3-BUTTON',
    route_name VARCHAR(255) NULL COMMENT '前端唯一路由名称，由前端路由表解析 path 和组件',
    icon_name  VARCHAR(255) NULL COMMENT '图标名，如 Histogram/Notebook',
    weight     INT          NOT NULL COMMENT '权重，用于排序',
    status     TINYINT(1)   NOT NULL COMMENT '状态：1-正常，0-禁用',
    remarks    TEXT         NULL COMMENT '备注',
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX index_code (code),
    INDEX index_name (name),
    INDEX idx_parent (parent_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='菜单';

CREATE TABLE IF NOT EXISTS role_menus
(
    id         VARCHAR(32) NOT NULL,
    role_id    VARCHAR(32) NOT NULL COMMENT '角色ID',
    menu_id    VARCHAR(32) NOT NULL COMMENT '菜单ID',
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (id),
    INDEX idx_role_menu (role_id, menu_id),
    INDEX idx_menu (role_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='角色菜单关系';

CREATE TABLE IF NOT EXISTS api_resources
(
    id            VARCHAR(32)  NOT NULL,
    code          VARCHAR(128) NOT NULL COMMENT '编码',
    name          VARCHAR(128) NOT NULL COMMENT '名称',
    method        VARCHAR(16)  NOT NULL COMMENT '方法 GET/POST/PUT/DELETE',
    path          VARCHAR(255) NOT NULL COMMENT '路径',
    resource_code VARCHAR(128) NOT NULL COMMENT '资源编码',
    action_code   VARCHAR(128) NOT NULL COMMENT '动作编码',
    auth_required TINYINT(1)   NOT NULL COMMENT '是否需要认证 1-是，0-否',
    status        TINYINT(1)   NOT NULL COMMENT '状态：1-正常，0-禁用',
    remarks       TEXT         NULL COMMENT '备注',
    created_at    DATETIME     NOT NULL,
    updated_at    DATETIME     NOT NULL,
    deleted_at    DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX idx_code (code),
    INDEX idx_method_path (method, path)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='API资源';

CREATE TABLE IF NOT EXISTS permission_api_resources
(
    id              VARCHAR(32) NOT NULL,
    permission_id   VARCHAR(32) NOT NULL COMMENT '权限ID',
    api_resource_id VARCHAR(32) NOT NULL COMMENT 'API资源ID',
    created_at      DATETIME    NOT NULL,
    updated_at      DATETIME    NOT NULL,
    deleted_at      DATETIME    NULL,
    PRIMARY KEY (id),
    INDEX idx_api_resources (permission_id, api_resource_id),
    INDEX idx_resources (api_resource_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='权限与API映射';