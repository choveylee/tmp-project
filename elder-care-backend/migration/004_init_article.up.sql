CREATE TABLE IF NOT EXISTS article_categories
(
    id         VARCHAR(24)  NOT NULL,
    name       VARCHAR(32)  NOT NULL COMMENT '名称',
    weight     INT          NOT NULL COMMENT '权重',
    status     TINYINT(1)   NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='文章分类';

CREATE TABLE IF NOT EXISTS articles
(
    id          VARCHAR(24)   NOT NULL,
    category_id VARCHAR(24)   NOT NULL COMMENT '分类ID',
    title       VARCHAR(1024) NOT NULL COMMENT '标题',
    abstract    TEXT          NOT NULL COMMENT '摘要',
    content     MEDIUMTEXT    NOT NULL COMMENT '内容',
    cover_url   VARCHAR(1024) NOT NULL COMMENT '封面URL',
    link_url    VARCHAR(1024) NOT NULL COMMENT '链接URL',
    publish_at  DATETIME      NOT NULL COMMENT '发布时间',
    status      INT           NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at  DATETIME      NOT NULL,
    updated_at  DATETIME      NOT NULL,
    deleted_at  DATETIME      NULL,
    PRIMARY KEY (id),
    INDEX idx_category (category_id, publish_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='文章';