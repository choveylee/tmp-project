CREATE TABLE IF NOT EXISTS banners
(
    id         VARCHAR(32)   NOT NULL,
    title      VARCHAR(1024) NOT NULL COMMENT '标题',
    abstract   TEXT          NOT NULL COMMENT '摘要',
    image_url  VARCHAR(1024) NOT NULL COMMENT '图片地址',
    link_url   VARCHAR(1024) NOT NULL COMMENT '链接地址',
    weight     INT           NOT NULL COMMENT '权重',
    status     INT           NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at DATETIME      NOT NULL,
    updated_at DATETIME      NOT NULL,
    deleted_at DATETIME      NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='轮播图';