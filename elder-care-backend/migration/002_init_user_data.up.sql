INSERT INTO roles (id, code, name, is_admin, weight, created_at, updated_at, deleted_at)
VALUES
    ('66f66c9bd2dab6c1fa1ef5a5', 'ADMIN', '管理员', true, 10, '2026-05-12 10:24:00', '2026-05-12 10:24:00', NULL);

INSERT INTO users (id, role_id, name, mobile, password, status, login_at, created_at, updated_at, deleted_at)
VALUES ('66f7786bbd5f7726b801af6d', '66f66c9bd2dab6c1fa1ef5a5', '管理员', '13800138000', '$2a$10$Y01981PtfQuj5uFIVZterOH7P8El9PodXOm/L2hOy9TPLivcYcMli', 1, NULL, '2026-01-01 00:00:00', '2026-01-01 00:00:00', NULL);