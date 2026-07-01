CREATE TABLE IF NOT EXISTS course_categories
(
    id         VARCHAR(24) NOT NULL,
    name       VARCHAR(32) NOT NULL COMMENT '名称',
    weight     INT         NOT NULL COMMENT '权重',
    status     TINYINT(1)  NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程分类';

CREATE TABLE IF NOT EXISTS courses
(
    id              VARCHAR(24)   NOT NULL,
    category_id     VARCHAR(24)   NOT NULL COMMENT '分类ID',
    author          VARCHAR(255)  NOT NULL COMMENT '作者',
    source          VARCHAR(1024) NOT NULL COMMENT '来源',
    title           VARCHAR(1024) NOT NULL COMMENT '标题',
    abstract        TEXT          NOT NULL COMMENT '摘要',
    course_type     INT           NOT NULL COMMENT '课程类型 0-视频，1-图文',
    cover_url       VARCHAR(1024) NOT NULL COMMENT '封面URL',
    link_url        VARCHAR(1024) NOT NULL COMMENT '链接URL',
    publish_at      DATETIME      NULL COMMENT '发布时间',
    favourite_count INT           NOT NULL COMMENT '收藏数',
    view_count      INT           NOT NULL COMMENT '浏览数',
    status          INT           NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at      DATETIME      NOT NULL,
    updated_at      DATETIME      NOT NULL,
    deleted_at      DATETIME      NULL,
    PRIMARY KEY (id),
    INDEX idx_category (category_id, publish_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程';

CREATE TABLE IF NOT EXISTS course_tags
(
    id         VARCHAR(24)   NOT NULL,
    course_id  VARCHAR(24)   NOT NULL COMMENT '课程ID',
    name       VARCHAR(1024) NOT NULL COMMENT '标签',
    created_at DATETIME      NOT NULL,
    deleted_at DATETIME      NULL,
    PRIMARY KEY (id),
    INDEX idx_course (course_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程标签';

CREATE TABLE IF NOT EXISTS course_details
(
    id           VARCHAR(24) NOT NULL,
    detail       MEDIUMTEXT  NOT NULL COMMENT '课程详情',
    summary      MEDIUMTEXT  NOT NULL COMMENT '课程概述',
    objective    MEDIUMTEXT  NOT NULL COMMENT '授课目标',
    outline      MEDIUMTEXT  NOT NULL COMMENT '课程大纲',
    `references` MEDIUMTEXT  NOT NULL COMMENT '参考资料',
    created_at   DATETIME    NOT NULL,
    updated_at   DATETIME    NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程详情';

CREATE TABLE IF NOT EXISTS course_catalogs
(
    id         VARCHAR(24)   NOT NULL,
    parent_id  VARCHAR(24)   NOT NULL COMMENT '父级目录ID',
    name       VARCHAR(1024) NOT NULL COMMENT '名称',
    weight     INT           NOT NULL COMMENT '权重',
    status     INT           NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at DATETIME      NOT NULL,
    updated_at DATETIME      NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程目录';

CREATE TABLE IF NOT EXISTS course_videos
(
    id         VARCHAR(24)   NOT NULL,
    catalog_id VARCHAR(24)   NOT NULL COMMENT '目录ID',
    video_url  VARCHAR(1024) NOT NULL COMMENT '视频URL',
    format     VARCHAR(255)  NOT NULL COMMENT '格式',
    language   VARCHAR(255)  NOT NULL COMMENT '语言',
    size       VARCHAR(255)  NOT NULL COMMENT '大小',
    duration   VARCHAR(255)  NOT NULL COMMENT '时长',
    upload_at  DATETIME      NOT NULL COMMENT '上传时间',
    weight     INT           NOT NULL COMMENT '权重',
    status     INT           NOT NULL COMMENT '状态：1-正常，0-禁用',
    created_at DATETIME      NOT NULL,
    updated_at DATETIME      NOT NULL,
    deleted_at DATETIME      NULL,
    PRIMARY KEY (id),
    INDEX idx_course (course_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程视频';

CREATE TABLE IF NOT EXISTS course_favourites
(
    id         VARCHAR(24) NOT NULL,
    course_id  VARCHAR(24) NOT NULL COMMENT '课程ID',
    user_id    VARCHAR(24) NOT NULL COMMENT '用户ID',
    created_at DATETIME    NOT NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (id),
    INDEX idx_course (course_id),
    INDEX idx_user (user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='课程收藏';