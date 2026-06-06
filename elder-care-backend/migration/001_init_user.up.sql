CREATE TABLE IF NOT EXISTS roles
(
    id         VARCHAR(32)  NOT NULL,
    code       VARCHAR(32)  NOT NULL,
    name       VARCHAR(255) NOT NULL COMMENT '名称',
    is_admin   TINYINT(1)   NOT NULL COMMENT '是否是管理员 1-是，0-否',
    weight     INT          NOT NULL COMMENT '权重，用于排序',
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX idx_code (code),
    INDEX idx_name (name)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='角色';

CREATE TABLE IF NOT EXISTS users
(
    id         VARCHAR(32)  NOT NULL,
    role_id    VARCHAR(32)  NOT NULL COMMENT '角色ID',
    mobile     VARCHAR(20)  NOT NULL COMMENT '手机号',
    name       VARCHAR(255) NOT NULL COMMENT '名字',
    password   VARCHAR(255) NOT NULL COMMENT '密码',
    status     TINYINT(1)   NOT NULL DEFAULT 1 COMMENT '状态：1-正常，0-禁用',
    login_at   DATETIME     NULL COMMENT '上次登录时间',
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX idx_mobile (mobile),
    KEY idx_role (role_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户';

CREATE TABLE IF NOT EXISTS access_logs
(
    `id`         VARCHAR(32)  NOT NULL,
    `user_id`    VARCHAR(32)  NOT NULL COMMENT '用户ID',
    `lang`       VARCHAR(255) NOT NULL COMMENT '语言',
    `real_ip`    VARCHAR(64)  NOT NULL COMMENT '真实IP',
    `created_at` DATETIME     NOT NULL,
    PRIMARY KEY (`id`),
    INDEX idx_user (`user_id`)
) NGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='访问日志';
